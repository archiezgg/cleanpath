package main

import (
	"flag"
)

var hFlag bool
var gFlag bool
var sFlag bool
var cFlag bool
var pFlag bool

func parseFlags() {
	flag.BoolVar(&hFlag, "h", false, "Prints the help page.")
	flag.BoolVar(&gFlag, "g", false, "Shows your current PATH.")
	flag.BoolVar(&sFlag, "s", false, "Prints status about the number of duplicates and the duplicates itself.")
	flag.BoolVar(&pFlag, "p", false, "Print what would be the result of the cleaning, but doesn't change path.")
	flag.Parse()
}
