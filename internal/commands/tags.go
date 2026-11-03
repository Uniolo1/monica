// SPDX-FileCopyrightText: NONE
// SPDX-License-Identifier: Unlicense

/*
 * This is free and unencumbered software released into the public domain.
 *
 * Anyone is free to copy, modify, publish, use, compile, sell, or distribute
 * this software, either in source code form or as a compiled binary, for any
 * purpose, commercial or non-commercial, and by any means.
 *
 * In jurisdictions that recognize copyright laws, the author or authors of this
 * software dedicate any and all copyright interest in the software to the
 * public domain. We make this dedication for the benefit of the public at large
 * and to the detriment of our heirs and successors. We intend this dedication
 * to be an overt act of relinquishment in perpetuity of all present and future
 * rights to this software under copyright law.
 *
 * THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
 * ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
 * WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package commands

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	lua "github.com/yuin/gopher-lua"

	"github.com/uniolo1/monica/internal/abstractions"
	"github.com/uniolo1/monica/internal/helpers"
)

type Tag struct {
	Contents string `json:"contents"` // contents of the tag
	Owner    string `json:"owner"`    // owner of the tag
	Uses     uint   `json:"uses"`     // how many times it has been used
	Lua      bool   `json:"is_lua"`   // weather or not its a lua tag
}

// dictionary of dictionary of tags
var tags map[string]map[string]Tag = make(map[string]map[string]Tag)

var tagMutex sync.Mutex
var maxTagSize int = 1000

// little wrapper I made to ensure that each zone has a map[string]
// NOTE: does not get mutex, caller has to do that!
func ensureTagZone(zone string) {
	_, exists := tags[zone]
	if !exists {
		tags[zone] = make(map[string]Tag)
	}
}

// helper used in 'tag info'
func formatDoesOrDoesNot(input bool) string {
	if input {
		return "does"
	} else {
		return "does not"
	}
}

// helper for evalTag
func openSafeLibs(L *lua.LState) {
	libs := []struct {
		name string
		fn   lua.LGFunction
	}{
		{lua.TabLibName, lua.OpenTable},
		{lua.StringLibName, lua.OpenString},
		{lua.MathLibName, lua.OpenMath},
		//{lua.CoroutineLibName, lua.OpenCoroutine}, // doubt many scripts will need this...
	}

	for _, lib := range libs {
		L.Push(L.NewFunction(lib.fn))
		L.Push(lua.LString(lib.name))
		L.Call(1, 0)
	}
}

// I will note I did make heavy use of AI here to make sure I don't have a security flaw
func evalTag(contents string, arguments string, usesLua bool) string {
	if !usesLua {
		return contents
	}

	L := lua.NewState(lua.Options{
		SkipOpenLibs: true,

		// limit resource usage
		MinimizeStackMemory: true,
		CallStackSize:       128,
		RegistrySize:        2048,
		RegistryMaxSize:     4096,
		RegistryGrowStep:    32,
	})
	defer L.Close()
	openSafeLibs(L)

	var output bytes.Buffer // output

	// maximum stuff that can be printed
	const maxOutputSize = 64 * 1024

	// override print function
	L.SetGlobal("print", L.NewFunction(func(L *lua.LState) int {
		for i := 1; i <= L.GetTop(); i++ {
			if output.Len() >= maxOutputSize {
				break
			}

			if i > 1 {
				output.WriteString("\t")
			}

			remaining := maxOutputSize - output.Len()
			value := L.ToStringMeta(L.Get(i)).String()

			if len(value) > remaining {
				output.WriteString(value[:remaining])
				break
			}

			output.WriteString(value)
		}

		if output.Len() < maxOutputSize {
			output.WriteString("\n")
		}

		return 0
	}))

	// setup lua environment
	L.SetGlobal("arguments", lua.LString(arguments))

	// timeout
	ctx, cancel := context.WithTimeout(
		context.Background(),
		4*time.Second,
	)
	defer cancel()
	L.SetContext(ctx)

	err := L.DoString(contents)
	if err != nil {
		return fmt.Sprintf("%v", err)
	}

	ret := output.String()
	if ret == "" {
		return "Tag (using Lua) did not output anything!"
	} else {
		return ret
	}
}

func removeTag(args *abstractions.CommandContext) (string, bool) {
	name := strings.ToLower(args.ArgumentGetter("name", abstractions.Command_argument_string))
	if name == "" {
		return "Missing required parameter: " + args.FormatQuote("name"), true
	}

	// get 'zone' field
	var zone string
	if args.UniqueChannelID != "" {
		zone = args.UniqueChannelID
	} else if args.UniqueCallerID != "" {
		zone = args.UniqueCallerID
	} else {
		return "Unknown zone!", true
	}

	tagMutex.Lock()
	defer tagMutex.Unlock()
	ensureTagZone(zone) // ensure tags[zone] exists!

	tag, exists := tags[zone][name]
	if !exists {
		return args.FormatQuote(name) + " does not exist!", true
	}

	if tag.Owner != args.UniqueCallerID {
		return "Cannot remove tag you do not own!", true
	}

	delete(tags, name)
	return "Removed tag " + args.FormatQuote(name) + "!", false
}

func modifyTag(args *abstractions.CommandContext) (string, bool) {
	name := strings.ToLower(args.ArgumentGetter("name", abstractions.Command_argument_string))
	contents := args.ArgumentGetter("contents", abstractions.Command_argument_string)

	if name == "" {
		return "Missing required parameter: " + args.FormatQuote("name"), true
	} else if contents == "" {
		return "Missing required parameter: " + args.FormatQuote("contents"), true
	}

	if len(contents) > maxTagSize {
		return fmt.Sprintf("Tag is %d bytes long, longer then the limit of %d bytes!", len(contents), maxTagSize), true
	}

	lua := helpers.StringToBool(args.ArgumentGetter("lua", abstractions.Command_argument_bool), false)

	var zone string
	if args.UniqueChannelID != "" {
		zone = args.UniqueChannelID
	} else if args.UniqueCallerID != "" {
		zone = args.UniqueCallerID
	} else {
		return "Unknown zone!", true
	}

	tagMutex.Lock()
	defer tagMutex.Unlock()
	ensureTagZone(zone) // ensure tags[zone] exists!

	ensureTagZone(zone) // ensure tags[zone] exists!
	var uses uint = 0

	tag, exists := tags[zone][name]
	if exists {
		uses = tag.Uses
		if args.UniqueCallerID != tag.Owner {
			return "Only the owner can modify a tag!", true
		}
	}

	tags[zone][name] = Tag{
		Contents: contents,
		Owner:    args.UniqueCallerID,
		Uses:     uses,
		Lua:      lua,
	}

	if exists {
		return args.FormatQuote(name) + " modified!", false
	} else {
		return args.FormatQuote(name) + " created!", false
	}
}

func getTagInfo(args *abstractions.CommandContext) (string, bool) {
	name := strings.ToLower(args.ArgumentGetter("name", abstractions.Command_argument_string))
	if name == "" {
		return "Missing required parameter: " + args.FormatQuote("name"), true
	}

	var zone string
	if args.UniqueChannelID != "" {
		zone = args.UniqueChannelID
	} else if args.UniqueCallerID != "" {
		zone = args.UniqueCallerID
	} else {
		return "Unknown zone!", true
	}

	tagMutex.Lock()
	defer tagMutex.Unlock()
	ensureTagZone(zone) // ensure tags[zone] exists!

	tag, exists := tags[zone][name]
	if !exists {
		return args.FormatQuote(name) + " does not exist!", true
	}

	// format it!
	return "Tag " + args.FormatQuote(name) + " " + formatDoesOrDoesNot(tag.Lua) + " use Lua, it is owned by " + args.FormatQuote(tag.Owner) +
			", and has been used " + args.FormatQuote(strconv.FormatUint(uint64(tag.Uses), 10)) + " times!" +
			"\n" + args.FormatBlock(tag.Contents),
		false
}

func useTag(args *abstractions.CommandContext) (string, bool) {
	name := strings.ToLower(args.ArgumentGetter("name", abstractions.Command_argument_string))
	if name == "" {
		return "Missing required parameter: " + args.FormatQuote("name"), true
	}

	var zone string
	if args.UniqueChannelID != "" {
		zone = args.UniqueChannelID
	} else if args.UniqueCallerID != "" {
		zone = args.UniqueCallerID
	} else {
		return "Unknown zone!", true
	}

	var tag Tag

	{ // inline scope here due to mutex
		tagMutex.Lock()
		defer tagMutex.Unlock()
		ensureTagZone(zone) // ensure tags[zone] exists!

		var exists bool // declare it here to fix an issue
		tag, exists = tags[zone][name]

		if !exists {
			return args.FormatQuote(name) + " does not exist!", true
		}

		// increment the uses
		tag.Uses++
		tags[zone][name] = tag
	} // mutex is unlocked here (lua code may take *some* time to execute)

	if tag.Lua {
		if os.Getenv("DISABLE_LUA") != "" && os.Getenv("DISABLE_LUA") != "false" {
			return "Lua tags have been disabled!", false
		}

		err := args.Defer(false)
		if err != nil {
			return fmt.Sprintf("Error deferring: %s", err), true
		}
	}

	arguments := args.ArgumentGetter("arguments", abstractions.Command_argument_string)
	reply := evalTag(tag.Contents, arguments, tag.Lua)

	if reply == "" {
		// for debugging: kept cuz why not
		return args.FormatQuote("evalTag") + " returned an empty string, this should NEVER HAPPEN, report this!\n" +
			args.FormatBlock(fmt.Sprintf("doLua: %v\nArgs: \"%s\"", tag.Lua, arguments)), true
	}

	return reply, false
}

func addTagCommands(adder abstractions.CommandAdder) error {
	var err error

	err = adder(abstractions.Command{
		Name:        "tag modify",
		Description: "Modify a tag",
		Arguments: []abstractions.CommandArguments{{
			Name:        "name",
			Description: "Name of the tag",
			Required:    true,
			Arg_type:    abstractions.Command_argument_string,
		}, {
			Name:        "contents",
			Description: "Contents of the tag",
			Required:    true,
			Arg_type:    abstractions.Command_argument_string,
		}, {
			Name:        "lua",
			Description: "Weather or not to use Lua when evaluating",
			Required:    false,
			Arg_type:    abstractions.Command_argument_bool,
		}},
		Handler: modifyTag,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "tag delete",
		Description: "Deletes a tag (tag-owner only!)",
		Arguments: []abstractions.CommandArguments{{
			Name:        "name",
			Description: "Name of the tag",
			Required:    true,
			Arg_type:    abstractions.Command_argument_string,
		}},
		Handler: removeTag,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "tag info",
		Description: "Get info about a tag",
		Arguments: []abstractions.CommandArguments{{
			Name:        "name",
			Description: "Name of the tag",
			Required:    true,
			Arg_type:    abstractions.Command_argument_string,
		}},
		Handler: getTagInfo,
	})
	if err != nil {
		return err
	}

	err = adder(abstractions.Command{
		Name:        "tag run",
		Description: "Run a tag",
		Arguments: []abstractions.CommandArguments{{
			Name:        "name",
			Description: "Name of the tag",
			Required:    true,
			Arg_type:    abstractions.Command_argument_string,
		}, {
			Name:        "arguments",
			Description: "Arguments of the tag",
			Required:    false,
			Arg_type:    abstractions.Command_argument_string,
		}},
		Handler: useTag,
	})
	if err != nil {
		return err
	}

	return persistTags()
}

// it is imperative this function be called when 'cmd' exits!
func SaveTags() error {
	tagFile := os.Getenv("TAGS_FILE")
	if tagFile == "" {
		return nil
	}

	data, err := json.MarshalIndent(tags, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(tagFile, data, 0644)
}

// logic to make tags persist
func persistTags() error {
	tagFile := os.Getenv("TAGS_FILE")

	if tagFile == "" {
		slog.Warn("'TAGS_FILE' is unset, TAGS WILL NOT PERSIST!")
		return nil
	}

	data, err := os.ReadFile(tagFile)
	if err != nil {
		if os.IsNotExist(err) {
			slog.Warn("'TAGS_FILE' does not exist, starting new!")
			return nil
		}
		return err
	}

	return json.Unmarshal(data, &tags)
}
