package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"./block"
	"./border"
	"./kmp"
	"./shiftand"
)

func runTest(funcs map[string]func(string, string) []int, text string, pattern string) {
	var wg sync.WaitGroup

	for name, fn := range funcs {
		wg.Add(1)
		go func(fn func(string, string) []int, text, pattern, name string) {
			defer wg.Done()
			start := time.Now()
			result := fn(text, pattern)
			elapsed := time.Since(start)
			fmt.Printf("%s finished in: %s\n", name, elapsed.String())
			fmt.Printf("Occurrences found: %v\n\n", result)
		}(fn, text, pattern, name)
	}

	wg.Wait()
}

func main() {
	funcs := map[string]func(string, string) []int{
		"Blocks method":         block.Find,
		"KMP method":            kmp.Find,
		"KMP method (realtime)": kmp.FindRT,
		"Shift-And method":      shiftand.Find,
		"Borders method":        border.Find,
	}

	txtFlag := flag.String("t", "", "A text file")
	patFlag := flag.String("p", "", "A pattern file")
	flag.Parse()

	txtBuf, txtErr := ioutil.ReadFile(*txtFlag)
	if txtErr != nil {
		fmt.Println("An error occurred while reading text file")
		os.Exit(1)
	}

	patBuf, patErr := ioutil.ReadFile(*patFlag)
	if patErr != nil {
		fmt.Println("An error occurred while reading pattern file")
		os.Exit(1)
	}

	text := string(txtBuf)
	pattern := string(patBuf)

	runTest(funcs, text, pattern)
	fmt.Println()
}
