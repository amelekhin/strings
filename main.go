package main

import (
	"fmt"

	"./border"
)

func main() {
	fmt.Println(border.Find("abab ababbaaba ab", "aba"))
}
