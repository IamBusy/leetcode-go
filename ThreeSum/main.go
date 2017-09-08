package ThreeSum

func threeSum(nums []int) [][]int {
	length := len(nums)
	countTable := make(map[int]int, length)
	for i := 0; i < length; i++ {
		if _, exist := countTable[nums[i]]; exist {
			countTable[nums[i]]++
		} else {
			countTable[nums[i]] = 1
		}
	}
	var target int
	var res [][]int
	resExist := map[int]map[int]int{}
	for i := range nums {
		target = 0 - nums[i]
		for j := i + 1; j < length && j != i; j++ {
			if count, exist := countTable[target-nums[j]]; exist {
				valid := false
				if nums[i] == nums[j] && nums[i] == target-nums[j] {
					if count >= 3 {
						valid = true
					}
				} else if nums[i] == target-nums[j] || nums[j] == target-nums[j] {
					if count >= 2 {
						valid = true
					}
				} else {
					valid = true
				}
				if valid {
					first, second, third := nums[i], nums[j], target-nums[j]
					var tmp int
					if first > second {
						tmp = second
						second = first
						first = tmp
					}
					if second > third {
						tmp = third
						third = second
						second = tmp
					}
					if first > second {
						tmp = second
						second = first
						first = tmp
					}

					if _, exist := resExist[first]; exist {
						if value, exist := resExist[first][second]; exist && value == third {
							continue
						}
					} else {
						resExist[first] = map[int]int{}
					}
					res = append(res, []int{first, second, third})
					resExist[first][second] = third
				}
			}
		}
	}
	return res
}
