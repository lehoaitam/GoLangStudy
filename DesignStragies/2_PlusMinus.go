package HackerRank

import "fmt"

func Main2(){
	var arr []int32
	arr = input2()

	var posFrac, negFrac, zeroFrac float32
	posFrac, negFrac, zeroFrac = PlusMinus(arr)

	output2(posFrac, negFrac, zeroFrac)

}
func input2() ([]int32){
	//declare initialize variable
	var i,n int32
	var arr []int32
	//read n
	fmt.Scanln(&n)
	arr = make([]int32, n)
	//read n x n number of array
	for i = 0; i < n; i++{
		fmt.Scan(&arr[i])
	}
	return arr
}
func PlusMinus(arr []int32)(float32, float32, float32){
	//declare initial variable
	var i, n, posNum, negNum, zeroNum int32
	var posFrac, negFrac, zeroFrac float32
	n = int32(len(arr))
	posNum = 0
	negNum = 0
	zeroNum = 0

	//calculate diff of diagonal
	for i = 0; i < n;i++{
		if arr[i] > 0{
			posNum++
		}else if arr[i] < 0{
			negNum++
		}else {
			zeroNum++
		}
	}
	//calculate result
	posFrac = float32(posNum) / float32(n)
	negFrac = float32(negNum) / float32(n)
	zeroFrac = float32(zeroNum) / float32(n)
	return posFrac, negFrac, zeroFrac
}
func output2(posFrac float32, negFrac float32, zeroFrac float32){
	//output to STDOUT
	fmt.Printf("%.6f\n",posFrac)
	fmt.Printf("%.6f\n",negFrac)
	fmt.Printf("%.6f\n",zeroFrac)
}