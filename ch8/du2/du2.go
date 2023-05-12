package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

var varbose = flag.Bool("v", false, "show verbose progress messages")

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	filesizes := make(chan int64)
	go func() {
		for _, dir := range roots {
			workDir(dir, filesizes)
		}
		close(filesizes)
	}()

	var tick <-chan time.Time
	if *varbose {
		tick = time.Tick(100 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-filesizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files	%.1fGB\n", nfiles, float64(nbytes)/1e9)
}

func workDir(dir string, filesizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			workDir(subdir, filesizes)
		} else {
			fileinfo, _ := entry.Info()
			filesizes <- fileinfo.Size()
		}
	}
}

func dirents(dir string) []fs.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stdout, "du1: %v\n", err)
		return nil
	}
	return entries
}
