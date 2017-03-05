package ZigZagConversion


func convert(s string, numRows int) string {
	var rowStr = make([]string,numRows,numRows)
	var res string = ""

	var i = 0
	length := len(s)

	if numRows==1{
		return s
	}


	for ;i<length ; i+=(numRows-1)*2 {
		j:=0
		for ; j+i <length && j<numRows;j++  {
			rowStr[j]+=string(s[i+j])
		}
		for ;j+i<length && j< 2*(numRows-1)  ; j++ {
			rowStr[2*(numRows-1)-j] += string(s[i+j])
		}
	}

	for i:=0;i<numRows ;i++  {
		res += rowStr[i]
	}
	return res


}
