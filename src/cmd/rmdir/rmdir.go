package main

import (
	g "../../getopt"
	d "../dirname"
	"errors"
	"os"
	"path"
)

var Flagp = false

func usage() {
	os.Stderr.WriteString("usage: rmdir [-p] dir...\n")
	os.Exit(1)
}

func warn(err *[]error, e ...error) {
	for _, ee := range e {
		os.Stderr.WriteString(ee.Error()+"\n")
	}
	*err = append(*err, e...)
}

func Rmdir(dir ...string) (err []error) {
	var (
		fi os.FileInfo
		e error
	)

	each: for _, s := range dir {
		s = path.Clean(s)

		fi, e = os.Lstat(s)
		switch {
			case e == nil && !fi.IsDir():
				e = errors.New(s+" is not a directory")
				fallthrough
			case e != nil:
				warn(&err, e)
				continue
			default:
				for Flagp && d.Dirname(s) != "." {
					if e := os.Remove(s); e != nil {
						warn(&err, e)
						continue each
					}
					s = d.Dirname(s)
				}

				if e = os.Remove(s); e != nil {
					warn(&err, e)
				}
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
			default:
				usage()
		}
	}

	if len(os.Args[g.Optind:]) < 1 {
		usage()
	}

	os.Exit(len(Rmdir(os.Args[g.Optind:]...)))
}
