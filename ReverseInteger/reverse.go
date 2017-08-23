package ReverseInteger

import "math"

func reverse(x int) int {
	nums := make([]uint8, 0, 10)
	idx := 0
	value := int(math.Abs(float64(x)))
	res := 0
	for value >= 10 {
		nums = append(nums, uint8(value%10))
		value = value / 10
		idx++
	}
	nums = append(nums, uint8(value))
	for i := 0; i <= idx; i++ {
		res += int(nums[i]) * int(math.Pow(float64(10), float64(idx-i)))
	}

	if res >= 2147483648 {
		return 0
	} else if x > 0 {
		return res
	} else {
		return -res
	}
}
