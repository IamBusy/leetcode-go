package LetterCombinationsOfAPhoneNumber

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	board := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	num := 1
	for i := range digits {
		if digits[i] == '7' || digits[i] == '9' {
			num *= 4
		} else {
			num *= 3
		}
	}
	repeat := 1
	res := make([]string, num)
	length := len(digits)

	// Traverse the digits
	for j := length - 1; j >= 0; j-- {
		idx := 0
		// Traverse the result
		for i := 0; i < num; i += repeat {
			for k := 0; k < repeat; k++ {
				res[i+k] = board[digits[j]][idx] + res[i+k]
			}
			idx++
			if idx >= 4 && (digits[j] == '7' || digits[j] == '9') {
				idx = 0
			} else if idx >= 3 && digits[j] != '7' && digits[j] != '9' {
				idx = 0
			}
		}
		if digits[j] == '7' || digits[j] == '9' {
			repeat *= 4
		} else {
			repeat *= 3
		}
	}
	return res
}
