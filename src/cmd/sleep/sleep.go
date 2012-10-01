package main

import (
	"os"
	"time"
)

func usage() {
	os.Stderr.WriteString("usage: sleep time\n")
	os.Exit(1)
}

func main() {
	switch len(os.Args){
		case 2:
			for _, c := range os.Args[1] {
				if !(c >= '0' && c <= '9') {
					usage()
				}
			}

			d, _ := time.ParseDuration(os.Args[1]+"s");
			time.Sleep(d)
		default:
			usage()
	}
	os.Exit(0)
}
