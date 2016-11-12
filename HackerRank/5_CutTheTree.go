package HackerRank
//
//import (
//	"fmt"
//	"math"
//	//"sort"
//)
//
//type Edge struct{
//	e1 int32
//	e2 int32
//}
//func Main5(){
//	//var i,j,v int32
//	//var arrVertex []int32
//	//var arrEdgeLength int32
//	//var arrEdge []Edge
//	//arrVertex, arrEdge = input5()
//	//arrEdgeLength = int32(len(arrEdge))
//	//var m1 map[int32]int32
//	//var m2 map[int32]int32
//	//var minDiff, diff, diff1, diff2 int32
//	//var ok1, ok2 bool
//	//minDiff = int32(math.MaxInt32)
//	//
//	//toggle := 1
//	//for i = 0; i < arrEdgeLength; i++{
//	//
//	//	m1 = make(map[int32]int32)
//	//	m2 = make(map[int32]int32)
//	//	var arrEdgeFlag []int32
//	//	arrEdge = make([]int32, arrEdgeLength)
//	//	arrEdgeFlag = make([]bool, arrEdgeLength)
//	//	diff1 = 0
//	//	diff2 = 0
//	//	v1 := arrEdge[i].e1
//	//	v2 := arrEdge[i].e2
//	//
//	//	for j = 0; j < arrEdgeLength; j++{
//	//		//cut one edge at one time
//	//		if i != j && arrEdgeFlag[j] == false{
//	//			minDiff = diff
//	//	}
//	//
//	//}
//	//fmt.Println(minDiff)
//
//
//}
//}
//if diff1 < diff2 {
//diff = diff2 - diff1
//}else{
//diff = diff1 - diff2
//}
//if diff < minDiff{
//
//}
//
//func input5() ([]int32,[]Edge){
//	//declare initialize variable
//	var i,n int32
//	var arrVertex []int32
//	var arrEdge []Edge
//	//read n
//	fmt.Scanln(&n)
//	arrVertex = make([]int32, n + 1)
//	arrEdge = make([]Edge, n - 1)
//	//read array vertext and array edge number of array
//	for i = 1; i <= n; i++{
//		fmt.Scan(&arrVertex[i])
//	}
//	for i = 0; i < n-1; i++{
//		fmt.Scanf("%d %d",&arrEdge[i].e1, &arrEdge[i].e2)
//	}
//	return arrVertex,arrEdge
//}
//
