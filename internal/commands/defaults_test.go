// SPDX-FileCopyrightText: NONE
// SPDX-License-Identifier: Unlicense

package commands

import "testing"

// little test to ensure non of the animegirl gifs are the same, since I am likely to include dups... oops!
func TestItemsAreUnique(t *testing.T) {
	seen := make(map[string]bool)

	for _, item := range DEFAULT_ANIMEGIRL_GIFS {
		if seen[item] {
			t.Errorf("duplicate item found: \"%q\"", item)
		}
		seen[item] = true
	}
}
