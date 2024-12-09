package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"
)

var (
	leftList    []int
	rightList   []int
	distanceSum int
)

func splitString(data string) (string, string) {
	splitString := strings.Fields(data)

	return splitString[0], splitString[1]
}

func appendToLeftList(value string) error {
	leftInt, err := strconv.Atoi(value)

	if err != nil {
		return err
	}

	leftList = append(leftList, leftInt)

	return nil
}

func appendToRightList(value string) error {
	rightInt, err := strconv.Atoi(value)

	if err != nil {
		fmt.Println(err)
		return err
	}

	rightList = append(rightList, rightInt)

	return nil
}

func buildLists() error {
	_, path, _, ok := runtime.Caller(0)

	if !ok {
		return errors.New("unable to get caller information")
	}

	dir := filepath.Dir(path)

	if readFile, err := os.Open(filepath.Join(dir, "puzzleinput.txt")); err != nil {
		fmt.Println(err)
		return err
	} else {
		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			leftVal, rightVal := splitString(fileScanner.Text())
			appendToLeftList(leftVal)
			appendToRightList(rightVal)
		}
	}

	return nil
}

func main() {
	err := buildLists()

	if err != nil {
		fmt.Println(err)
		return
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	for i := 0; i < len(leftList); i++ {
		var diff = int(math.Abs(float64(leftList[i]) - float64(rightList[i])))
		distanceSum += diff
	}

	println(distanceSum)
}
