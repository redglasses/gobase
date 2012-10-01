package main

import (
	g "../../getopt"
	"bufio"
	"os"
	"strconv"
)

var Flagn = uint64(10)

func usage() {
	os.Stderr.WriteString("usage: head [-n number] [file...]\n")
	os.Exit(1)
}

func warn(err *[]error, e error) {
	os.Stderr.WriteString(e.Error()+"\n")
	*err = append(*err, e)
}

func Head(file ...string) (err []error) {
	var (
		f *os.File
		e error
	)
	for _, s := range file {
		if s == "-" {
			f = os.Stdin
		} else if f, e = os.Open(s); e != nil {
			warn(&err, e)
			continue
		}

		for i, r, e := uint64(0), bufio.NewReader(f), error(nil);
		    i < Flagn && e == nil; i++ {
			s, e = r.ReadString('\n')
			os.Stdout.WriteString(s)
		}

		if s != "-" {
			f.Close()
		}
	}

	return
}

func main() {
	exitCode := 0
	defer os.Exit(exitCode)

	parse: for {
		switch g.Getopt("n:") {
			case g.EOF:
				break parse
			case 'n':
				if i, e := strconv.ParseUint(g.Optarg, 10, 64); e != nil {
					usage()
				} else {
					Flagn = i
				}
			default:
				usage()
		}
	}

	if len(os.Args[g.Optind:]) == 0 {
		exitCode = len(Head("-"))
	} else {
		exitCode = len(Head(os.Args[g.Optind:]...))
	}
}
