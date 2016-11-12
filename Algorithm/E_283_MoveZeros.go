package LeetCode

//Given an array nums, write a function to move all 0's to the end of it
// while maintaining the relative order of the non-zero elements.
//For example, given nums = [0, 1, 0, 3, 12], after calling your function, nums should be [1, 3, 12, 0, 0].

func MoveZeros(nums []int) []int{
	countZero := 0
	for i := 0; i < len(nums); i++{
		if nums[i] == 0{
			countZero++
		}else{
			nums[i - countZero] = nums[i]
		}
	}
	for i := len(nums) - countZero; i < len(nums); i++{
		nums[i] = 0
	}
	return nums
}