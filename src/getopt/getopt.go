package getopt

import (
	"path"
	"os"
	"unicode/utf8"
)

const EOF = utf8.RuneError

var (
	Optarg string
	Opterr = true
	Optind = 0
	Optopt rune
	PosixlyCorrect = os.Getenv("POSIXLY_CORRECT") != ""

	done bool
	chrind, nonind, i, step int
)

func permuteOsArgs(head int, tail int) {
	t := os.Args[tail]
	for ; tail > head; tail-- {
		os.Args[tail] = os.Args[tail-1]
	}
	os.Args[head] = t
}

func warn(s string) {
	if Opterr {
		os.Stderr.WriteString(path.Base(os.Args[0])+": "+s+"\n")
	}
}

func Skip() {
	if !done {
		Optind, chrind = Optind+1, 1
	}
}

func Getopt(optstring string) (r rune) {
	if Optind == 0 {
		nonind, Optind, chrind, done = 0, 1, 1, false

		if len(optstring) == 0 {
			done = true
			return EOF
		}

		if optstring[0] == ':' {
			Opterr = false
		}
		/* os.Args will be permuted */
		if !PosixlyCorrect {
			nonind++
		}
	}

	if chrind == 1 {
		/* check for done */
		if done {
			return EOF
		} else if Optind >= len(os.Args) {
			done = true
			if nonind > 0 {
				Optind = nonind
			}
			return EOF
		} else if os.Args[Optind] == "--" {
			done, Optind = true, Optind+1
			return EOF
		}
		/* move to next option argument */
		for i = Optind ; i < len(os.Args); i++ {
			if os.Args[i][0] == '-' && os.Args[i] != "-" {
				break
			} else if nonind == 0 {
				done = true
				return EOF
			}
		}
		/* option argument not found */
		if i >= len(os.Args) {
			done = true
			return EOF
		}

		Optind = i
	}

	Optopt, _ = utf8.DecodeRuneInString(os.Args[Optind][chrind:])
	step, r, Optarg = 1, Optopt, string(Optopt)

	/* find opt in optstring */
	for i = 0; i < len(optstring); {
		o, n := utf8.DecodeRuneInString(optstring[i:])
		if o == Optopt {
			break
		}
		i += n
	}

	if i >= len(optstring) || Optopt == ':' {
		r = '?'
		warn("illegal option -- "+Optarg)
	} else if i < len(optstring) - 1 && optstring[i+1] == ':' {
		if chrind < len(os.Args[Optind]) - 1 {
			Optarg = os.Args[Optind][chrind+1:]
			chrind = len(os.Args[Optind])
		} else if Optind < len(os.Args) - 1 {
			Optarg = os.Args[Optind+1];
			step = 2
		} else {
			r = ':'
			warn("option requires an argument -- "+Optarg)
		}
	}

	if chrind += utf8.RuneLen(Optopt); chrind >= len(os.Args[Optind]) {
		chrind = 1
		Optind += step

		/* do permutation */
		for ; nonind > 0 && Optind > nonind && step > 0; step-- {
			permuteOsArgs(nonind, Optind-step)
			nonind++
		}
	}
	return
}
