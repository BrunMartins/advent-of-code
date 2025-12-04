//Day24
package main

import (
	"advent-of-code/common"
	"bufio"
	"os"
	"time"
	"log"
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
	start := time.Now()

	var err error
	puzzleInput, err = common.OpenPuzzleInput(nil)

	if err != nil {
		panic(err)
	}

	println(puzzleInput)

	//Content here

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}