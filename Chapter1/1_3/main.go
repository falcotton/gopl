package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	echo1()
	echo2()
	echo3()
}

func echo1() {
	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	sec := time.Since(start).Seconds()
	fmt.Printf("echo1: %.2es \n", sec)
}

func echo2() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	sec := time.Since(start).Seconds()
	fmt.Printf("echo2: %.2es \n", sec)
}

func echo3() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	sec := time.Since(start).Seconds()
	fmt.Printf("echo3: %.2es \n", sec)

}
