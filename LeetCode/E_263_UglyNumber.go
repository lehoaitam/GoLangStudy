package LeetCode

import (
	"math/big"
	"fmt"
)

func IsUgly(num int) bool{
	if num < 1 {
		return false
	}
	if num == 1{
		return true
	}
	testNum := new(big.Int)
	testNum.SetString("3046798700052480000000000000",10)

	bNum := new(big.Int)
	bNum.SetInt64(int64(num))

	result := new(big.Int)
	result = result.Mod(testNum, bNum)
	fmt.Println("Mod:", result)

	zero := new(big.Int)
	zero.SetInt64(int64(0))

	fmt.Println("Zero:",zero)
	fmt.Println("result == zero:",result.Int64() == zero.Int64())
	return result.Int64() == zero.Int64()
}
