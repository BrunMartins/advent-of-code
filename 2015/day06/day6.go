// Day6
package main

import (
	"advent-of-code/common"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	puzzleInput           *os.File
	lightMatrix1          [][]bool
	lightMatrix2          [][]int
	rows                  = 1000
	cols                  = 1000
	lightsThatAreOn       = 0
	lightsTotalBrightness = 0
)

func getInputLineScanner() *bufio.Scanner {
	fileScanner := bufio.NewScanner(puzzleInput)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

func iterateLightCommands() {
	scanner := getInputLineScanner()

	for scanner.Scan() {
		lightCommand := scanner.Text()

		runLightCommand(lightCommand)
	}
}

func getCoords(coordString string) (int, int) {
	coords := common.ArrayAtoI(strings.Split(coordString, ","))
	return coords[0], coords[1]
}

func runLightCommand(command string) {
	commandParams := common.SplitString(command)
	fmt.Println(commandParams)

	// Define Coord Sets and Command
	var cmd int
	var startX, startY, endX, endY int

	if commandParams[1] == "on" {
		cmd = 1
	} else if strings.Contains(commandParams[1], ",") {
		cmd = 2
	}

	if cmd == 2 {
		startX, startY = getCoords(commandParams[1])
		endX, endY = getCoords(commandParams[3])
	} else {
		startX, startY = getCoords(commandParams[2])
		endX, endY = getCoords(commandParams[4])
	}

	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			if cmd == 2 {
				lightMatrix1[x][y] = !lightMatrix1[x][y]
				lightMatrix2[x][y] += cmd
			} else if cmd == 1 {
				lightMatrix2[x][y] += cmd
				lightMatrix1[x][y] = true
			} else {
				lightMatrix1[x][y] = false
				if lightMatrix2[x][y] > 0 {
					lightMatrix2[x][y] -= 1
				}
			}

		}
	}
}

func createMatrix[T any](init T) [][]T {
	matrix := make([][]T, rows)
	for i := range matrix {
		matrix[i] = make([]T, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = init
		}
	}

	return matrix
}

func calculateNumberOfLightsOn() {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if lightMatrix1[i][j] {
				lightsThatAreOn++
			}
		}
	}
}

func calculateLightsTotalBrightness() {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			lightsTotalBrightness += lightMatrix2[i][j]
		}
	}
}

func main() {
	var err error
	puzzleInput, err = common.OpenPuzzleInput()

	if err != nil {
		panic(err)
	}

	lightMatrix1 = createMatrix(false)
	lightMatrix2 = createMatrix(0)
	iterateLightCommands()
	calculateNumberOfLightsOn()
	calculateLightsTotalBrightness()

	println("There are", lightsThatAreOn, "lights on.")
	println("The total brightness of the lights is:", lightsTotalBrightness)
}
