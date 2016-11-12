package LeetCode


func ReverseVowels(s string) string{
	input := []byte(s)

	i := 0
	j := len(s) - 1;
	var ci, cj byte
	for i < j {
		ci = input[i]
		cj = input[j]
		bi := isVowel(ci)
		bj := isVowel(cj)
		if bi && bj{
			//reverse
			input[i] = cj
			input[j] = ci
			i++
			j--
		}
		if !bi{
			i++
		}
		if !bj{
			j--
		}
	}
	return string(input)
}

func isVowel(ch byte) bool{
	if ch == 'a' || ch == 'i' || ch == 'u' || ch == 'e' || ch == 'o' {
		return true
	}
	if ch == 'A' || ch == 'I' || ch == 'U' || ch == 'E' || ch == 'O' {
		return true
	}
	return false
}