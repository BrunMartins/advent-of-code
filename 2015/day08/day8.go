// Day8
package main

import (
	"advent-of-code/common"
	"os"
)

var (
	puzzleInput *os.File
)

func calculateListSize() {}

func iteratePresentsList() {
	scanner := common.GetInputLineScanner(puzzleInput)

	for scanner.Scan() {

	}
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
