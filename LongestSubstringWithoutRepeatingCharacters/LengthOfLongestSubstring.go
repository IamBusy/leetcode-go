package LongestSubstringWithoutRepeatingCharacters

import "math"

func lengthOfLongestSubstring(s string) int  {
	length := len(s)
	if length<=1{
		return length
	}

	var p1,p2,maxLength  = 0,0,1

	hashtable := make(map[uint8]int)


	for ;p1<length ;p1++ {
		hashtable[s[p1]]=1

		p2 = int(math.Max(float64(p1+1),float64(p2)))
		subLen:= p2-p1

		for ; p2<length ; p2++{
			if _,ok := hashtable[s[p2]];ok {
				if subLen > maxLength{
					maxLength = subLen
				}
				break
			}else if p2==length-1{

				if subLen+1>maxLength{
					return int(math.Max(float64(maxLength),float64(subLen+1)))
				}

			}
			subLen++
			hashtable[s[p2]]=1

		}

		delete(hashtable,s[p1])
	}
	return maxLength

}