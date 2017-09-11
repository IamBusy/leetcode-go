package FourSum

import "fmt"

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	if length < 4 {
		return [][]int{}
	}
	countTable := make(map[int]int, length)
	for i := 0; i < length; i++ {
		if _, exist := countTable[nums[i]]; exist {
			countTable[nums[i]]++
		} else {
			countTable[nums[i]] = 1
		}
	}

	availableTwoSum := make(map[int][][3]int)
	availableTwoSumLength := 0
	var low int
	var high int
	var res [][]int
	//time1 := time.Now()
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			sum := nums[i] + nums[j]
			if nums[i] < nums[j] {
				low = nums[i]
				high = nums[j]
			} else {
				low = nums[j]
				high = nums[i]
			}
			if arr, exist := availableTwoSum[sum]; exist {
				repeat := false
				for k := 0; k < len(arr); k++ {
					if arr[k][0] == low && arr[k][1] == high {
						repeat = true
						break
					}
				}
				if !repeat {
					availableTwoSum[sum] = append(arr, [3]int{low, high, 0})
					availableTwoSumLength++
				}
			} else {
				availableTwoSum[sum] = [][3]int{
					{low, high, 0},
				}
				availableTwoSumLength++
			}
		}
	}
	//time2 := time.Now()
	//fmt.Println(availableTwoSum)
	resTable := map[int]map[int]map[int]int{}
	for value := range availableTwoSum {
		if a2, exist := availableTwoSum[target-value]; exist {
			a1 := availableTwoSum[value]
			for i := 0; i < len(a1); i++ {
				for j := 0; j < len(a2); j++ {
					unique := make(map[int]int)
					cand := []int{
						a1[i][0],
						a1[i][1],
						a2[j][0],
						a2[j][1],
					}
					sort.Ints(cand)
					for k := 0; k < 4; k++ {
						if _, exist := unique[cand[k]]; exist {
							unique[cand[k]]++
						} else {
							unique[cand[k]] = 1
						}
					}
					valid := true
					for k := range unique {
						if countTable[k] < unique[k] {
							valid = false
							break
						}
					}
					//for k:=0; k<4; k++ {
					//
					//}

					if valid {
						if _, exist := resTable[cand[0]]; exist {
							if _, exist := resTable[cand[0]][cand[1]]; exist {
								if val, exist := resTable[cand[0]][cand[1]][cand[2]]; exist && val == cand[3] {
									continue
								}
							} else {
								resTable[cand[0]][cand[1]] = map[int]int{}
							}
						} else {
							resTable[cand[0]] = map[int]map[int]int{}
							resTable[cand[0]][cand[1]] = map[int]int{}
						}
						resTable[cand[0]][cand[1]][cand[2]] = cand[3]
						res = append(res, cand)
					}

				}
			}
		}
	}
	//time3 := time.Now()
	//fmt.Println("first loop :",time2.Sub(time1))
	//fmt.Println("second loop :",time3.Sub(time2))
	return res
}
