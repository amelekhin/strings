package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"./block"
	"./border"
	"./kmp"
	"./shiftand"
)

func runTest(funcs map[string]func(string, string) []int, text string, pattern string) {
	for name, fn := range funcs {
		start := time.Now()
		result := fn(text, pattern)
		elapsed := time.Since(start)
		fmt.Printf("%s finished in: %s\n", name, elapsed.String())
		fmt.Println(result)
	}
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

	txtBuf, errTxt := ioutil.ReadFile(*txtFlag)
	patBuf, errPat := ioutil.ReadFile(*patFlag)

	if errTxt != nil || errPat != nil {
		fmt.Println("An error occurred while reading files")
		os.Exit(1)
	}

	text := string(txtBuf)
	pattern := string(patBuf)

	runTest(funcs, text, pattern)
	fmt.Println()
}
