// Day4
package main

import (
	"advent-of-code/common"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

var (
	puzzleInput *os.File
)

func generateHash() {
	scanner := common.GetInputLineScanner(puzzleInput)
	scanner.Scan()
	input := scanner.Text()
	puzzleInput.Close()
	counter := 0

	for {
		hash := md5.New()
		str := fmt.Sprintf("%s%d", input, counter)

		hash.Write([]byte(str))

		hashSum := hash.Sum(nil)
		hashString := hex.EncodeToString(hashSum)

		if hashString[0:5] == "00000" {
			println("The lowest number to generate a 5 leading 0 result was: ", counter)
			break
		}
		counter++
	}

	counter = 0

	for {
		hash := md5.New()
		str := fmt.Sprintf("%s%d", input, counter)

		hash.Write([]byte(str))

		hashSum := hash.Sum(nil)
		hashString := hex.EncodeToString(hashSum)

		if hashString[0:6] == "000000" {
			println("The lowest number to generate a 6 leading 0 result was: ", counter)
			break
		}
		counter++
	}
}

func main() {
	var err error
	puzzleInput, err = common.OpenPuzzleInput(nil)

	if err != nil {
		panic(err)
	}

	generateHash()

	// Content here
}
