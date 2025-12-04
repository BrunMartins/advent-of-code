package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	day := flag.Int("day", 0, "Puzzle day (1-25)")
	year := flag.Int("year", 0, "Puzzle year")
	flag.Parse()

	today := time.Now()
	currentYear := today.Year()
	currentDay := today.Day()

	if *year != 0 {
		currentYear = *year
	}
	if *day != 0 {
		currentDay = *day
	}

	// Validate inputs
	if currentDay < 1 || currentDay > 25 {
		fmt.Printf("Error: Day must be between 1 and 25, got %d\n", currentDay)
		os.Exit(1)
	}

	if currentYear < 2015 {
		fmt.Printf("Error: Year must be 2015 or later, got %d\n", currentYear)
		os.Exit(1)
	}

	// Build the path to the Go file (relative to root directory)
	dayFolder := fmt.Sprintf("day%02d", currentDay)
	yearFolder := fmt.Sprintf("%d", currentYear)
	goFile := fmt.Sprintf("day%d.go", currentDay)
	filePath := filepath.Join("..", "..", yearFolder, dayFolder, goFile)
	dir := filepath.Join("..", "..", yearFolder, dayFolder)

	// Check if the file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("Error: Go file does not exist at %s\n", filePath)
		os.Exit(1)
	}

	fmt.Printf("Running Day %d, Year %d (Go)\n", currentDay, currentYear)
	fmt.Printf("File: %s\n", filePath)
	fmt.Printf("Directory: %s\n", dir)
	fmt.Println("=" + fmt.Sprintf("%*s", 50, "") + "=")

	// Run the Go file
	cmd := exec.Command("go", "run", goFile)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	start := time.Now()
	err := cmd.Run()
	elapsed := time.Since(start)

	if err != nil {
		fmt.Printf("\nError running Go file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nTotal execution time: %v\n", elapsed)
}