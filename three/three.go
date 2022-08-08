package three

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/michalzoldak97/simple_concurrent_filepro/utils"
)

var (
	rWLock = sync.Mutex{}
	wGroup = sync.WaitGroup{}
)

func calcChars(charCount map[int32]int32, line string) {
	res := utils.InitResultMap()
	for _, char := range line {
		res[char]++
	}
	rWLock.Lock()
	for k, v := range res {
		charCount[k] += v
	}
	rWLock.Unlock()
	wGroup.Done()
}

func Run() {
	charCount := utils.InitResultMap()

	file, err := os.Open(utils.F_PTH)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	defer utils.ShowExecTime("three", time.Now())

	var buff bytes.Buffer

	for scanner.Scan() {
		if len(buff.Bytes()) < 9000000 {
			buff.WriteString(strings.ToLower(scanner.Text()))
		} else {
			wGroup.Add(1)
			go calcChars(charCount, buff.String())
			buff.Reset()
			buff.WriteString(strings.ToLower(scanner.Text()))
		}
	}

	if len(buff.Bytes()) > 0 {
		wGroup.Add(1)
		go calcChars(charCount, buff.String())
	}

	wGroup.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	utils.PrintResMap(charCount)
}
