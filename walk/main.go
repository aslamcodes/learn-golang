package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type config struct {
	root string
	list bool
	ext  string
	size int64
}

func main() {
	root := flag.String("root", "", "Root directory to start")
	list := flag.Bool("list", false, "list file only")
	ext := flag.String("ext", "", "file extension to filter out")
	size := flag.Int64("size", 0, "minimum file size")

	flag.Parse()

	c := config{
		root: *root,
		list: *list,
		ext:  *ext,
		size: *size,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(root string, out io.Writer, c config) error {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filterOut(path, c.ext, c.size, info) {
			return nil
		}

		return listFile(path, out)
	})
	return nil
}
