package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var n int = 100000
	ex1 := sredn(experiment1, n)
	ex2 := sredn(experiment2, n)

	fmt.Println("\nTOTAL:")
	fmt.Println("Sredn of ex1:", ex1)
	fmt.Println("Sredn of ex2:", ex2)
}

func sredn(f func() int, n int) int {
	var sum int
	for i := 0; i < n; i++ {
		sum += f()
	}
	return sum / n
}

func experiment1() int {
	fmt.Println("Concatenation with +:")
	t1 := time.Now()
	var s string
	for i := 0; i < len(os.Args); i++ {
		s += os.Args[i] + "\n"
	}
	fmt.Println(s)
	t2 := time.Now()
	time_micro := t2.Sub(t1).Microseconds()
	fmt.Println("Microseconds:", time_micro)
	return int(time_micro)
}

func experiment2() int {
	fmt.Println("Concatenation with strings.Join:")
	t1 := time.Now()
	var s string
	//for i := 0; i < len(os.Args); i++ {
	//	s = strings.Join(strconv.Itoa(i), " ", os.Args[i], "\n")
	//}
	s = strings.Join(os.Args, "\n")
	fmt.Println(s)
	t2 := time.Now()
	time_micro := t2.Sub(t1).Microseconds()
	fmt.Println("Microseconds:", time_micro)
	return int(time_micro)
}
