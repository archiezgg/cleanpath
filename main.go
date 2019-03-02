package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var pFlag bool
var hFlag bool
var gFlag bool

func main() {
	parseFlags()
	if pFlag {
		printCleanPath()
		return
	}

	if gFlag {

	}

	if hFlag {
		welcome()
		return
	}

	welcome()
}

func parseFlags() {
	flag.BoolVar(&hFlag, "h", false, "Prints the help page.")
	flag.BoolVar(&gFlag, "g", false, "Shows your current PATH.")
	flag.BoolVar(&pFlag, "p", false, "Print what would be the result of the cleaning, but do nothing.")
	flag.Parse()
}

func welcome() {
	fmt.Println(`
Welcome to cleanpath, a lightweight Go CLI app to clean your path from duplicates!
	
Usage: cleanpath [OPTIONS]
Options:
	-h,		Show help page.
	-g,		Shows your current PATH.
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
