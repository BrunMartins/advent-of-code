// Day3
package main

import (
	"advent-of-code/common"
	"bufio"
	"fmt"
	"os"
)

type Coords struct {
	x int
	y int
}

var (
	puzzleInput     *os.File
	gpsInstructions []rune
	santaHouseGrid  = make(map[string]bool)
	santaPos        = Coords{0, 0}
	roboSantaPos    = Coords{0, 0}
)

func getInputLineScanner() *bufio.Scanner {
	fileScanner := bufio.NewScanner(puzzleInput)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func parseGPSInstructions() {
	scanner := getInputLineScanner()

	for scanner.Scan() {
		chars := []rune(scanner.Text())
		gpsInstructions = append(gpsInstructions, chars...)
	}
}

func distributePresentsAlone() {
	santaHouseGrid[fmt.Sprintf("%d,%d", 0, 0)] = true

	for _, instr := range gpsInstructions {
		executeMove(instr, &santaPos)
	}
}

func executeMove(instr rune, coords *Coords) {
	switch instr {
	case 60:
		coords.x--
	case 94:
		coords.y++
	case 62:
		coords.x++
	case 118:
		coords.y--
	}
	santaHouseGrid[fmt.Sprintf("%d,%d", coords.x, coords.y)] = true
}

func distributePresentsWithRoboSanta() {
	santaHouseGrid[fmt.Sprintf("%d,%d", 0, 0)] = true
	roboSantaTurn := false

	for _, instr := range gpsInstructions {
		if roboSantaTurn {
			executeMove(instr, &roboSantaPos)
		} else {
			executeMove(instr, &santaPos)
		}

		roboSantaTurn = !roboSantaTurn
	}
}

func resetPuzzle() {
	santaHouseGrid = make(map[string]bool)
	santaPos = Coords{0, 0}
	roboSantaPos = Coords{0, 0}
}

func main() {
	var err error
	puzzleInput, err = common.OpenPuzzleInput()

	if err != nil {
		panic(err)
	}
	defer puzzleInput.Close()

	parseGPSInstructions()
	distributePresentsAlone()
	println("Number of unique houses visited by Santa: ", len(santaHouseGrid))

	resetPuzzle()

	distributePresentsWithRoboSanta()
	println("Number of unique houses visited by Santa and Robo Santa: ", len(santaHouseGrid))

}
