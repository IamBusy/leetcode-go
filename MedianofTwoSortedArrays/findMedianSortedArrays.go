package MedianofTwoSortedArrays


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	len1 := len(nums1)
	len2 := len(nums2)
	length := len1+len2

	//var merged [length]int
	var merged = make([]int,length,length)
	var p1,p2 int =0,0
	for p1< len1 && p2<len2  {
		if nums1[p1] < nums2[p2]{
			merged[p1+p2] = nums1[p1]
			p1++
		}else {
			merged[p1+p2] = nums2[p2]
			p2++
		}
	}
	for p1<len1  {
		merged[p1+p2] = nums1[p1]
		p1++
	}
	for p2<len2  {
		merged[p1+p2] = nums2[p2]
		p2++
	}

	if length%2 == 0{
		return float64(merged[length/2]+merged[length/2-1])/2.0
	}else{
		return float64(merged[length/2+1])
	}

}
