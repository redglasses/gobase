package main

import (
	g "../../getopt"
	p "path"
	"os"
)

func usage() {
	os.Stderr.WriteString("usage: basename string [suffix]\n")
	os.Exit(1)
}
func Basename(path string, suffix string) string {
	path = p.Base(path)

	if len(suffix) > 0 && len(suffix) < len(path) &&
	   path[len(path)-len(suffix):] == suffix {
		return path[0:len(path)-len(suffix)]
	}
	return path
}

func main() {
	suffix := ""
	parse: for {
		switch g.Getopt("") {
			case g.EOF:
				break parse
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
