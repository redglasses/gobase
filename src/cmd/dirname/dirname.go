package main

import (
	"os"
	"path/filepath"
)

func usage() {
	os.Stderr.WriteString("usage: dirname name...")
	os.Exit(1)
}

func Dirname(name string) string {
	return filepath.Dir(filepath.Clean(name))
}

func main() {
	if len(os.Args) == 1 {
		usage()
	}

	for _, s := range os.Args[1:] {
		os.Stdout.WriteString(Dirname(s)+"\n")
	}

	os.Exit(0)
}
