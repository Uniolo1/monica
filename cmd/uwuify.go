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

package main

import "strings"

func transfer_case(original, replacement string) string {
	// replacement is assumed to be lowercase

	// if they are the same
	if original == replacement {
		return original
	}

	// if they original is upper
	if original == strings.ToUpper(original) {
		return strings.ToUpper(replacement)
	}

	if len(original) > 0 &&
		original[:1] == strings.ToUpper(original[:1]) {
		return strings.ToUpper(replacement[:1]) + replacement[1:]
	}

	return replacement
}

// based on a C library I made - https://github.com/uniolo1/libuwu
func uwuify(input string) string {
	words := strings.Fields(input)

	for i, word := range words {
		// NOTE smaller selection of words then in libuwu
		switch strings.ToLower(word) {
		case "love":
			words[i] = transfer_case(word, "wuv")
		case "loved":
			words[i] = transfer_case(word, "wuved")
		case "this":
			words[i] = transfer_case(word, "dis")
		case "small":
			words[i] = transfer_case(word, "smol")
		case "angry":
			words[i] = transfer_case(word, "angi")
		case "boy":
			words[i] = transfer_case(word, "boi")
		default:
			result := []rune(word)

			for j, char := range result {
				switch char {
				case 'r', 'l':
					result[j] = 'w'
				case 'R', 'L':
					result[j] = 'W'
				}
			}

			words[i] = string(result)
		}
	}

	return strings.Join(words, " ")
}
