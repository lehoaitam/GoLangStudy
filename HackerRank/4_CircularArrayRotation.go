package HackerRank

import "fmt"

func Main4(){
	var i,n,k int32
	var arr []int32
	var qArr []int32

	n,k,arr,qArr = input4()

	for i = 0; i < int32(len(qArr)); i++{
		var index int32
		//check rotated index
		if qArr[i] - k < 0{
			//rotation index to first positoin
			index = (qArr[i] - k + n)
			for  index < 0{
				index = index + n
			}
		}else{
			index = qArr[i] - k
		}
		fmt.Println(arr[index])
	}

}

func input4() (int32,int32,[]int32,[]int32){
	//declare initialize variable
	var i,n,k,q int32
	var arr []int32
	var qArr []int32
	//read n
	fmt.Scanf("%v %v %v",&n,&k,&q)
	arr = make([]int32, n)
	qArr = make([]int32, q)
	//read n x n number of array
	for i = 0; i < n; i++{
		fmt.Scan(&arr[i])
	}
	for i = 0; i < q; i++{
		fmt.Scanln(&qArr[i])
	}
	return n,k,arr,qArr
}