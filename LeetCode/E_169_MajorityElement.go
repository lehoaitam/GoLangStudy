package LeetCode

func MajorityElement(nums []int)int{
	mapNum := make(map[int]int)
	haftLenArr := len(nums) / 2
	for _,num := range nums{
		mapNum[num]++
		if mapNum[num] > haftLenArr{
			return num
		}
	}
	return 0
}