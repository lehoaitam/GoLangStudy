package HackerRank

import "fmt"

func Main1(){
	var arr [][]int32
	arr = input1()

	var result int32
	result = diagonalDiff1(arr)

	output1(result)

}
func input1() ([][]int32){
	//declare initialize variable
	var i,j,n int32
	var arr [][]int32
	//read n
	fmt.Scanln(&n)
	arr = make([][]int32, n)
	//read n x n number of matrix
	for i = 0; i < n; i++{
		arr[i] = make([]int32,n)
		for j = 0; j < n - 1; j++{
			fmt.Scan(& arr[i][j])
		}
		fmt.Scanln(&arr[i][j])

	}
	return arr
}
func diagonalDiff1(arr [][]int32)int32{
	//declare initial variable
	var i, n, d1, d2, result int32
	n = int32(len(arr))
	d1 = 0
	d2 = 0
	result = 0
	//calculate diff of diagonal
	for i = 0; i < n;i++{
		d1 += arr[i][i]
		d2 += arr[n-1-i][i]
	}
	//calculate result
	result = d1 - d2
	if result < 0 {
		result *= -1
	}
	return result
}
func output1(v int32){
	//output to STDOUT
	fmt.Println(v)
}
