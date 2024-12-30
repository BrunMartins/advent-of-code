package main

import (
	"advent-of-code/common"
	"bufio"
	"os"
	"strings"
)

var (
	puzzleInput    *os.File
	totalSqFtPaper = 0
	totalFtRibbon  = 0
)

func getInputLineScanner() *bufio.Scanner {
	fileScanner := bufio.NewScanner(puzzleInput)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func calculateAllDimensions(dim string) {
	dimensions := common.ArrayAtoI(strings.Split(dim, "x"))

	lxw := dimensions[0] * dimensions[1]
	wxh := dimensions[1] * dimensions[2]
	hxl := dimensions[2] * dimensions[0]

	minVal := common.Min(lxw, common.Min(wxh, hxl))

	params := make(map[string]interface{}, 3)

	params["lxw"] = lxw
	params["wxh"] = wxh
	params["hxl"] = hxl

	result := 2*lxw + 2*wxh + 2*hxl

	totalSqFtPaper += result + minVal
}

func calculateRibbonLength(dim string) {
	dimensions := common.ArrayAtoI(strings.Split(dim, "x"))

	l := dimensions[0]
	w := dimensions[1]
	h := dimensions[2]

	perimeter1 := 2 * (l + w)
	perimeter2 := 2 * (w + h)
	perimeter3 := 2 * (h + l)
	smallestPerimeter := common.Min(perimeter1, common.Min(perimeter2, perimeter3))

	volume := l * w * h

	result := smallestPerimeter + volume

	totalFtRibbon += result
}

func iteratePresents() {
	scanner := getInputLineScanner()

	for scanner.Scan() {
		text := scanner.Text()
		calculateAllDimensions(text)
		calculateRibbonLength(text)
	}
}

func main() {
	var err error
	puzzleInput, err = common.OpenPuzzleInput()

	if err != nil {
		panic(err)
	}

	iteratePresents()

	println("The total square feet needed for all presents is: ", totalSqFtPaper)
	println("The total feet of ribbon needed is: ", totalFtRibbon)
}
