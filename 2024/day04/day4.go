package main

import (
	"advent-of-code/common"
	"bufio"
	"fmt"
	"log"
	"time"
)

var (
	searchString    = "XMAS"
	totalOccurences = 0
)

func findTheXMAS() error {
	var err error
	file, err := common.OpenPuzzleInput()

	if err != nil {
		return err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	var letterSoup = []string{}
	for fileScanner.Scan() {
		letterSoup = append(letterSoup, fileScanner.Text())
	}

	totalOccurences += searchHorizontally(letterSoup, searchString)
	totalOccurences += searchVertically(letterSoup, searchString)
	totalOccurences += searchDiagonally(letterSoup, searchString)

	return nil
}

func searchHorizontally(soup []string, target string) int {
	var (
		rows           = len(soup)
		cols           = len(soup[0])
		targetInverted = common.ReverseString(target)
		targetLen      = len(target)
		occurrences    int
	)
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			substr := ""
			for i := 0; i < targetLen; i++ {
				if col+i < cols {
					substr += string(soup[row][col+i])
					continue
				}
				substr += string(soup[row][col])
			}
			if substr == target || substr == targetInverted {
				occurrences++
			}
		}
	}

	return occurrences
}
func searchVertically(soup []string, target string) int {
	var (
		rows           = len(soup)
		cols           = len(soup[0])
		targetInverted = common.ReverseString(target)
		targetLen      = len(target)
		occurrences    int
	)
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {

			substr := ""
			for i := 0; i < targetLen; i++ {
				if row+i < rows {
					substr += string(soup[row+i][col])
					continue
				}
				substr += string(soup[row][col])
			}
			if substr == target || substr == targetInverted {
				occurrences++
			}
		}
	}

	return occurrences
}

func searchDiagonally(soup []string, target string) int {
	var (
		rows           = len(soup)
		cols           = len(soup[0])
		targetInverted = common.ReverseString(target)
		targetLen      = len(target)
		occurrences    int
	)
	for row := 0; row <= rows-targetLen; row++ {
		for col := 0; col <= cols-targetLen; col++ {
			substr := ""
			if row+targetLen <= rows && col+targetLen <= cols {
				for i := 0; i < targetLen; i++ {
					if row+i < rows && col+1 < cols {
						substr += string(soup[row+i][col+i])
						continue
					} else if row+i < 0 && col+1 >= cols {
						substr += string(soup[row+i][col])
						continue
					} else if row+i <= rows && col+1 < cols {
						substr += string(soup[row][col+i])
						continue
					}

					substr += string(soup[row][col])
				}
				if substr == target || substr == targetInverted {
					occurrences++
				}
			}

			if substr == target || substr == targetInverted {
				occurrences++
			}

			if row-targetLen-1 >= 0 && col+targetLen <= cols {
				for i := 0; i < targetLen; i++ {
					if row-i > 0 && col+1 < cols {
						substr += string(soup[row-i][col+i])
						continue
					} else if row-i > 0 && col+1 >= cols {
						substr += string(soup[row-i][col])
						continue
					} else if row-i <= 0 && col+1 < cols {
						substr += string(soup[0][col+i])
						continue
					}

					substr += string(soup[row][col])
				}
				if substr == target || substr == targetInverted {
					occurrences++
				}
			}
		}
	}

	return occurrences
}

func main() {
	start := time.Now()
	err := findTheXMAS()

	if err != nil {
		fmt.Println(err)
		return
	}

	println(fmt.Sprintf("Total number of 'XMAS' found: %d", totalOccurences))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
