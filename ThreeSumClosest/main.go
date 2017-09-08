package ThreeSumClosest

func threeSum(countTable map[int]int, nums []int, aim int) bool {
	length := len(nums)
	var target int
	for i := range nums {
		target = aim - nums[i]
		for j := i + 1; j < length && j != i; j++ {
			if count, exist := countTable[target-nums[j]]; exist {
				if nums[i] == nums[j] && nums[i] == target-nums[j] {
					if count >= 3 {
						return true
					}
				} else if nums[i] == target-nums[j] || nums[j] == target-nums[j] {
					if count >= 2 {
						return true
					}
				} else {
					return true
				}
			}
		}
	}
	return false
}

func threeSumClosest(nums []int, target int) int {
	res := nums[0] + nums[1] + nums[2]
	maxBias := res - target
	length := len(nums)
	if maxBias < 0 {
		maxBias = -maxBias
	}
	countTable := make(map[int]int, length)
	for i := 0; i < length; i++ {
		if _, exist := countTable[nums[i]]; exist {
			countTable[nums[i]]++
		} else {
			countTable[nums[i]] = 1
		}
	}

	if threeSum(countTable, nums, target) {
		return target
	}
	for i := 1; i < maxBias; i++ {
		if threeSum(countTable, nums, target+i) {
			return target + i
		} else if threeSum(countTable, nums, target-i) {
			return target - i
		}
	}
	return res
}
