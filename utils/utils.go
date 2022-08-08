package utils

import (
	"fmt"
	"time"
)

const F_PTH = "./data/sample.txt"

func InitResultMap() map[int32]int32 {
	resMap := make(map[int32]int32)

	for ch := 'a'; ch <= 'z'; ch++ {
		resMap[int32(ch)] = 0
	}
	return resMap
}

func PrintResMap(resMap map[int32]int32) {
	for ch := 'a'; ch <= 'z'; ch++ {
		fmt.Printf("%s ==> %d   ", string(ch), resMap[int32(ch)])
	}
}

func TransferRes(resMap map[int32]int32, a []int32) {
	i := 0
	for ch := 'a'; ch <= 'z'; ch++ {
		resMap[int32(ch)] = a[i]
		i++
	}
}

func ShowExecTime(name string, start time.Time) {
	fmt.Printf("Operation %s took %v\n", name, time.Since(start))
}
