package main

import (
	g "../../getopt"
	"os"
)

var (
	Flagp = false
	Flagm = os.ModePerm
)

func usage() {
	os.Stderr.WriteString("usage: mkdir [-p] dir...\n")
	os.Exit(1)
}

func Mkdir(path ...string) (err []error) {
	mk := os.Mkdir
	if Flagp { mk = os.MkdirAll }
	for _, s := range path {
		if e := mk(s, Flagm); e != nil {
			os.Stderr.WriteString(e.Error()+"\n")
			err = append(err, e)
		}
	}
	return
}

func main() {
	parse: for {
		switch g.Getopt("p") {
			case g.EOF:
				break parse
			case 'p':
				Flagp = true
			case 'm':
				/* TODO: implement [-m mode] option */
				fallthrough
			default:
				usage()
		}
	}

	if len(os.Args[g.Optind:]) == 0 { usage() }

	os.Exit(len(Mkdir(os.Args[g.Optind:]...)))
}
