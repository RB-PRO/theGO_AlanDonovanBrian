package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s string
	var sep string = "\n"
	for i := 0; i < len(os.Args); i++ {
		s += strconv.Itoa(i) + " " + os.Args[i] + sep
	}
	fmt.Println(s)
}
