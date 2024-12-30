package main

import (
	"advent-of-code/common"
	"bufio"
)

type Trailhead struct {
	trailheadStart []int
	trailheadClass int
}

var (
	trailList []Trailhead
	trails    []string
	rows      int
	cols      int
)

func populateTrails() error {
	var err error
	file, err := common.OpenPuzzleInput()

	if err != nil {
		return err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		trails = append(trails, fileScanner.Text())
	}
	rows = len(trails)
	cols = len(trails[0])

	return nil
}

func findAllTrailheads() {
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if trails[row][col] == '0' {
				trailList = append(trailList, Trailhead{[]int{row, col}, 0})
			}
		}
	}
}

func main() {

}
