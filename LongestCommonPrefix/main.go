package LongestCommonPrefix

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	} else if length == 1 {
		return strs[0]
	}
	var substr string
	minLength := len(strs[0])

	for i := 0; i < length; i++ {
		if len(strs[i]) < minLength {
			minLength = len(strs[i])
		}
	}
Loop:
	for i := 0; i < minLength; i++ {
		for j := 1; j < length; j++ {
			if strs[j][i] != strs[0][i] {
				break Loop
			}
		}
		substr += string(strs[0][i])
	}
	return substr
}
