package LeetCode

func TitleToNumber(s string)int{
	if s == ""{
		return 0
	}else{
		result := 0
		base := 1
		for i := len(s) - 1; i >= 0; i--{
			temp := s[i] - 'A' + 1
			result += base * int(temp)
			base *= 26
		}
		return result
	}
}