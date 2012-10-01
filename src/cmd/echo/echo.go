package main

import "os"

var Flagn = false

func Echo(str ...string) {
	for i, s := range str {
		if i > 0 { os.Stdout.WriteString(" ") }
		os.Stdout.WriteString(s)
	}

	if !Flagn { os.Stdout.WriteString("\n") }
}

func main() {
	i := 1

	if len(os.Args) > 1 && os.Args[1] == "-n" {
		Flagn, i = true, i+1
	}

	Echo(os.Args[i:]...)
	os.Exit(0)
}
