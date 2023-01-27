package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if os.Args[1] == "" {
		fmt.Print(0)
		return
	}
	args := os.Args[1]
	split := strings.Split(args, " ")

	fmt.Print(len(split))
}
