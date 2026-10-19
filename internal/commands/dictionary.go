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
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func formatTabs(s string) string {
	var result strings.Builder
	column := 0

	for _, r := range s {
		switch r {
		case '\t':
			spaces := 6 - (column % 6)
			result.WriteString(strings.Repeat(" ", spaces))
			column += spaces
		case '\n':
			result.WriteRune(r)
			column = 0
		default:
			result.WriteRune(r)
			column++
		}
	}

	return result.String()

}

func getDefinition(query string) string {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// I use datamuse because it does not require an API key - and it's free!
	apiURL := "https://api.datamuse.com/words?sp=" +
		url.QueryEscape(query) + "&md=d"

	resp, err := client.Get(apiURL)
	if err != nil {
		return fmt.Sprintf("ERROR: %s", err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("Got HTTP status code '%d' when looking up definition!", resp.StatusCode)
	}

	var data []struct {
		Defs []string `json:"defs"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return fmt.Sprintf("ERROR: %s", err.Error())
	}

	if len(data) == 0 || len(data[0].Defs) == 0 {
		return "No definition was found!"
	}

	// [:3] limits definition to a max of 3, chosen because :3
	return strings.Join(data[0].Defs[:3], "\n")
}
