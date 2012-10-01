package main

import (
	g "../../getopt"
	"path/filepath"
	"os"
)

func usage() {
	os.Stderr.WriteString("usage: dirname string\n")
	os.Exit(1)
}

func Dirname(path string) string {
	return filepath.Dir(path)
}

func main() {
	parse: for {
		switch g.Getopt("") {
			case g.EOF:
				break parse
			default:
				usage()
		}
	}

	if len(os.Args[g.Optind:]) == 1 {
		os.Stdout.WriteString(Dirname(os.Args[g.Optind])+"\n")
		os.Exit(0)
	}

	usage()
}
