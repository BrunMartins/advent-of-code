// Day11
package main

import (
	"advent-of-code/common"
	"log"
	"os"
	"strings"
	"time"
)

var (
	puzzleInput *os.File
	newPassword string
)

func populateInput() string {
	scanner := common.GetInputLineScanner(puzzleInput)
	scanner.Scan()
	return scanner.Text()
}

// Function to check if the password contains an increasing straight of at least three letters
func hasIncreasingStraight(password string) bool {
	for i := 0; i < len(password)-2; i++ {
		if password[i]+1 == password[i+1] && password[i+1]+1 == password[i+2] {
			return true
		}
	}
	return false
}

// Function to check if the password contains any forbidden letters (i, o, l)
func hasNoForbiddenLetters(password string) bool {
	forbidden := []rune{'i', 'o', 'l'}
	for _, char := range forbidden {
		if strings.ContainsRune(password, char) {
			return false
		}
	}
	return true
}

// Function to check if the password contains at least two different, non-overlapping pairs of letters
func hasTwoNonOverlappingPairs(password string) bool {
	pairCount := 0
	i := 0
	for i < len(password)-1 {
		if password[i] == password[i+1] {
			pairCount++
			i += 2 // Skip the next character to avoid overlapping
		} else {
			i++
		}
	}
	return pairCount >= 2
}

// Function to verify if the password meets all the conditions
func isValidPassword(password string) bool {
	return hasIncreasingStraight(password) &&
		hasNoForbiddenLetters(password) &&
		hasTwoNonOverlappingPairs(password)
}

func computeNewPassword(password string) {
	runes := []rune(password)
	i := len(runes) - 1

	// Increment the password
	for i >= 0 {
		if runes[i] < 'z' {
			runes[i]++
			break
		} else {
			runes[i] = 'a' // Wrap around to 'a'
			i--            // Move to the next character to the left
		}
	}

	newPassword = string(runes)
}

func main() {
	start := time.Now()

	var err error
	puzzleInput, err = common.OpenPuzzleInput(nil)

	if err != nil {
		panic(err)
	}
	newPassword = populateInput()
	for !isValidPassword(newPassword) {
		computeNewPassword(newPassword)
	}

	println("Santa's new password is: ", newPassword)

	computeNewPassword(newPassword)
	for !isValidPassword(newPassword) {
		computeNewPassword(newPassword)
	}

	println("Santa's part 2 password is: ", newPassword)
	//Content here

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
