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
