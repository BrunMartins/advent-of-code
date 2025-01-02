package main

import (
	"advent-of-code/common"
	"bufio"
	"fmt"
)

type Trailhead struct {
	trailheadStart []int
	trailheadClass int
}

type Direction struct {
	row int
	col int
}

var (
	trailHeads           []Trailhead
	trails               []string
	rows                 int
	cols                 int
	sumOfClassifications = 0
)

func populateTrails() error {
	var err error
	file, err := common.OpenPuzzleInput(&common.Truthy)

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
	// println("Looking for trailheads")
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if string(trails[row][col]) == "0" {
				trailHead := Trailhead{[]int{row, col}, 0}
				// fmt.Println("Trailhead found: ", trailHead)
				trailHeads = append(trailHeads, trailHead)
			}
		}
	}
}

func calculateTrailheadClassifications() {
	for _, trailhead := range trailHeads {
		findEndOfTrail(0, trailhead.trailheadStart)
	}
}

func findEndOfTrail(val int, initialPos []int) {

	var directions []Direction

	if (initialPos[1] - 1) >= 0 {
		directions = append(directions, Direction{initialPos[0], initialPos[1] - 1})
	}

	if (initialPos[1] + 1) <= cols {
		directions = append(directions, Direction{initialPos[0], initialPos[1] + 1})
	}

	if (initialPos[0] + 1) <= rows {
		directions = append(directions, Direction{initialPos[0] + 1, initialPos[1]})
	}

	if (initialPos[0] - 1) >= 0 {
		directions = append(directions, Direction{initialPos[0] - 1, initialPos[1]})
	}

	if val < 9 {
		for _, direction := range directions {

			if (int(trails[direction.row][direction.col]) - '0') == 9 {
				fmt.Println("Position on end: ", initialPos)
				fmt.Println("Reached end of trail at: ", direction)
				sumOfClassifications++
				break
			}

			if ((int(trails[direction.row][direction.col]) - '0') - val) == 1 {
				// fmt.Println("Checking ", direction)
				findEndOfTrail(val+1, []int{direction.row, direction.col})
			}
		}
	}
}

func main() {
	populateTrails()
	findAllTrailheads()
	calculateTrailheadClassifications()
	fmt.Println("Sum of all trailhead classifications: ", sumOfClassifications)
}
