package main

import (
	g "../../getopt"
	"errors"
	"io"
	"path"
	"os"
	"unicode"
	"unicode/utf8"
)

var (
	Flagf = false
	Flagi = false
	Flagr = false
)

func usage() {
	os.Stderr.WriteString("usage: rm [-f|-i] [-Rr] file...\n")
	os.Exit(1)
}

func warn(err *[]error, e ...error) {
	if !Flagf {
		for _, ee := range e {
			os.Stderr.WriteString(ee.Error()+"\n")
			*err = append(*err, ee)
		}
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
		if !unicode.IsSpace(r) { break }
	}
	/* right trim */
	for ; n > 0; n -= i {
		r, i = utf8.DecodeLastRune((*buf)[:n])
		if !unicode.IsSpace(r) { break }
	}

	if !(h < n) { return "" }

	/* convert to lower case */
	for i, r := range string((*buf)[h:n]) {
		if r >= 'A' && r <= 'Z' { r-=32 }
		utf8.EncodeRune((*buf)[h+i:n], r)
	}

	return string((*buf)[h:n])
}

func Rm(file ...string) (err []error) {
	var (
		e error
		f *os.File
		fi os.FileInfo
		names []string
	)

	for _, s := range file {
		if s = path.Clean(s); s == "." || s == ".." {
			continue
		}
		if fi, e = os.Lstat(s); e != nil {
			warn(&err, e)
			continue
		}

		switch {
			case fi.IsDir():
				if !Flagr {
					warn(&err, errors.New(s+": is a directory"))
					continue
				}
				if !Flagi {
					e = os.RemoveAll(s)
					break
				}

				for f, e = os.Open(s); e == nil; {
					warn(&err, Rm(names...)...)

					names, e = f.Readdirnames(512)
					for i, ss := range names {
						names[i] = s+string(os.PathSeparator)+ss
					}
				}

				if e == io.EOF {
					f.Close()
				}

				fallthrough
			default:
				if Flagi {
					buf := make([]byte, 32)
					switch interact("remove "+s+"?(y/n)[n]", &buf) {
						case "y", "yes":
						default:
							continue
					}
				}
				e = os.Remove(s)
		}

		if e != nil {
			warn(&err, e)
		}
	}

	return
}

func main() {
	parse: for {
		switch g.Getopt("fiRr") {
			case g.EOF:
				break parse
			case 'f':
				Flagf, Flagi = true, false
			case 'i':
				Flagi, Flagf = true, false
			case 'R':
				fallthrough
			case 'r':
				Flagr = true
			default:
				usage()
		}
	}

	if len(os.Args[g.Optind:]) == 0 { usage() }

	os.Exit(len(Rm(os.Args[g.Optind:]...)))
}
