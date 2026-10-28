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

package helpers

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func RollDice(input string) (uint, []uint, error) {
	input = strings.TrimSpace(input)

	// parse the string
	parts := strings.FieldsFunc(input, func(r rune) bool {
		return r == 'd' || r == 'D'
	})

	// check if the user inputted a valid dice format
	if len(parts) != 2 {
		return 0, nil, fmt.Errorf("invalid dice format %q: expected NdS", input)
	}

	count, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, nil, fmt.Errorf("invalid dice count: %w", err)
	}

	sides, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, nil, fmt.Errorf("invalid dice sides: %w", err)
	}

	if count < 1 || count > 20 {
		return 0, nil, fmt.Errorf("dice count must be between 1 and 20")
	}

	if sides < 2 || sides > 500 {
		return 0, nil, fmt.Errorf("dice sides must be between 2 and 500")
	}

	results := make([]uint, count)
	var total uint = 0

	for i := range results {
		results[i] = uint(rand.Intn(sides) + 1)
		total += results[i]
	}

	return total, results, nil
}
