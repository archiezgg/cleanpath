package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var hFlag bool
var gFlag bool
var fFlag bool
var pFlag bool
var path = os.Getenv("PATH")

func main() {
	parseFlags()
	if hFlag {
		welcome()
		return
	}
	if gFlag {
		printCurrentPath()
		return
	}
	if fFlag {
		fmt.Printf("You have %d duplicates in your path.\n", getDuplicateNumber())
		return
	}
	if pFlag {
		return
	}
	welcome()
}

func parseFlags() {
	flag.BoolVar(&hFlag, "h", false, "Prints the help page.")
	flag.BoolVar(&gFlag, "g", false, "Shows your current PATH.")
	flag.BoolVar(&fFlag, "f", false, "Prints the amount of duplicates found in your path.")
	flag.BoolVar(&pFlag, "p", false, "Print what would be the result of the cleaning, but doesn't change path.")
	flag.Parse()
}

func welcome() {
	fmt.Println(`
Welcome to cleanpath, a lightweight Go CLI app to clean your path from duplicates!
	
Usage: cleanpath [OPTIONS]
Options:
	-h,		Shows help page.
	-g,		Shows your current PATH.
	-f,		Prints the amount of duplicates found in your path.
	-c,		Cleans path of duplicates.
	-p,		Prints what would be the result of the cleaning, but doesn't change path.
	`)
}

func printCurrentPath() {
	fmt.Printf("Your current path is:\n %s\n", path)
}

func getSplitPath() []string {
	f := func(r rune) bool {
		return r == rune(':')
	}

	s := strings.FieldsFunc(path, f)
	return s
}

func ifContains(slice []string, word string) bool {
	for _, v := range slice {
		if v == word {
			return true
		}
	}
	return false
}

func getDuplicateNumber() int {
	s := getSplitPath()
	var tmp []string
	var duplicates []string

	for _, v := range s {
		if ifContains(tmp, v) {
			duplicates = append(duplicates, v)
		} else {
			tmp = append(tmp, v)
		}
	}

	return len(duplicates)
}
