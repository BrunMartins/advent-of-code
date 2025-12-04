// Day 4
const common = require('../../common');

function getInputLines(puzzleInput) {
	return puzzleInput.split('\n');
}

function solvePuzzlePart1(puzzleInput) {
	let result = 0;
	// TODO: Solve part 1
	
	return result;
}

function solvePuzzlePart2(puzzleInput) {
	let result = 0;
	// TODO: Solve part 2
	
	return result;
}

function processPuzzle(puzzleInput) {
	const part1 = solvePuzzlePart1(puzzleInput);
	const part2 = solvePuzzlePart2(puzzleInput);
	
	return { part1, part2 };
}

function main() {
	const start = Date.now();
	
	const puzzleInput = common.readPuzzleInput(false, __dirname);
	console.log('Puzzle input loaded');
	
	const { part1, part2 } = processPuzzle(puzzleInput);
	
	const elapsed = Date.now() - start;
	console.log(`Execution took ${elapsed}ms`);
	
	console.info('The solution according to part 1 rules is: ', part1);
	console.info('The solution according to part 2 rules is: ', part2);
}

if (require.main === module) {
	main();
}