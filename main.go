package main

import (
	"fmt"
	"os"
)

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

func printStatus() {
	splitUniqAndDuplicates()
	if getDuplicateNumber() < 1 {
		fmt.Println("You don't have any duplicates!")
	} else {
		fmt.Printf("You have %d number of duplicates in your path, which are the followings:\n\n", getDuplicateNumber())
		for _, v := range duplicates {
			fmt.Println(v)
		}
	}
}

func printCleanPath() {
	newPath := createNewPath()
	fmt.Println(newPath)
}
