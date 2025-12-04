// Day10
package main

import (
	"advent-of-code/common"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	puzzleInput  *os.File
	input        string
	output       strings.Builder
	part1Times   = 40
	part2Times   = 50
	finalContent string
)

func populateInput() {
	scanner := common.GetInputLineScanner(puzzleInput)
	scanner.Scan()
	input += scanner.Text()
}

func compressContent(in string) {
	counter := 1
	for i := 0; i < len(in); i++ {
		if i+1 < len(in) {
			if in[i+1] == in[i] {
				counter++
				continue
			} else {
				output.WriteString(fmt.Sprintf("%d%c", counter, in[i]))
				counter = 1
			}
		} else {
			output.WriteString(fmt.Sprintf("%d%c", counter, in[i]))
		}
	}

	finalContent = output.String()
	output.Reset()
}

func part1() {
	finalContent = input
	for i := 0; i < part1Times; i++ {
		compressContent(finalContent)
	}
	fmt.Println("Length of part1 result: ", len(finalContent))
}

func part2() {
	finalContent = input
	for i := 0; i < part2Times; i++ {
		compressContent(finalContent)
	}
	fmt.Println("Length of part2 result: ", len(finalContent))
}

func main() {
	start := time.Now()

	var err error
	puzzleInput, err = common.OpenPuzzleInput(nil)

	if err != nil {
		panic(err)
	}

	populateInput()
	part1()
	part2()

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
