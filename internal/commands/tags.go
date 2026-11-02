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
	"strconv"
	"strings"
	"sync"

	"github.com/uniolo1/monica/internal/abstractions"
)

type Tag struct {
	Contents string `json:"contents"` // contents of the tag
	Owner    string `json:"owner"`    // owner of the tag
	Uses     uint   `json:"uses"`     // how many times it has been used
	Lua      bool   `json:"is_lua"`   // weather or not its a lua tag
	NSFW     bool   `json:"is_nsfw"`  // I need better friends...
}

// dictionary of dictionary of tags
var tags map[string]map[string]Tag = make(map[string]map[string]Tag)

var tagsMutex map[string]*sync.Mutex = make(map[string]*sync.Mutex)
var tagsMutexMutex sync.Mutex  // these do practically the same thing
var globalTagsMutex sync.Mutex // these do practically the same thing

func getZoneMutex(zone string) *sync.Mutex {
	tagsMutexMutex.Lock()
	defer tagsMutexMutex.Unlock()

	mutex, exists := tagsMutex[zone]
	if !exists {
		mutex = &sync.Mutex{}
		tagsMutex[zone] = mutex
	}

	return mutex
}

// little wrapper I made to ensure that each zone has a map[string]
// now that I think about it, should probably merge with getZoneMutex... meh, I made this second anyways
func ensureTagZone(zone string) {
	globalTagsMutex.Lock()
	defer globalTagsMutex.Unlock()

	_, exists := tags[zone]
	if !exists {
		tags[zone] = make(map[string]Tag)
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
		return "Unknown caller!", true
	}

	// mutex shenanigans
	zoneMutex := getZoneMutex(zone)
	zoneMutex.Lock()
	defer zoneMutex.Unlock()
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

	var zone string
	if args.UniqueChannelID != "" {
		zone = args.UniqueChannelID
	} else if args.UniqueCallerID != "" {
		zone = args.UniqueCallerID
	} else {
		return "Unknown caller!", true
	}

	// mutex shenanigans
	zoneMutex := getZoneMutex(zone)
	zoneMutex.Lock()
	defer zoneMutex.Unlock()

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
		Lua:      false,
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
		return "Unknown caller!", true
	}

	// mutex shenanigans
	zoneMutex := getZoneMutex(zone)
	zoneMutex.Lock()
	defer zoneMutex.Unlock()
	ensureTagZone(zone) // ensure tags[zone] exists!

	tag, exists := tags[zone][name]
	if !exists {
		return args.FormatQuote(name) + " does not exist!", true
	}

	// format it!
	return "Tag " + args.FormatQuote(name) +
			" is owned by " + args.FormatQuote(tag.Owner) +
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
		return "Unknown caller!", true
	}

	// mutex shenanigans
	zoneMutex := getZoneMutex(zone)
	zoneMutex.Lock()
	defer zoneMutex.Unlock()
	ensureTagZone(zone) // ensure tags[zone] exists!

	tag, exists := tags[zone][name]

	if !exists {
		return args.FormatQuote(name) + " does not exist!", true
	} else if tag.Lua {
		return "Lua support is not yet implemented!", true
	} else if tag.NSFW && !args.NSFW {
		return "Cannot use tag marked as NSFW here!", true
	}

	// increment the uses
	tag.Uses++
	tags[zone][name] = tag

	return tag.Contents, false
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

	return nil
}
