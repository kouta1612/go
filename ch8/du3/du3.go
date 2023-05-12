package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
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
	var n sync.WaitGroup
	for _, dir := range roots {
		n.Add(1)
		go workDir(dir, &n, filesizes)
	}
	go func() {
		n.Wait()
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

func workDir(dir string, n *sync.WaitGroup, filesizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go workDir(subdir, n, filesizes)
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
