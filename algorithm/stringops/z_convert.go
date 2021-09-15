package stringops

// PAYPALISHIRING
func ZConvert(s string, numRows int) string {
	sLen := len(s)
	if numRows == 1 || sLen <= numRows {
		return s
	}

	tmpArr := make([][]uint8, numRows)
	cursor := 0 // 0 ... numRows - 1
	isDown := false
	var res string

	for i := 0; i < sLen; i++ {
		tmpArr[cursor] = append(tmpArr[cursor], s[i])
		if cursor == 0 || cursor == numRows - 1 {
			isDown = !isDown
		}
		if isDown {
			cursor++
		} else {
			cursor--
		}
	}

	for i := 0; i < numRows; i++ {
		res += string(tmpArr[i])
	}
	return res
}
