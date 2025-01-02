package main

import (
	"advent-of-code/common"
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	unsafeTreshold = 3
	unsafeReports  int
	safeReports    int
	reportsInFile  int
	reportResults  []string
)

func isReportSafe(report []int) bool {

	var direction string
	var problemDampened = false
	for i := 1; i < len(report); i++ {

		diff := report[i] - report[i-1]
		if direction == "" {
			if diff > 0 {
				direction = "inc"
			} else {
				direction = "dec"
			}
		}
		absDiff := math.Abs(float64(diff))
		if absDiff == 0 || diff > unsafeTreshold || diff < (-unsafeTreshold) || (direction == "inc" && diff < 0) || (direction == "dec" && diff > 0) {
			if problemDampened {
				return false
			}

			problemDampened = true
		}
	}

	return true
}

func analyzeReports() error {
	file, err := common.OpenPuzzleInput(nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		report := common.SplitString(fileScanner.Text())
		reportIsSafe := isReportSafe(common.ArrayAtoI(report))
		reportsInFile++

		if reportIsSafe {
			reportResults = append(reportResults, fmt.Sprintf("%s safe", report))
			safeReports++
		} else {
			reportResults = append(reportResults, fmt.Sprintf("%s unsafe", report))
			unsafeReports++
		}
	}

	return nil
}

func main() {
	start := time.Now()
	err := analyzeReports()

	if err != nil {
		fmt.Println(err)
		return
	}

	println(fmt.Sprintf("Total number of reports in file is: %d", reportsInFile))
	println(fmt.Sprintf("Total number of analyzed reports is: %d", unsafeReports+safeReports))
	println(fmt.Sprintf("Number of safe reports is: %d", safeReports))
	println(fmt.Sprintf("Number of unsafe reports is: %d", unsafeReports))

	mode := os.O_WRONLY | os.O_CREATE

	file, _ := os.OpenFile(filepath.Join(".", "report-results.txt"), mode, 0644)

	io.WriteString(file, strings.Join(reportResults, "\n"))
	file.Close()

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
