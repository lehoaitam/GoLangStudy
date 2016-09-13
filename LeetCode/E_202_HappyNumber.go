package LeetCode

func IsHappyNumber(n int) bool{
	slow := n
	fast := next(n)
	for slow != fast{
		slow = next(slow)
		fast = next(next(fast))
	}
	return fast == 1
}
func IsHappyNumber2(n int) bool{
	cnt := 1
	lim := 2
	slow := n
	fast := next(n)
	for slow != fast{
		if cnt == lim{
			cnt = 1
			lim *= 2
			slow = fast
		}else{
			cnt++
		}

		fast = next(fast)
	}
	return fast == 1
}
func next(n int) int{
	sum := 0
	for n > 0 {
		temp := n % 10
		sum += temp * temp
		n = n / 10
	}
	return sum
}