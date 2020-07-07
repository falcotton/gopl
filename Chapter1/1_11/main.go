package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	var urls []string

	for _, arg := range os.Args[1:2] {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stdout, "read error:%s", arg)
		}
		urls = readList(f)
	}

	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f %7d %s", secs, nbytes, url)
}

func readList(f *os.File) []string {
	input := bufio.NewScanner(f)
	var urls []string
	for input.Scan() {
		urls = append(urls, input.Text())
	}
	return urls
}
