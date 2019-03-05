package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

var hFlag bool
var gFlag bool
var sFlag bool
var cFlag bool
var pFlag bool
var uniq []string
var duplicates []string

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
	if sFlag {
		printStatus()
		return
	}
	if pFlag {
		printCleanPath()
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
Welcome to cleanpath, a lightweight Go CLI application to show your duplicates in your PATH variable!
	
Usage: cleanpath [OPTIONS]
Options:
	-h,		Shows help page.
	-g,		Shows your current PATH.
	-s,		Prints status about the number of duplicates and the duplicates itself.
	-p,		Prints what would your clean path look like.
	`)
}

func printCurrentPath() {
	fmt.Printf("Your current path is:\n%s\n", os.Getenv("PATH"))
}

func getSplitPath() []string {
	f := func(r rune) bool {
		return r == rune(':')
	}

	s := strings.FieldsFunc(os.Getenv("PATH"), f)
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
	duplicates = nil
	uniq = nil
	s := getSplitPath()
	for _, v := range s {
		if ifContains(uniq, v) {
			duplicates = append(duplicates, v)
		} else {
			uniq = append(uniq, v)
		}
	}
	sort.Strings(duplicates)
}

func getDuplicateNumber() int {
	splitUniqAndDuplicates()
	return len(duplicates)
}

func printStatus() {
	splitUniqAndDuplicates()
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
	splitUniqAndDuplicates()
	newPath := ""
	for k, p := range uniq {
		if k == len(uniq)-1 {
			newPath += p
		} else {
			newPath += p + ":"
		}
	}
	return newPath
}

func printCleanPath() {
	if getDuplicateNumber() < 1 {
		fmt.Println("Your PATH is perfect as it is.")
	} else {
		newPath := createNewPath()
		fmt.Printf("Without duplicates, your path would look like this:\n%s\n", newPath)
	}
}
