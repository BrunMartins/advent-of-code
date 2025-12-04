// Day12
package main

import (
	"advent-of-code/common"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	puzzleInput *os.File
	input       string
	m           interface{}
)

// Function to recursively extract numbers from a JSON-like structure
func extractNumbers(data interface{}, exclude bool) []float64 {
	var numbers []float64

	switch v := data.(type) {
	case map[string]interface{}: // If the data is a map (JSON object)
		for _, value := range v {
			if value == "red" { // If a value is "red", exclude its children and siblings
				exclude = true
			}
			if !exclude {
				numbers = append(numbers, extractNumbers(value, exclude)...)
			}
			if value == "red" { // Reset exclude after processing the "red" subtree
				exclude = false
			}
		}
	case []interface{}: // If the data is a slice (JSON array)
		for _, item := range v {
			if item == "red" { // If an item is "red", exclude its siblings
				exclude = true
			}
			if !exclude {
				numbers = append(numbers, extractNumbers(item, exclude)...)
			}
			if item == "red" { // Reset exclude after processing the "red" subtree
				exclude = false
			}
		}
	case float64: // If the data is a number (JSON numbers are float64 in Go)
		if !exclude {
			numbers = append(numbers, v) // Add the number to the list
		}
	case int: // Handle integers (if they appear directly)
		if !exclude {
			numbers = append(numbers, float64(v))
		}
	}

	fmt.Println(numbers)

	return numbers
}

func main() {
	start := time.Now()

	var err error
	puzzleInput, err = common.OpenPuzzleInput(nil)

	if err != nil {
		panic(err)
	}

	input = common.GetWholeInputContent(puzzleInput)

	json.Unmarshal([]byte(input), &m)

	numbers := extractNumbers(m, false)

	result := 0.0

	for _, number := range numbers {
		result += number
	}

	println("The result is: ", int(result))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
