package two

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/michalzoldak97/simple_concurrent_filepro/utils"
)

var (
	charCount = utils.InitResultMap()
	wg        = sync.WaitGroup{}
)

func calcChars(inChan chan string) {
	for line := range inChan {
		for _, char := range line {
			charCount[char]++
		}
	}
	wg.Done()
}

func Run() {
	file, err := os.Open("./data/sample.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	defer utils.ShowExecTime("two", time.Now())

	linesChan := make(chan string)

	wg.Add(1)
	go calcChars(linesChan)

	for scanner.Scan() {
		linesChan <- strings.ToLower(scanner.Text())
	}

	close(linesChan)

	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	utils.PrintResMap(charCount)
}
