package LeetCode

import "strconv"

//Write a function that takes an unsigned integer and returns the number of ’1' bits it has (also known as the Hamming weight).

func HammingWeight(n int64) int64{
	var result = n
	var mask1,_ = strconv.ParseInt("01010101010101010101010101010101", 2, 64)
	var mask2,_  = strconv.ParseInt("00110011001100110011001100110011", 2, 64)
	var mask3,_ = strconv.ParseInt("00001111000011110000111100001111", 2, 64)
	var mask4,_ = strconv.ParseInt("00000000111111110000000011111111", 2, 64)
	var mask5,_ = strconv.ParseInt("00000000000000001111111111111111", 2, 64)
	result = (result & mask1) + ((result >> 1) & mask1)
	result = (result & mask2) + ((result >> 2) & mask2)
	result = (result & mask3) + ((result >> 3) & mask3)
	result = (result & mask4) + ((result >> 4) & mask4)
	result = (result & mask5) + ((result >> 5) & mask5)
	return result
}