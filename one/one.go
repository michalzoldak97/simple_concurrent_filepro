package one

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/michalzoldak97/simple_concurrent_filepro/utils"
)

// to check if map key exists: if _, ok := charCount[char]; ok

func calcChars(charCount map[int32]int32, line string) {
	for _, char := range line {
		charCount[char]++
	}
}

func Run() {
	charCount := utils.InitResultMap()

	file, err := os.Open(utils.F_PTH)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	defer utils.ShowExecTime("one", time.Now())

	for scanner.Scan() {
		calcChars(
			charCount,
			strings.ToLower(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	utils.PrintResMap(charCount)
}
