package main

import (
	"fmt"

	"./block"
	"./border"
	"./shiftand"
)

func main() {
	fmt.Println(border.Find("abab ababbaaba ab", "aba"))
	fmt.Println(block.Find("abab ababbaaba ab", "aba"))
	fmt.Println(shiftand.Find("abab ababbaaba ab", "aba"))
}
