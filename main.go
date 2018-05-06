package main

import (
	"fmt"

	"./block"
	"./border"
)

func main() {
	fmt.Println(border.Find("abab ababbaaba ab", "aba"))
	fmt.Println(block.Find("abab ababbaaba ab", "aba"))
}
