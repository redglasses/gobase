package main

import (
	g "../../getopt"
	"os"
)

var Flagu = false

func usage() {
	os.Stderr.WriteString("usage: cat [-u] [file...]\n")
	os.Exit(1)
}

func warn(err *[]error, e error) {
	os.Stderr.WriteString(e.Error()+"\n")
	*err = append(*err, e)
}

func Cat(file ...string) (err []error) {
	var (
		n int
		e error
		f *os.File
		buf = []byte{0}
	)

	if !Flagu {
		buf = make([]byte, 8192)
	}

	if len(file) == 0 {
		file = []string{"-"}
	}

	for i, _ := range file {
		if file[i] == "-" {
			f = os.Stdin
		} else if f, e = os.Open(file[i]); e != nil {
			warn(&err, e)
			continue
		}

		for n, e = f.Read(buf); n > 0; n, e = f.Read(buf) {
			if _, e = os.Stdout.Write(buf[:n]); e != nil {
				warn(&err, e)
			}
		}

		if file[i] != "-" {
			f.Close()
		}
	}

	return
}

func main() {
	parse: for {
		switch g.Getopt("u") {
			case g.EOF:
				break parse
			case 'u':
				Flagu = true
			default:
				usage()
		}
	}

	os.Exit(len(Cat(os.Args[g.Optind:]...)))
}
