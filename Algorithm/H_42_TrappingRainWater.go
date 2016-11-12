package LeetCode


func Trap(height []int) int {
	//empty arr
	if height == nil{
		return 0
	}
	//number of item less than 3
	length := len(height)
	if length < 3 {
		return 0
	}
	//1. find max of input arr
	max := height[0]
	for j1,_ := range height{
		if max < height[j1] {
			max = height[j1]
		}
	}

	//2. find list of max
	listMax := make([]int,0)
	for j2,_ := range height{
		if max == height[j2]{
			listMax = append(listMax, j2)
		}
	}
	listMaxLength := len(listMax)

	//3. iterate from 1 to first of listMax (listMax[0])

	result := 0
	minusValue := 0
	curValue := height[0]
	curPos := 0
	i := 1
	for ; i <= listMax[0] ; i++{

		if curValue > height[i] {
			minusValue += height[i]
		}else{
			//calculate trap of water
			result += (curValue * (i - curPos - 1) ) - minusValue
			curPos = i
			curValue = height[i]
			minusValue = 0
		}
	}

	//4. iterate from lenght - 1 to last of listMax (listMax[length-1)
	i = length - 2
	minusValue = 0
	curValue = height[length - 1]
	curPos = length - 1
	for ; i >= listMax[listMaxLength - 1] ; i--{

		if curValue > height[i] {
			minusValue += height[i]
		}else{
			//calculate trap of water
			result += (curValue * (curPos - i - 1) ) - minusValue
			curPos = i
			curValue = height[i]
			minusValue = 0
		}
	}

	//5. iterate from listMax[1] to listMax[length - 2]
	if listMaxLength > 1 {
		minusValue = 0
		for i = listMax[0] + 1; i <= listMax[listMaxLength-1] - 1; i++{
			minusValue += height[i]
		}
		result += max * (listMax[listMaxLength - 1] - listMax[0] - 1) - minusValue

	}

	return result
}