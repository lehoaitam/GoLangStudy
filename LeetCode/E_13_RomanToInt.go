package LeetCode

import "strings"

func RomanToInt(str string)int{
	result := 0
	romanStr := "IVXLCDM"
	arrInt := []int{1,5,10,50,100,500,1000}
	for i := len(str) - 1; i >= 0; i--{
		j := strings.IndexByte(romanStr,str[i])
		result += arrInt[j]
		if i > 0{
			k := strings.IndexByte(romanStr,str[i-1])
			if (k == 0 || k == 2 || k == 4) && (j - k <= 2 && j > k){
				result -= arrInt[k]
				i--
			}
		}
	}
	return result
}
