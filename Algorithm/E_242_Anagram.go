package LeetCode


//Given two strings s and t, write a function to determine if t is an anagram of s.

func IsAnagram(s string, t string) bool {
	if(s == "" && t == "") {
		return true
	} else if(s == "" || t == ""){
		return false
	} else if(len(s) != len(t)){
		return false
	} else{
		var m1 map[rune]int
		var m2 map[rune]int
		m1 = make(map[rune]int)
		m2 = make(map[rune]int)
		for _, ch := range s {
			m1[ch]++

		}
		for _, ch := range t {
			m2[ch]++
		}
		if mapEqual(m1,m2){
			return true
		}else{
			return false
		}
	}
}
func mapEqual(m1 map[rune]int, m2 map[rune]int) bool{
	//check if two map have same length
	if len(m1) != len(m2){
		return false
	}else{
		//go through first map and get value of same key in second
		for k1,v1 := range m1{
			v2 := m2[k1]
			if v1 != v2{
				return false
			}

		}
	}
	return true

}
