// Day5
package main

import (
	"advent-of-code/common"
	"bufio"
	"os"
	"regexp"
)

var (
	puzzleInput           *os.File
	niceWordCount1        int
	niceWordCount2        int
	vowels                = []rune{'a', 'e', 'i', 'o', 'u'}
	disallowedPairsRegexp = regexp.MustCompile(`ab|cd|pq|xy`)
)

func getInputLineScanner() *bufio.Scanner {
	fileScanner := bufio.NewScanner(puzzleInput)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func checkNiceString1(str string) {
	match := disallowedPairsRegexp.FindString(str)

	if match == "" {
		vowelCount := 0
		var prevChar rune
		containsPair := false
		for _, char := range str {
			for _, vowel := range vowels {
				if char == vowel {
					vowelCount++
				}
			}

			if !containsPair && prevChar != 0 {
				if prevChar == char {

					containsPair = true
				}
			}

			prevChar = char
		}

		if vowelCount >= 3 && containsPair {
			niceWordCount1++
		}
	}
}

func checkNiceString2(str string) {
	// Check for a pair of letters appearing at least twice without overlapping
	hasPair := false
	pairMap := make(map[string]int) // Map to store pairs and their positions

	for i := 0; i < len(str)-1; i++ {
		pair := str[i : i+2]
		if pos, exists := pairMap[pair]; exists {
			// Ensure the pairs do not overlap
			if i-pos >= 2 {
				hasPair = true
				break
			}
		} else {
			pairMap[pair] = i
		}
	}

	if !hasPair {
		return
	}

	// Check for a letter repeating with exactly one letter between them
	hasRepeat := false
	for i := 0; i < len(str)-2; i++ {
		if str[i] == str[i+2] {
			hasRepeat = true
			break
		}
	}

	if hasRepeat {
		niceWordCount2++
	}
}

func iterateStrings() {
	scanner := getInputLineScanner()

	for scanner.Scan() {
		text := scanner.Text()
		checkNiceString1(text)
		checkNiceString2(text)
	}
}

func main() {
	var err error
	puzzleInput, err = common.OpenPuzzleInput(nil)

	if err != nil {
		panic(err)
	}

	iterateStrings()

	println("There are", niceWordCount1, "nice words in this input.")
	println("There are", niceWordCount2, "new nice words in this input.")
	// Content here
}
