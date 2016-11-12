package LeetCode
func ReverseString(s string) string {
	var arrRune []rune
	arrRune = make([]rune,len(s))
	i := len(s) - 1
	for _, r := range s{
		arrRune[i] = r
		i--
	}
	return string(arrRune)
}