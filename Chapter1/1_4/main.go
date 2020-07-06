package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts,filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup3: %v \n", err)
				continue
			}
			countLines(f, counts,filenames)
			f.Close()
		}

	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t", n, line)
			for key,_ := range filenames[line]{
				fmt.Printf("%s\t",key)
			}
			fmt.Println()
		}
	}
}

func countLines(f *os.File, counts map[string]int,filenames map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if _,ok := filenames[input.Text()];ok{
			filenames[input.Text()][f.Name()]++
		}else{
			filenames[input.Text()] = make(map[string]int)
			filenames[input.Text()][f.Name()]++
		}
	}
}

