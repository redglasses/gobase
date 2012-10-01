package main

import (
	g "../../getopt"
	"os"
	"path/filepath"
)

var FlagP = true

func usage() {
	os.Stderr.WriteString("Usage: pwd [-L|-P]")
	os.Exit(1)
}

func warn(e error) {
	os.Stderr.WriteString(e.Error()+"\n")
}

func Pwd() (path string, err error) {
	if FlagP {
		if path, err = os.Getwd(); err != nil {
			warn(err)
		}

		if path, err = filepath.Abs(path); err != nil {
			warn(err)
		}
	} else {
		path = os.Getenv("PWD")
	}

	return
}

func main() {
	parse: for {
		switch g.Getopt("LP") {
			case g.EOF:
				break parse
			case 'L':
				FlagP = false
			case 'P':
				FlagP = true
			default:
				usage()
		}
	}

	if p, e := Pwd(); e == nil {
		os.Stdout.WriteString(p+"\n")
		os.Exit(0)
	}
	os.Exit(1)
}
