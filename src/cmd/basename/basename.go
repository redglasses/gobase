package main

import (
	g "github.com/redglasses/gobase/src/getopt"
	"os"
	"path/filepath"
)

var Flagd = false

func usage() {
	os.Stderr.WriteString("usage: basename [-d] string [suffix]\n")
	os.Exit(1)
}

func Basename(name string, suffix string) string {
	fn := filepath.Base
	if Flagd { fn = filepath.Dir }

	name = fn(filepath.Clean(name))
	if slen, nlen := len(suffix), len(name);
	   slen > 0 && slen < nlen && name[nlen-slen:] == suffix {
		name = name[0 : nlen-slen]
	}

	return name
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
			os.Stdout.WriteString(Basename(os.Args[g.Optind],
			                               suffix)+"\n")
		default:
			usage()
	}

	os.Exit(0)
}
