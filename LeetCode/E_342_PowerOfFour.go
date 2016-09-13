package LeetCode

func IsPowerOfFour(num int) bool {
	// 4^0 = 1, 4^2 = 1000, 4^3 = 100000000
	return (num > 0) && ((num & (num - 1) == 0) && (num & 0x55555555 == num))
}