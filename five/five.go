package five

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
	wGroup = sync.WaitGroup{}
	wg     = sync.WaitGroup{}
)

type kvp struct {
	k int32
	v int32
}

func updateRes(kvpChan chan kvp, charCount map[int32]int32) {

	for kvPair := range kvpChan {
		charCount[kvPair.k] += kvPair.v
	}

	wg.Done()
}

func calcChars(line string, kvpChan chan kvp) {
	res := utils.InitResultMap()
	for _, char := range line {
		res[char]++
	}

	for k, v := range res {
		kvpChan <- kvp{k, v}
	}

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

	defer utils.ShowExecTime("five", time.Now())

	var buff bytes.Buffer
	kvpChan := make(chan kvp)

	wg.Add(1)
	go updateRes(kvpChan, charCount)

	for scanner.Scan() {
		if len(buff.Bytes()) < 10000000 {
			buff.WriteString(strings.ToLower(scanner.Text()))
		} else {
			wGroup.Add(1)
			go calcChars(buff.String(), kvpChan)
			buff.Reset()
			buff.WriteString(strings.ToLower(scanner.Text()))
		}
	}

	if len(buff.Bytes()) > 0 {
		wGroup.Add(1)
		go calcChars(buff.String(), kvpChan)
	}

	wGroup.Wait()
	close(kvpChan)
	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	utils.PrintResMap(charCount)
}
