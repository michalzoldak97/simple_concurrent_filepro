package four

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/michalzoldak97/simple_concurrent_filepro/utils"
)

const (
	numT     int = 12
	alphabet     = "abcdefghijklmnopqrstuvwxyz"
)

var (
	wg   = sync.WaitGroup{}
	freq [26]int32
)

func calcChars(inChan chan string) {
	for line := range inChan {
		for _, char := range line {
			idx := strings.Index(alphabet, string(char))
			if idx >= 0 {
				atomic.AddInt32(&freq[idx], 1)
			}
		}
	}
	wg.Done()
}

func Run() {
	charCount := utils.InitResultMap()

	file, err := os.Open(utils.F_PTH)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	defer utils.ShowExecTime("four", time.Now())

	inChan := make(chan string, 4096)

	for i := 0; i < numT; i++ {
		wg.Add(1)
		go calcChars(inChan)
	}

	for scanner.Scan() {
		inChan <- strings.ToLower(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	close(inChan)
	wg.Wait()
	utils.TransferRes(charCount, freq[:])
	utils.PrintResMap(charCount)
}
