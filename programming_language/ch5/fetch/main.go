package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, url := range os.Args[1:] {
		filename, num, err := fetch(url)
		if err != nil {
			log.Fatalf("fetch error: %v", err)
		}
		fmt.Printf("%s(row: %d) is created.\n", filename, num)
	}
}

func fetch(url string) (fileName string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
