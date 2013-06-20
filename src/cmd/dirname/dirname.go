package main

import (
	"os"
	"path"
	"strings"
)

func usage() {
	os.Stderr.WriteString("usage: dirname name...")
	os.Exit(1)
}

func Dirname(name string) string {
	name = path.Dir(name)

	if '/' == os.PathSeparator {
		return name
	}

	return strings.Replace(name, "/", string(os.PathSeparator), -1)
}

func main() {
	if len(os.Args) == 1 {
		usage()
	}

	for _, s := range os.Args[1:] {
		os.Stdout.WriteString(s+"\n")
	}

	os.Exit(0)
}
