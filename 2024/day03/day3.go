package main

import (
	"advent-of-code/common"
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"
)

var (
	multiplicationRegex = regexp.MustCompile(`(don't\(\)|do\(\)|mul\((\d{1,3},\d{1,3})\))`)
	finalResult         = 0
	dontCounter         = 0
	enabled             = true
)

func multiplyArray(numbers []int) int {
	sum := 1
	for _, num := range numbers {
		sum *= num
	}
	return sum
}

func evaluateMultiplications(puzzleInput string) {
	matches := multiplicationRegex.FindAllStringSubmatch(puzzleInput, -1)
	for _, match := range matches {
		if match[1] == "do()" {
			enabled = true
		} else if match[1] == "don't()" {
			dontCounter++
			println(match[1])
			enabled = false
		} else if enabled {

			finalResult += multiplyArray(common.ArrayAtoI(strings.Split(match[2], ",")))
		}
	}
}

func fixComputer() error {

	file, err := common.OpenPuzzleInput(nil)
	if err != nil {
		return err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		evaluateMultiplications(fileScanner.Text())
	}

	println(dontCounter)

	return nil
}

func main() {
	start := time.Now()
	err := fixComputer()

	if err != nil {
		fmt.Println(err)
		return
	}

	println(fmt.Sprintf("Total sum of all the multiplications in the currupted file is: %d", finalResult))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
