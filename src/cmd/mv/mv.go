package main

import (
	g "../../getopt"
	"path"
	"os"
	"unicode"
	"unicode/utf8"
)

var Flagf = false
var Flagi = false

func usage() {
	os.Stderr.WriteString("usage:\tmv [-f|-i] source_file target_file\n"+
				"\tmv [-f|-i] source_file... target_dir\t")
	os.Exit(1)
}

func warn(err *[]error, e error) {
	if !Flagf {
		os.Stderr.WriteString(e.Error()+"\n")
		*err = append(*err, e)
	}
}

func interact(msg string, buf *[]byte) (ans string) {
	var r rune
	var i, n, h int

	os.Stderr.WriteString(msg)
	/* Read input */
	for n, _ = os.Stdin.Read(*buf); n==0; n, _ = os.Stdin.Read(*buf) {}

	/* left trim */
	for h = 0; h < n ; h += i {
		r, i = utf8.DecodeRune((*buf)[h:])
		if unicode.IsSpace(r) { break }
	}
	/* right trim */
	for ; n > 0; n -= i {
		r, i = utf8.DecodeLastRune((*buf)[:n])
		if unicode.IsSpace(r) { break }
	}

	if !(h < n) { return "" }

	/* convert to lower case */
	for i, r := range string((*buf)[h:n]) {
		if r >= 'A' && r <= 'Z' { r-=32 }
		utf8.EncodeRune((*buf)[h+i:n], r)
	}

	return string((*buf)[h:n])
}

func Mv(sources []string, target string) (err []error) {
	var (
		fi os.FileInfo
		e error
		getTarget = func(s string) string { return target }
	)

	if fi, e = os.Lstat(target); e == nil && fi.IsDir() &&
	(target[len(target)-1] == os.PathSeparator || len(sources) > 1) {
		getTarget = func(s string) string {
			return path.Clean(target)+string(os.PathSeparator)+path.Base(s)
		}
	}

	each: for _, s := range sources {
		t := getTarget(s)

		if _, e = os.Lstat(s); e != nil {
			warn(&err, e)
			continue
		}

		if _, e = os.Lstat(t); e == nil {
			if Flagf {
				e = os.Remove(t)
			} else if Flagi {
				buf := make([]byte, 32)
				switch interact("overwrite "+t+"?(y/n)[n]", &buf) {
					case "y", "yes":
						e = os.Remove(t)
					default:
						continue each
				}
			}

			if e != nil { warn(&err, e) }
		}

		if e = os.Rename(s, t); e != nil {
			warn(&err, e)
		}
	}

	return
}

func main() {
	parse: for {
		switch g.Getopt("fi") {
			case g.EOF:
				break parse
			case 'f':
				Flagf, Flagi = true, false
			case 'i':
				Flagi, Flagf = true, false
			default:
				usage()
		}
	}

	if len(os.Args[g.Optind:]) < 2 {
		usage()
	}

	os.Exit(len(Mv(os.Args[g.Optind:len(os.Args)-1],os.Args[len(os.Args)-1])))
}
