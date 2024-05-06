package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

/*
Реализовать утилиту wget с возможностью скачивать сайты целиком.

*/

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "wrong input: use\t wget \"link to website\"")
		return
	}
	wget(os.Args[1])
}

func wget(link string) {
	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	strs := strings.Split(link, "/")

	var filename string
	if strs[len(strs)-1] == "" {
		filename = "index.html"
	} else {
		filename = strs[len(strs)-1]
	}

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
