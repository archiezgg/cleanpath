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
var path = os.Getenv("PATH")

func main() {
	parseFlags()
	if gFlag {
		getCurrentPath()
		return
	}
	if pFlag {
		printCleanPath()
		return
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
	-h,		Shows help page.
	-g,		Shows your current PATH.
	-c,		Cleans path of duplicates.
	-p,		Prints what would be the result of the cleaning, but do nothing.
	`)
}

func getCurrentPath() {
	fmt.Printf("Your current path is:\n %s\n", path)
}

func printCleanPath() {
	f := func(r rune) bool {
		return r == rune(':')
	}

	a := strings.FieldsFunc(path, f)
	fmt.Println(a)
}
