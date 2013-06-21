package main

import (
	g "github.com/redglasses/gobase/src/getopt"
	"os"
	"path/filepath"
)

var Flagd string

func usage() {
	os.Stderr.WriteString("usage: cleanname [-d pwd] name...\n")
	os.Exit(1)
}

func Cleanname(name string) string {
	name = filepath.Clean(name)

	if len(Flagd) > 0 && filepath.VolumeName(name) == "" {
		name = filepath.Clean(Flagd) + string(os.PathSeparator) + name
	}

	return name
}

func main() {
	if len(os.Args) == 1 { usage() }
	parse: for {
		switch g.Getopt("d:") {
		case g.EOF:
			break parse
		case 'd':
			Flagd = g.Optarg
		default:
			usage()
		}
	}

	if len(os.Args[g.Optind:]) == 0 {
		usage()
	}

	for _, n := range os.Args[g.Optind:] {
		os.Stdout.WriteString(Cleanname(n)+"\n")
	}

	os.Exit(0)
}
