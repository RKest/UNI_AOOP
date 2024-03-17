package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

const numCons = 5

var logger = log.New(os.Stdout, "", 0)

func main() {
	var wg sync.WaitGroup
	var globalWg sync.WaitGroup
	globalC := make(chan map[string]int)
	filepathChan := produce(&wg)
	for _ = range numCons {
		go consume(&wg, filepathChan, globalC)
	}
	globalWg.Add(1)
	go globalConsume(globalC, &globalWg)
	wg.Wait()
	close(globalC)
	globalWg.Wait()
}

func produce(wg *sync.WaitGroup) <-chan string {
	c := make(chan string)
	wd, _ := os.Getwd()
	wg.Add(1)
	go func() {
		_ = filepath.Walk(wd, func(path string, info fs.FileInfo, err error) error {
			if !info.IsDir() && strings.HasSuffix(path, ".txt") {
				wg.Add(1)
				c <- path
			}
			return nil
		})
		wg.Done()
		close(c)
	}()
	return c
}

func consume(wg *sync.WaitGroup, c <-chan string, outC chan map[string]int) {
	for s := range c {
		doConsume(wg, s, outC)
	}
}

func doConsume(wg *sync.WaitGroup, s string, outC chan map[string]int) {
	file, err := os.Open(s)
	if err != nil {
		return
	}
	defer file.Close()
	countAndPrint(wg, s, file, outC)
}

func countAndPrint(wg *sync.WaitGroup, filename string, r io.Reader, outC chan map[string]int) {
	wordsCount := make(map[string]int)
	s, _ := io.ReadAll(r)
	words := strings.Fields(string(s))
	for _, w := range words {
		if _, ok := wordsCount[w]; ok {
			wordsCount[w] += 1
		} else {
			wordsCount[w] = 1
		}
	}
	outC <- wordsCount
	countAndPrintInMap(filename, wordsCount)
	wg.Done()
}

func countAndPrintInMap(filename string, m map[string]int) {
	wc := make([]*KV, 0, len(m))
	for key, val := range m {
		wc = append(wc, &KV{key, val})
	}
	sort.Sort(ByVal(wc))
	formatted := make([]string, 0, 10)
	end := min(10, len(wc))
	for _, v := range wc[:end] {
		formatted = append(formatted, fmt.Sprintf("%s\t%d", v.key, v.value))
	}
	logger.Printf("Top %d most occurred words in file %s:\n%s", end, filename, strings.Join(formatted, "\n"))
}

func globalConsume(c <-chan map[string]int, wg *sync.WaitGroup) {
	gm := make(map[string]int)
	for m := range c {
		for k, v := range m {
			if _, ok := gm[k]; ok {
				gm[k] += v
			} else {
				gm[k] = v
			}
		}
	}
	countAndPrintInMap("All files", gm)
	wg.Done()
}

// KV Boiler plate to sort values in map
type KV struct {
	key   string
	value int
}

type ByVal []*KV

func (s ByVal) Len() int {
	return len(s)
}
func (s ByVal) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByVal) Less(i, j int) bool {
	return s[i].value > s[j].value
}
