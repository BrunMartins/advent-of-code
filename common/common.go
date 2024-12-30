package common

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func SplitString(data string) []string {
	splitString := strings.Fields(data)

	return splitString
}

func OpenPuzzleInput() (*os.File, error) {
	_, path, _, ok := runtime.Caller(1)

	if !ok {
		return nil, errors.New("unable to get caller information")
	}

	dir := filepath.Dir(path)
	file, err := os.Open(filepath.Join(dir, "puzzleinput.txt"))

	return file, err
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
