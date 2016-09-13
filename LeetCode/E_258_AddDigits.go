package LeetCode

//Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.
func AddDigits(n int) int{
	return 1 + ( n - 1 ) % 9
}
