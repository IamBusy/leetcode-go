package StringtoInteger

import "math"

func myAtoi(str string) int {
	idx := 0
	length := len(str)
	numbers := make([]int, 0, 10)
	isNeg := false
	result := 0
	for ; idx < length; idx++ {
		if str[idx] != ' ' {
			break
		}
	}
	if idx >= length {
		return 0
	}
	if str[idx] == '+' {
		isNeg = false
	} else if str[idx] == '-' {
		isNeg = true
	} else if str[idx] >= '0' && str[idx] <= '9' {
		numbers = append(numbers, int(str[idx])-48)
	} else {
		return 0
	}
	idx++

	for ; idx < length && str[idx] <= '9' && str[idx] >= '0'; idx++ {
		numbers = append(numbers, int(str[idx])-48)
	}

	if len(numbers) >= 11 || (len(numbers) == 10 && numbers[0] >= 3) {
		result = -math.MinInt32
	} else {
		for i := len(numbers) - 1; i >= 0; i-- {
			result += numbers[i] * int(math.Pow(float64(10), float64(len(numbers)-i-1)))
		}
	}

	if isNeg {

		if result >= -math.MinInt32 {
			return math.MinInt32
		} else {
			return -result
		}

	} else {
		if result >= math.MaxInt32 {
			return math.MaxInt32
		} else {
			return result
		}

	}

}
