// Day8
package main

import (
	"advent-of-code/common"
	"os"
	"strconv"
	"strings"
)

var (
	puzzleInput               *os.File
	totalStringSizeInMemory   = 0
	totalStringSizeForEncoded = 0
)

func decodeEscapedString(input string) string {
	// Remove the surrounding double quotes
	unquoted := input[1 : len(input)-1]

	// Replace escape sequences
	unquoted = strings.ReplaceAll(unquoted, `\\`, `\`) // Replace escaped backslashes
	unquoted = strings.ReplaceAll(unquoted, `\"`, `"`) // Replace escaped double quotes

	// Handle hexadecimal escape sequences
	var result strings.Builder
	i := 0
	for i < len(unquoted) {
		if i+3 < len(unquoted) && unquoted[i] == '\\' && unquoted[i+1] == 'x' {
			// Check if the next two characters are valid hexadecimal
			hexPart := unquoted[i+2 : i+4]
			if _, err := strconv.ParseInt(hexPart, 16, 0); err == nil {
				// Convert the hex value to a character
				decodedChar, _ := strconv.ParseInt(hexPart, 16, 0)
				result.WriteByte(byte(decodedChar))
				i += 4 // Skip the `\xNN` sequence
				continue
			}
		}
		// If no valid hex sequence, copy the current character
		result.WriteByte(unquoted[i])
		i++
	}

	return result.String()
}

func encodeString(input string) string {
	var result strings.Builder

	// Add the surrounding double quotes
	result.WriteByte('"')

	for _, char := range input {
		switch char {
		case '\\':
			result.WriteString(`\\`) // Escape backslash
		case '"':
			result.WriteString(`\"`) // Escape double quote
		default:
			result.WriteRune(char) // Keep other characters as-is
		}
	}

	// Add the closing double quote
	result.WriteByte('"')

	return result.String()
}

func iteratePresentsList() {
	scanner := common.GetInputLineScanner(puzzleInput)

	for scanner.Scan() {
		str := scanner.Text()
		stringCodeSize, stringMemorySize := len(str), 0
		encodedStringSize := len(encodeString(str))

		stringMemorySize = len(decodeEscapedString(str))

		totalStringSizeForEncoded += (encodedStringSize - stringCodeSize)
		totalStringSizeInMemory += (stringCodeSize - stringMemorySize)
	}
}

func main() {
	var err error
	puzzleInput, err = common.OpenPuzzleInput(nil)

	if err != nil {
		panic(err)
	}

	iteratePresentsList()

	println("Part One: ", totalStringSizeInMemory)
	println("Part two: ", totalStringSizeForEncoded)
}
