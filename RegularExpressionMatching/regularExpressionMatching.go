package RegularExpressionMatching

// Assert p[pi]!='*'
func nextMatchPosition(s, p string, si, pi int) []int {
	res := []int{}
	if pi < len(p) {
		// Optionally match
		if pi+1 < len(p) && p[pi+1] == '*' {
			if s[si] == p[pi] || p[pi] == '.' {
				res = append(res, pi+1)
				if pi+2 < len(p) {
					res = append(res, nextMatchPosition(s, p, si, pi+2)...)
				}
			} else {
				if pi+2 < len(p) {
					res = append(res, nextMatchPosition(s, p, si, pi+2)...)
				}
			}
		} else {
			if p[pi] == s[si] || p[pi] == '.' {
				res = append(res, pi)
			}
		}
	}
	return res
}

func isMatch(s string, p string) bool {
	if len(s) == 0 {
		if len(p) == 0 {
			return true
		} else if len(p) == 1 {
			return false
		} else {
			if len(p)%2 == 1 {
				return false
			}
			idxInP := 0
			for ; idxInP < len(p); idxInP += 2 {
				if idxInP+1 < len(p) && p[idxInP+1] != '*' {
					return false
				}
			}
			return true
		}
	}
	if len(p) == 0 {
		return false
	}
	var matchedIdx []int

	matchedIdx = nextMatchPosition(s, p, 0, 0)

	for i := 1; i < len(s); i++ {
		tmpIdx := []int{}
		for _, idxInP := range matchedIdx {
			if p[idxInP] == '*' {
				// still match at the position of *
				if s[i] == s[i-1] || p[idxInP-1] == '.' {
					tmpIdx = append(tmpIdx, idxInP)
				}

				// match at next position, and it won't be *
				if idxInP+1 < len(p) {
					tmpIdx = append(tmpIdx, nextMatchPosition(s, p, i, idxInP+1)...)
				}
			} else {
				tmpIdx = append(tmpIdx, nextMatchPosition(s, p, i, idxInP+1)...)
			}
		}
		matchedIdx = tmpIdx
	}

	valid := []int{}

Check:
	for _, idxInP := range matchedIdx {
		i := idxInP
		for idxInP++; idxInP < len(p); idxInP += 2 {
			if idxInP+1 < len(p) && p[idxInP+1] != '*' {
				continue Check
			}
		}
		if idxInP == len(p) {
			valid = append(valid, i)
		}

	}
	return len(valid) > 0
}
