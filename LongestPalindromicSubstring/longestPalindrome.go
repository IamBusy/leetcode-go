package LongestPalindromicSubstring

import (
	"math"
)

func longestPalindrome(s string) string  {
	length := len(s)
	longestStr := ""
	for i:=0;i<2*length-1 ;i++  {

		//var oddStr string = string(s[i])
		//var evenStr string= ""
		//var leftCent int = i
		//var rightCent int = i
		var substr string
		var j,leftCent,rightCent ,maxJ int
		if i%2 == 0{
			substr = string(s[i/2])
			leftCent = i/2
			rightCent= i/2
			maxJ = int(math.Min(float64(length-i/2-1),float64(i/2)))
		}else{
			if s[i/2] != s[i/2+1]{
				continue
			}
			substr = string(s[i/2])+string(s[i/2+1])
			leftCent = i/2
			rightCent = i/2+1
			maxJ = int(math.Min(float64(length-i/2-2),float64(i/2)))
		}

		for j=1;j<=int(maxJ) ; j++{
			if s[leftCent-j] == s[rightCent+j]{
				substr = string(s[leftCent-j])+substr+string(s[rightCent+j])
			}else {
				break
			}
		}

		if len(substr)>len(longestStr){
			longestStr = substr
		}
	}
	return longestStr
}