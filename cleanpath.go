package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var hFlag bool
var gFlag bool
var sFlag bool
var pFlag bool
var path = os.Getenv("PATH")
var uniq []string
var duplicates []string

func main() {
	parseFlags()
	splitUniqAndDuplicates()
	if hFlag {
		welcome()
		return
	}
	if gFlag {
		printCurrentPath()
		return
	}
	if sFlag {
		printStatus()
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
	flag.BoolVar(&sFlag, "s", false, "Prints status about the number of duplicates and the duplicates itself.")
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
	-s,		Prints status about the number of duplicates and the duplicates itself.
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

func splitUniqAndDuplicates() {
	s := getSplitPath()
	for _, v := range s {
		if ifContains(uniq, v) {
			duplicates = append(duplicates, v)
		} else {
			uniq = append(uniq, v)
		}
	}
}

func getDuplicateNumber() int {
	return len(duplicates)
}

func printStatus() {
	if getDuplicateNumber() < 1 {
		fmt.Println("You don't have any duplicates!")
	} else {
		fmt.Printf("You have %d number of duplicates in your path, which are the followings:\n", getDuplicateNumber())
		for _, v := range duplicates {
			fmt.Println(v)
		}
	}
}

func createNewPath() string {
	newPath := ""
	for _, p := range uniq {
		newPath += p + ":"
	}
	return newPath
}
