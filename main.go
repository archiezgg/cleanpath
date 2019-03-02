package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	p := flag.Bool("p", false, "Print what would be the result of the cleaning, but do nothing.")
	flag.Parse()
	if *p {
		printCleanPath()
		return
	}

	welcome()

}

func welcome() {
	fmt.Println(`

Welcome to cleanpath, a lightweight Go CLI app to clean your path from duplicates!
	
Usage: CLI Template [OPTIONS]
Options:
	-c,		Clean path of duplicates.
	-p,		Print what would be the result of the cleaning, but do nothing.
	`)
}

func printCleanPath() {
	p := os.Getenv("PATH")
	f := func(r rune) bool {
		return r == rune(':')
	}

	a := strings.FieldsFunc(p, f)
	fmt.Println(a)
}
