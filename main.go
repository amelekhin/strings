package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"./block"
	"./border"
	"./kmp"
	"./rabinkarp"
	"./shiftand"
)

type searchMethod struct {
	name string
	fn   func(string, string) []int
}

func runTest(funcs []searchMethod, text string, pattern string) {
	for _, searchMethod := range funcs {
		start := time.Now()
		result := searchMethod.fn(text, pattern)
		elapsed := time.Since(start)
		fmt.Printf("%s finished in: %s\n", searchMethod.name, elapsed.String())
		fmt.Printf("Occurrences found: %v\n\n", len(result))
	}
}

func main() {
	funcs := []searchMethod{
		{name: "Borders method", fn: border.Find},
		{name: "Blocks method", fn: block.Find},
		{name: "KMP method", fn: kmp.Find},
		{name: "KMP method (realtime)", fn: kmp.FindRT},
		{name: "Rabin-Karp", fn: rabinkarp.Find},
		{name: "Shift-And method", fn: shiftand.Find},
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

	border := strings.Repeat("=", 40)
	fmt.Printf("File: %s\n%s\n\n", *txtFlag, border)
	runTest(funcs, text, pattern)
}
