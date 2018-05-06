package main

import (
	"fmt"

	"./block"
	"./border"
	"./kmp"
	"./shiftand"
)

func main() {
	fmt.Println(border.Find("abab ababbaaba ab", "aba"))
	fmt.Println(block.Find("abab ababbaaba ab", "aba"))
	fmt.Println(kmp.Find("abab ababbaaba ab", "aba"))
	fmt.Println(kmp.FindRT("abab ababbaaba ab", "aba"))
	fmt.Println(shiftand.Find("abab ababbaaba ab", "aba"))
}
