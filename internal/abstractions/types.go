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

package abstractions

type CommandArgumentGetter func(string, Command_argument_type) string
type CommandCodeOrQuoteFormatter func(string) string
type CommandHandler func(*CommandContext) (
	string, // the reply to the message
	bool, // weather the message should be ephemeral (only shown to user that sent the command)
)
type CommandAdder func(Command) error

// the <sarcasm>infamously awesome</sarcasm> go enums!
type Command_argument_type int

const (
	Command_argument_type_string = iota
	Command_argument_type_int
)

type CommandArguments struct {
	Name        string
	Description string
	Required    bool                  // NOTE: Due to platform-specific considerations, required parameters may not be present :3
	Sensitive   bool                  // platform's may interpret this differently or ignore it
	Arg_type    Command_argument_type // NOTE: converted into a string (due to platform-specific considerations), must be converted back manually!
}

type Command struct {
	Name        string
	Description string
	NSFW        bool // who knows might add one in the future, you never know
	Arguments   []CommandArguments
	Handler     CommandHandler
}

// what a command is given as its parameters
type CommandContext struct {
	ArgumentGetter    CommandArgumentGetter // REMINDER: Always returns 'string' regardless of argument type
	FormatCodeOrQuote CommandCodeOrQuoteFormatter
	Framework         string // framework being used
}
