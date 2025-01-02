// Day7
package main

import (
	"advent-of-code/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	op     string   // Operation (e.g., "AND", "OR", "NOT", "LSHIFT", "RSHIFT", "ASSIGN")
	inputs []string // Input wires or values
	output string   // Output wire
}

var (
	puzzleInput          *os.File
	wires                = make(map[string]uint16)
	instructions         []Instruction
	originalInstructions []Instruction
)

func getInputLineScanner() *bufio.Scanner {
	fileScanner := bufio.NewScanner(puzzleInput)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}

// ParseInstruction parses a raw instruction string into an Instruction struct
func ParseInstruction(raw string) Instruction {
	parts := strings.Split(raw, " -> ")
	output := parts[1]
	expr := strings.Fields(parts[0])

	var op string
	var inputs []string

	switch len(expr) {
	case 1:
		op = "ASSIGN"
		inputs = []string{expr[0]}
	case 2:
		op = expr[0]
		inputs = []string{expr[1]}
	case 3:
		op = expr[1]
		inputs = []string{expr[0], expr[2]}
	}

	return Instruction{op, inputs, output}
}

// CanProcess checks if all inputs for an instruction are available
func CanProcess(instr Instruction) bool {
	for _, input := range instr.inputs {
		// Check if the input is a number (constants are always available)
		if _, err := strconv.Atoi(input); err != nil {
			// If it's not a number, check if the wire exists
			if _, exists := wires[input]; !exists {
				return false
			}
		}
	}

	return true
}

func RestartProcessing() {
	wires = make(map[string]uint16)
	puzzleInput.Seek(0, 0)
	instructions = []Instruction{}
}

func Process(instr Instruction) {
	var result uint16

	switch instr.op {
	case "ASSIGN":
		value, err := strconv.Atoi(instr.inputs[0])
		if err != nil {
			result = wires[instr.inputs[0]]
		} else {
			result = uint16(value)
		}
	case "NOT":
		value := wires[instr.inputs[0]]
		result = ^value
	case "AND":
		left := getValue(instr.inputs[0], wires)
		right := getValue(instr.inputs[1], wires)
		result = left & right
	case "OR":
		left := getValue(instr.inputs[0], wires)
		right := getValue(instr.inputs[1], wires)
		result = left | right
	case "LSHIFT":
		value := wires[instr.inputs[0]]
		shift, err := strconv.Atoi(instr.inputs[1])
		if err != nil {
			panic("Invalid shift value")
		}
		result = value << shift
	case "RSHIFT":
		value := wires[instr.inputs[0]]
		shift, err := strconv.Atoi(instr.inputs[1])
		if err != nil {
			panic("Invalid shift value")
		}
		result = value >> shift
	default:
		panic("Unknown operation: " + instr.op)
	}

	// Store the result in the output wire
	wires[instr.output] = result
}

// getValue retrieves the value of a wire or constant
func getValue(input string, wires map[string]uint16) uint16 {
	if value, err := strconv.Atoi(input); err == nil {
		// Input is a constant
		return uint16(value)
	}
	// Input is a wire
	return wires[input]
}

func fetchInstructions() {
	scanner := getInputLineScanner()
	for scanner.Scan() {
		originalInstructions = append(originalInstructions, ParseInstruction(scanner.Text()))
	}
}

func processInstructions() {

	for len(instructions) > 0 {
		progress := false
		var remainingInstructions []Instruction

		for _, instr := range instructions {
			if CanProcess(instr) {
				Process(instr)
				progress = true
			} else {
				remainingInstructions = append(remainingInstructions, instr)
			}
		}

		if !progress {
			panic("Unresolved dependencies")
		}

		instructions = remainingInstructions
	}
}

func main() {
	var err error
	puzzleInput, err = common.OpenPuzzleInput(nil)
	if err != nil {
		panic(err)
	}
	defer puzzleInput.Close()

	fetchInstructions()

	instructions = make([]Instruction, len(originalInstructions))
	copy(instructions, originalInstructions)

	processInstructions()
	valueA := wires["a"]
	println("Part One: The \"a\" wire has a value of", valueA)

	RestartProcessing()

	instructions = make([]Instruction, len(originalInstructions))
	copy(instructions, originalInstructions)

	filtered := []Instruction{}
	for _, instr := range instructions {
		if instr.output != "b" {
			filtered = append(filtered, instr)
		} else {
			fmt.Println(instr)
		}
	}

	instructions = append(filtered, Instruction{"ASSIGN", []string{strconv.FormatUint(uint64(valueA), 10)}, "b"})
	processInstructions()
	println("Part Two: The \"a\" wire has a value of", wires["a"])
}
