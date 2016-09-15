package main

import (
	"fmt"
	"GoLangStudy/LeetCode"
)

func main() {
	//var n int64 = 10
	//fmt.Printf("number of bit 1 of %d : %d", n, LeetCode.HammingWeight(n))
	//fmt.Println("Anagram: ",LeetCode.IsAnagram("anagram","anmgraa"))
	var arr []int
	arr = make([]int, 5)
	arr[0] = 1
	arr[1] = 1
	arr[2] = 1
	arr[3] = 0
	arr[4] = 3
	fmt.Println("IsUgly: ",LeetCode.IsUgly(2))
}

