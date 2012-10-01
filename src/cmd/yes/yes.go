package main

import (
	"os"
	"strings"
)

func Yes(s ...string) {
	msg := "yes"

	if len(s) > 0 {
		msg = strings.Join(s, " ")
	}

	for  {
		os.Stdout.WriteString(msg+"\n")
	}
}

func main() {
	Yes(os.Args[1:]...)
}
