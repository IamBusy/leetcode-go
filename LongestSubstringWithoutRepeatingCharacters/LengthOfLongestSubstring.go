package LongestSubstringWithoutRepeatingCharacters

func lengthOfLongestSubstring(s string) int  {
	length := len(s)
	if length<=1{
		return length
	}

	var p1,p2,maxLength  = 0,0,1

	for ;p1<length ;p1++ {
		hashtable := make(map[uint8]int)
		hashtable[s[p1]]=1
		subLen:=1
		for p2=p1+1; p2<length ; p2++{
			if _,ok := hashtable[s[p2]];ok {
				if subLen > maxLength{
					maxLength = subLen
				}
				break
			}else if p2==length-1{

				if subLen+1>maxLength{
					maxLength = subLen+1
				}

			}
			subLen++
			hashtable[s[p2]]=1

		}
	}
	return maxLength

}