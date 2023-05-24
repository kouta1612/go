package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/user"
	"path"
	"path/filepath"
)

func main() {
	// doFilePath()
	doHttpAndFilepath()
}

func doHttpAndFilepath() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if ok, err := path.Match("/data/*.txt", r.URL.Path); err != nil || !ok {
			http.NotFound(w, r)
			return
		}

		name := filepath.Join(cwd, "ch2", "data", filepath.Base(r.URL.Path))
		f, err := os.Open(name)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		defer f.Close()
		io.Copy(w, f)
	})
	http.ListenAndServe(":8080", nil)
}

func doFilePath() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dir := filepath.Join(u.HomeDir, ".config", "myapp")
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
