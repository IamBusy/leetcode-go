package TwoSum


func twoSum(nums []int, target int) []int {
	length := len(nums)
	var m = make(map[int]int)
	for i:=0;i<length;i++ {
		if _,ok:=m[nums[i]];ok{
			if nums[i]*2==target{
				return []int{m[nums[i]],i}
			}
		}
		m[nums[i]]=i
	}
	for k,v:= range nums{
		if _,ok:=m[target-v];ok && v!=target-v{
			return []int{k,m[target-v]}
		}
	}
	return []int{}
}