// Day3
package main

import (
	"advent-of-code/common"
	"bufio"
	"os"
)

var (
	puzzleInput *os.File
)

func getInputLineScanner() *bufio.Scanner {
	fileScanner := bufio.NewScanner(puzzleInput)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func main() {
	var err error
	puzzleInput, err = common.OpenPuzzleInput(nil)

	if err != nil {
		panic(err)
	}

	println(puzzleInput)

	// Content here
}
