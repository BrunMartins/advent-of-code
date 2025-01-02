package common

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type StringOrStringArray interface {
	string | []string
}

var (
	puzzleInput     = "puzzleInput.txt"
	defaultTestMode = false
	puzzleTestInput = "puzzle-test.txt"
	Truthy          = true
	Falsy           = false
)

func SplitString(data string) []string {
	splitString := strings.Fields(data)

	return splitString
}

func createPathToPuzzleInput(file string) (*string, error) {
	_, path, _, ok := runtime.Caller(2)

	if !ok {
		return nil, errors.New("unable to get caller information")
	}

	dir := filepath.Dir(path)
	finalPath := filepath.Join(dir, file)
	return &finalPath, nil
}

func OpenPuzzleInput(testMode *bool) (*os.File, error) {
	var file *string
	if testMode != &defaultTestMode {
		file, _ = createPathToPuzzleInput(puzzleTestInput)
	} else {
		file, _ = createPathToPuzzleInput(puzzleInput)
	}

	puzzleFile, err := os.Open(*file)

	return puzzleFile, err
}

func ArrayAtoI(report []string) []int {
	var report2 = []int{}

	for _, i := range report {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		report2 = append(report2, j)
	}
	return report2
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ReverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func StringToStringArray(str string) []string {
	strs := make([]string, len(str)) // Create a string slice of the same length
	for i, r := range str {
		strs[i] = string(r) // Convert each rune to a string
	}
	return strs
}

func StringArrayContains(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}
