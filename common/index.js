const fs = require("fs");
const path = require("path");
const readline = require("readline");

/**
 * JavaScript Common Module for Advent of Code
 * Provides utility functions similar to the Go common module
 */

// Constants
const PUZZLE_INPUT = "puzzleinput.txt";
const PUZZLE_TEST_INPUT = "puzzle-test.txt";
const DEFAULT_TEST_MODE = false;

/**
 * Read puzzle input file and return as string
 * @param {boolean} testMode - Whether to read test input or regular input
 * @param {string} callerDir - Directory of the calling script (use __dirname from caller)
 * @returns {string} - The puzzle input content
 */
function readPuzzleInput(testMode = DEFAULT_TEST_MODE, callerDir = null) {
  try {
    const filename = testMode ? PUZZLE_TEST_INPUT : PUZZLE_INPUT;
    // Use caller's __dirname if provided, otherwise try current working directory
    const baseDir = callerDir || process.cwd();
    const inputPath = path.join(baseDir, filename);
    console.log(`Reading from: ${inputPath}`);
    return fs.readFileSync(inputPath, "utf8").trim();
  } catch (error) {
    console.error(`Error reading puzzle input: ${error.message}`);
    console.error(
      `Attempted to read from: ${path.join(
        callerDir || process.cwd(),
        filename
      )}`
    );
    process.exit(1);
  }
}

/**
 * Read puzzle input and return as array of lines
 * @param {string|null} puzzleInput - The puzzle input string, or null to read from file
 * @param {boolean} testMode - Whether to read test input or regular input
 * @returns {string[]} - Array of input lines
 */
function getInputLines(puzzleInput = null, testMode = DEFAULT_TEST_MODE) {
  const input = puzzleInput || readPuzzleInput(testMode);
  return input.split("\n").filter((line) => line.length > 0);
}

/**
 * Read puzzle input and return as single string (no line breaks)
 * @param {string|null} puzzleInput - The puzzle input string, or null to read from file
 * @param {boolean} testMode - Whether to read test input or regular input
 * @returns {string} - Whole input content without line breaks
 */
function getWholeInputContent(
  puzzleInput = null,
  testMode = DEFAULT_TEST_MODE
) {
  const input = puzzleInput || readPuzzleInput(testMode);
  return input.replace(/\n/g, "");
}

/**
 * Split string into array of words (equivalent to Go's strings.Fields)
 * @param {string} data - String to split
 * @returns {string[]} - Array of words
 */
function splitString(data) {
  return data
    .trim()
    .split(/\s+/)
    .filter((word) => word.length > 0);
}

/**
 * Convert array of strings to array of integers
 * @param {string[]} strArray - Array of string numbers
 * @returns {number[]} - Array of integers
 */
function arrayAtoI(strArray) {
  return strArray.map((str) => {
    const num = parseInt(str, 10);
    if (isNaN(num)) {
      throw new Error(`Cannot convert '${str}' to integer`);
    }
    return num;
  });
}

/**
 * Find minimum of two numbers
 * @param {number} a - First number
 * @param {number} b - Second number
 * @returns {number} - Minimum value
 */
function min(a, b) {
  return Math.min(a, b);
}

/**
 * Find maximum of two numbers
 * @param {number} a - First number
 * @param {number} b - Second number
 * @returns {number} - Maximum value
 */
function max(a, b) {
  return Math.max(a, b);
}

/**
 * Reverse a string
 * @param {string} str - String to reverse
 * @returns {string} - Reversed string
 */
function reverseString(str) {
  return str.split("").reverse().join("");
}

/**
 * Convert string to array of individual characters
 * @param {string} str - String to convert
 * @returns {string[]} - Array of character strings
 */
function stringToStringArray(str) {
  return str.split("");
}

/**
 * Check if array contains a specific string
 * @param {string[]} array - Array to search in
 * @param {string} target - String to search for
 * @returns {boolean} - True if found, false otherwise
 */
function stringArrayContains(array, target) {
  return array.includes(target);
}

/**
 * Create a readline interface for user input
 * @returns {Object} - Readline interface and helper function
 */
function createReadlineInterface() {
  const readline = require("readline");
  const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
  });

  const askQuestion = (question) => {
    return new Promise((resolve) => {
      rl.question(question, (answer) => {
        resolve(answer);
      });
    });
  };

  return { rl, askQuestion };
}

/**
 * Clear the console
 */
function clearConsole() {
  console.clear();
}

/**
 * Parse numbers from a string using regex
 * @param {string} str - String containing numbers
 * @param {boolean} includeNegative - Whether to include negative numbers
 * @returns {number[]} - Array of found numbers
 */
function parseNumbers(str, includeNegative = false) {
  const pattern = includeNegative ? /-?\d+/g : /\d+/g;
  const matches = str.match(pattern);
  return matches ? matches.map((num) => parseInt(num, 10)) : [];
}

/**
 * Sum an array of numbers
 * @param {number[]} numbers - Array of numbers to sum
 * @returns {number} - Sum of all numbers
 */
function sum(numbers) {
  return numbers.reduce((acc, num) => acc + num, 0);
}

/**
 * Create a 2D grid from input lines
 * @param {string[]} lines - Array of input lines
 * @returns {string[][]} - 2D array representing the grid
 */
function createGrid(lines) {
  return lines.map((line) => line.split(""));
}

/**
 * Measure execution time of a function
 * @param {Function} fn - Function to measure
 * @param {...any} args - Arguments to pass to the function
 * @returns {Object} - Object containing result and execution time
 */
function measureTime(fn, ...args) {
  const start = Date.now();
  const result = fn(...args);
  const elapsed = Date.now() - start;
  return { result, elapsed };
}

function askQuestion(question) {
  return new Promise((resolve) => {
    rl.question(question, (answer) => {
      resolve(answer);
    });
  });
}

module.exports = {
  // File I/O
  readPuzzleInput,
  getInputLines,
  getWholeInputContent,

  // String utilities
  splitString,
  reverseString,
  stringToStringArray,
  stringArrayContains,
  parseNumbers,

  // Array utilities
  arrayAtoI,
  sum,
  createGrid,

  // Math utilities
  min,
  max,

  // Interactive utilities
  createReadlineInterface,
  clearConsole,
  askQuestion,

  // Performance utilities
  measureTime,

  // Constants
  PUZZLE_INPUT,
  PUZZLE_TEST_INPUT,
  DEFAULT_TEST_MODE,
};
