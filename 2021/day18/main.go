package main

import (
	"fmt"
	"os"

	"2021/common"
)

func main() {
	lines := common.ReadLines(os.Args[1])

	fmt.Println("Part I", lines)
}
