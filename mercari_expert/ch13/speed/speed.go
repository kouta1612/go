package speed

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var maxsize int64 = 1e5

func uploadHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contentLength := r.ContentLength
		if contentLength > maxsize {
			contentLength = maxsize
		}
		_, err := io.CopyN(io.Discard, r.Body, contentLength)
		if err != nil {
			log.Printf("failed to write body: %s", err)
			return
		}
	}
}

func downloadHandler() http.HandlerFunc {
	src := rand.NewSource(0)
	return func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()
		size := queries.Get("size")
		max, err := strconv.Atoi(size)
		if err != nil {
			max = int(maxsize)
		}
		read := rand.New(src)
		_, err = io.CopyN(w, read, int64(max))
		if err != nil {
			log.Printf("failed to write random data: %s\n", err)
			return
		}
	}
}
