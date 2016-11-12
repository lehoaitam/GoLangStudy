package LeetCode

func SingleNumber(nums []int)int{
	var result int
	if len(nums) == 1 {
		return nums[0]
	}else{
		result = nums[0]
		for i := 1; i < len(nums) ; i++{
			result = result ^ nums[i]
		}
	}
	return result

}
