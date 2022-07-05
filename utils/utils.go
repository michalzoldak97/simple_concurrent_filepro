package utils

import (
	"fmt"
	"time"
)

func InitResultMap() map[int32]int {
	resMap := make(map[int32]int)

	for ch := 'a'; ch <= 'z'; ch++ {
		resMap[int32(ch)] = 0
	}
	return resMap
}

func PrintResMap(resMap map[int32]int) {
	for ch := 'a'; ch <= 'z'; ch++ {
		fmt.Printf("%s ==> %d   ", string(ch), resMap[int32(ch)])
	}
}

func ShowExecTime(name string, start time.Time) {
	fmt.Printf("Operation %s took %v\n", name, time.Since(start))
}
