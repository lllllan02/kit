package kit

import "unicode"

func splitToStrings(s string, upperCase bool) []string {
	var runes [][]rune
	charType := 0
	lastCharType := 0

	for _, r := range s {
		switch true {
		case isLower(r):
			charType = 1
		case isUpper(r):
			charType = 2
		case isDigit(r):
			charType = 3
		default:
			charType = 4
		}

		if charType == lastCharType {
			runes[len(runes)-1] = append(runes[len(runes)-1], r)
		} else {
			runes = append(runes, []rune{r})
		}
		lastCharType = charType
	}

	for i := 0; i < len(runes)-1; i++ {
		if isUpper(runes[i][0]) && isLower(runes[i+1][0]) {
			last, _ := Last(runes[i])
			runes[i+1] = InsertAt(runes[i+1], 0, last)
			runes[i] = DropRight(runes[i], 1)
		}
	}

	result := make([]string, 0, len(runes))

	for _, r := range runes {
		if len(r) > 0 && (unicode.IsLetter(r[0]) || isDigit(r[0])) {
			if upperCase {
				result = append(result, string(toUpperAll(r)))
			} else {
				result = append(result, string(toLowerAll(r)))
			}
		}
	}

	return result
}

// isDigit checks if a character is digit ('0' to '9')
func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// isLower checks if a character is lower case ('a' to 'z')
func isLower(r rune) bool {
	return r >= 'a' && r <= 'z'
}

// isUpper checks if a character is upper case ('A' to 'Z')
func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

// toLower converts a character  'A' to 'Z' to its lower case
func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

// toLowerAll converts a character  'A' to 'Z' to its lower case
func toLowerAll(rs []rune) []rune {
	for i := range rs {
		rs[i] = toLower(rs[i])
	}
	return rs
}

// toUpper converts a character  'a' to 'z' to its upper case
func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 32
	}
	return r
}

// toUpperAll converts a character  'a' to 'z' to its upper case
func toUpperAll(rs []rune) []rune {
	for i := range rs {
		rs[i] = toUpper(rs[i])
	}
	return rs
}
