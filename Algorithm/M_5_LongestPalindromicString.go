package LeetCode


func LongestPalindromicString(s string) string{
	i create some error and commit!
	if s == "" || len(s) < 2{
		return s
	}
	sByte := []byte(s)
	n := len(sByte)
	t := make([]byte,n*2+1)
	var i int
	var j int
	i = 0
	j = 1
	t[0] = '#'
	for i < n{
		t[j] = sByte[i]
		j++
		t[j] = '#'
		j++
		i++
	}
	maxLeng := 0
	maxPos := 0
	n = 2*n +1
	p := make([]int,n)
	c := 1
	r := 1
	p[0] = 0
	p[1] = 1
	for i = 2; i < n ; i++{
		//calculate left position
		j = 2 * c - i

		//initialize p[i]
		if i < r{
			if p[j] < r - i{
				p[i] = p[j]
			}else{
				p[i] = r - i
			}
		}else{
			p[i] = 0
		}

		//expand to right
		for (i-1-p[i] >= 0 ) && (i+1+p[i] < n) &&  t[i-1-p[i]] == t[i+1+p[i]]{
			p[i]++
		}

		//update c and r
		if i + p[i] > r{
			r = i + p[i]
			c = i
		}
		if p[i] > maxLeng{
			maxLeng = p[i]
			maxPos = i
		}
	}

	//create return result
	n = maxLeng * 2 + 1
	resultBytes := make([]byte, n)
	j = maxPos - maxLeng
	i = 0
	k := 0
	for i < n{
		if t[j] != '#'{
			resultBytes[k] = t[j]
			k++
		}
		j++
		i++
	}
	strResult := string(resultBytes[:k])

	return strResult
}
