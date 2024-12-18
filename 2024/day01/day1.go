package main

import (
	"advent-of-code/common"
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"time"
)

var (
	leftList        []int
	rightList       []int
	distanceSum     int
	similarityScore int
)

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
			splitValues := common.SplitString(fileScanner.Text())
			var (
				leftVal  = splitValues[0]
				rightVal = splitValues[1]
			)
			appendToLeftList(leftVal)
			appendToRightList(rightVal)
		}
	}

	return nil
}

func main() {
	start := time.Now()
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
		counter := 0
		for j := 0; j < len(rightList); j++ {
			if leftList[i] == rightList[j] {
				counter++
			}
		}
		similarityScore += leftList[i] * counter
		counter = 0
	}

	println(distanceSum)
	println(similarityScore)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
