package main

import (
	"advent-of-code/common"
	"bufio"
	"fmt"
	"log"
	"time"
)

var (
	target                = "XMAS"
	target2               = "MAS"
	soup                  []string
	rows                  int
	cols                  int
	totalOccurencesOfXmas = 0
	totalOccurencesOfMasX = 0
)

func findTheXMAS() error {
	var err error
	file, err := common.OpenPuzzleInput(nil)

	if err != nil {
		return err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		soup = append(soup, fileScanner.Text())
	}

	rows = len(soup)
	cols = len(soup[0])

	countXmas(target)
	countXmas(target2)

	return nil
}

func countXmas(target string) {
	var diagOnlyF bool = false
	var diagOnlyT bool = true
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if soup[row][col] == 'X' {
				solveTheSoup(&target, row, col, &diagOnlyF)
			}
			if soup[row][col] == 'A' {
				solveTheSoup(&target, row, col, &diagOnlyT)
			}
		}
	}
}

func solveTheSoup(target *string, row int, col int, diagonalOnly *bool) int {
	if target == nil {
		return 0
	}

	if diagonalOnly == nil {
		*diagonalOnly = false
	}

	// println(string(*target))
	targetLen := len(string(*target)) - 1
	valueAtCoord := string(soup[row][col])

	if !*diagonalOnly {
		coordHorSubstr := valueAtCoord
		coordHorInvSubstr := valueAtCoord
		coordVerSubstr := valueAtCoord
		coordVerInvSubstr := valueAtCoord
		coordDiagDownRightSubstr := valueAtCoord
		coordDiagDownLeftSubstr := valueAtCoord
		coordDiagUpRightSubstr := valueAtCoord
		coordDiagUpLeftSubstr := valueAtCoord
		colNOOBRt := col+targetLen <= cols-1
		colNOOBLf := col-targetLen >= 0
		rowNOOBDn := row+targetLen <= rows-1
		rowNOOBUp := row-targetLen >= 0

		for i := 1; i <= targetLen; i++ {
			// println(col + i)
			// Not Out Of Bounds Checks
			if colNOOBRt {
				// Horizontal Checking
				coordHorSubstr += string(soup[row][col+i])

				if rowNOOBDn {
					// Diagonal Down Right Check
					coordDiagDownRightSubstr += string(soup[row+i][col+i])
				}
			}

			if colNOOBLf {
				// Horizontal Inverted Checking
				coordHorInvSubstr += string(soup[row][col-i])

				if rowNOOBUp {
					// Diagonal Up Left Check
					coordDiagUpLeftSubstr += string(soup[row-i][col-i])
				}
			}

			if rowNOOBDn {
				// Vertical Checking
				coordVerSubstr += string(soup[row+i][col])

				if colNOOBLf {
					coordDiagDownLeftSubstr += string(soup[row+i][col-i])
				}

			}
			if rowNOOBUp {
				// Vertical Inverted Checking
				coordVerInvSubstr += string(soup[row-i][col])
				if colNOOBRt {
					coordDiagUpRightSubstr += string(soup[row-i][col+i])
				}
			}
		}

		// Comparisons
		if !*diagonalOnly {

			if coordHorSubstr == *target {
				// println("Horizontal: ", coordHorSubstr)
				totalOccurencesOfXmas++
			}

			if coordHorInvSubstr == *target {
				// println("Horizontal Inv: ", coordHorInvSubstr)
				totalOccurencesOfXmas++
			}

			if coordVerSubstr == *target {
				// println("Vertical: ", coordVerSubstr)
				totalOccurencesOfXmas++
			}

			if coordVerInvSubstr == *target {
				// println("Vertical Inv: ", coordVerInvSubstr)
				totalOccurencesOfXmas++
			}

			if coordDiagDownRightSubstr == *target {
				// println("Diag Down Right: ", coordDiagDownRightSubstr)
				totalOccurencesOfXmas++
			}

			if coordDiagUpLeftSubstr == *target {
				// println("Diag Up Left: ", coordDiagUpLeftSubstr)
				totalOccurencesOfXmas++
			}

			if coordDiagDownLeftSubstr == *target {
				// println("Diag Down Left: ", coordDiagDownLeftSubstr)
				totalOccurencesOfXmas++
			}

			if coordDiagUpRightSubstr == *target {
				// println("Diag Up Right: ", coordDiagUpRightSubstr)
				totalOccurencesOfXmas++
			}
		}
	} else {
		// firstMasFound := false
		X := [][][][]int{
			{
				{
					{row - 1, col - 1}, {row, col}, {row + 1, col + 1},
				},
				{
					{row + 1, col - 1}, {row, col}, {row - 1, col + 1},
				},
			},
		}

		isFirstMatch := true
		for _, coords := range X {
			for _, coord := range coords {
				substr := ""
				// fmt.Println(coord)
				if coord[0][0] < 0 || coord[0][0] >= rows || coord[1][0] < 0 || coord[1][0] >= rows || coord[2][0] < 0 || coord[2][0] >= rows || coord[0][1] < 0 || coord[0][1] >= cols || coord[1][1] < 0 || coord[1][1] >= cols || coord[2][1] < 0 || coord[2][1] >= cols {
					break
				}
				substr = string(soup[coord[0][0]][coord[0][1]]) + string(soup[coord[1][0]][coord[1][1]]) + string(soup[coord[2][0]][coord[2][1]])

				if substr == "SAM" || substr == "MAS" {
					if isFirstMatch {
						isFirstMatch = false
						// fmt.Println(coord, " is the first match of MAS or SAM")
					} else {
						isFirstMatch = true
						totalOccurencesOfMasX++
						// fmt.Println(coords, " is a full match")
					}
				} else {
				}
			}
		}
	}

	return totalOccurencesOfXmas
}

func main() {
	start := time.Now()
	err := findTheXMAS()

	// t := "MAS"
	// d := true

	// solveTheSoup(&t, 1, 2, &d)

	if err != nil {
		fmt.Println(err)
		return
	}

	println(fmt.Sprintf("Total number of 'XMAS' found: %d", totalOccurencesOfXmas))
	println(fmt.Sprintf("Total number of 'MAS' Xs found: %d", totalOccurencesOfMasX))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
