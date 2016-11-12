package LeetCode

func Rob(nums []int) int {
	last, now := 0, 0

	for _, num := range nums{
		last, now = now, max(last + num, now)


	}
	return now;
}
func max(num1 int, num2 int) int{
	if num1 > num2{
		return num1
	}else{
		return num2
	}
}