package main

import (
	"os"
	"sort"
	"strings"
)

var uniq []string
var duplicates []string

func getSplitPath() []string {
	return strings.Split(os.Getenv("PATH"), ":")
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
