package main

import (
	g "github.com/redglasses/gobase/src/getopt"
	p "path"
	"os"
)

var Flagd = false

func usage() {
	os.Stderr.WriteString("usage: basename [-d] string [suffix]\n")
	os.Exit(1)
}

func Basename(path string, suffix string) string {
	if Flagd {
		path = p.Dir(path)
	} else {
		path = p.Base(path)
	}

	if len(suffix) > 0 && len(suffix) < len(path) &&
	   path[len(path)-len(suffix):] == suffix {
		return path[0:len(path)-len(suffix)]
	}
	return path
}

func main() {
	suffix := ""
	parse: for {
		switch g.Getopt("d") {
			case g.EOF:
				break parse
			case 'd':
				Flagd = true
			default:
				usage()
		}
	}

	switch len(os.Args[g.Optind:]) {
		case 2:
			suffix = os.Args[g.Optind+1]
			fallthrough
		case 1:
			os.Stdout.WriteString(Basename(os.Args[g.Optind],suffix)+"\n")
		default:
			usage()
	}

	os.Exit(0)
}
