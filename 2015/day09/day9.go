// Day9
package main

import (
	"advent-of-code/common"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

var (
	puzzleInput      *os.File
	shortestDistance int
	trips            = make(map[string]map[string]int, 1)
	locationList     []string
	longestDistance  int
)

func processTrip(trip string) {
	tripParts := common.SplitString(trip)
	loc1, loc2 := tripParts[0], tripParts[2]
	distance, err := strconv.Atoi(tripParts[4])

	if err != nil {
		panic(err)
	}

	// Create bi-directional entries in the map
	if _, exists := trips[loc1]; !exists {
		trips[loc1] = make(map[string]int)
	}
	if _, exists := trips[loc2]; !exists {
		trips[loc2] = make(map[string]int)
	}
	trips[loc1][loc2] = distance
	trips[loc2][loc1] = distance
}

func Permutations(arr []string) [][]string {
	if len(arr) == 0 {
		return [][]string{}
	}
	if len(arr) == 1 {
		return [][]string{arr}
	}

	var result [][]string
	for i, v := range arr {
		remaining := append([]string{}, arr[:i]...)
		remaining = append(remaining, arr[i+1:]...)
		for _, perm := range Permutations(remaining) {
			result = append(result, append([]string{v}, perm...))
		}
	}
	return result
}

func CalculateRouteDistance(route []string) int {
	totalDistance := 0
	for i := 0; i < len(route)-1; i++ {
		totalDistance += trips[route[i]][route[i+1]]
	}
	return totalDistance
}

func calculateShortestRoute() {
	scanner := common.GetInputLineScanner(puzzleInput)

	for scanner.Scan() {
		processTrip(scanner.Text())
	}
	for location := range trips {
		locationList = append(locationList, location)
	}

	allRoutes := Permutations(locationList)

	shortestDistance = math.MaxInt
	longestDistance = math.MinInt
	for _, route := range allRoutes {
		distance := CalculateRouteDistance(route)
		if distance < shortestDistance {
			shortestDistance = distance
		}

		if distance > longestDistance {
			longestDistance = distance
		}
	}
}

func main() {
	start := time.Now()

	var err error
	puzzleInput, err = common.OpenPuzzleInput(nil)

	if err != nil {
		panic(err)
	}

	calculateShortestRoute()

	println("The shortest distance possible for a trip is:", shortestDistance)
	println("The longest distance possible for a trip is:", longestDistance)

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
