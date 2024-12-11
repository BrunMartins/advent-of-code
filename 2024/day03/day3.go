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
	multiplicationRegex = regexp.MustCompile(`(mul\((\d{1,3}\,\d{1,3})\)|do\(\)|don\'t\(\))`)
	finalResult         = 0
)

func multiplyArray(numbers []int) int {
	sum := 1
	for _, num := range numbers {
		sum *= num
	}
	return sum
}

func evaluateMultiplications(puzzleInput string) {
	matches := multiplicationRegex.FindAllStringSubmatch(puzzleInput, 30)
	enabled := true
	for _, match := range matches {
		if match[1] == "do()" {
			print("Sum ")
			enabled = true
			continue
		}

		if match[1] == "don't()" {
			print("Don't sum ")
			enabled = false
			continue
		}

		fmt.Println(match[0])
		if enabled {
			mulResult := multiplyArray(common.ArrayAtoI(strings.Split(match[2], ",")))
			finalResult += mulResult
		}
	}
}

func fixComputer() error {

	file, err := common.OpenPuzzleInput()
	if err != nil {
		return err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		evaluateMultiplications(fileScanner.Text())
	}

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
