// Day 3
const common = require("../../common");

function solvePuzzlePart1(puzzleInput) {
  let result = [];

  const batteryBanks = common.getInputLines(puzzleInput);

  for (const batteryBank of batteryBanks) {
    let heighestRating = 0;
    const batteryJoltages = String(batteryBank)
      .split("")
      .map((b) => Number(b));
    for (let i = 0; i < batteryJoltages.length; i++) {
      const firstDigit = batteryJoltages[i];
      for (let j = i + 1; j < batteryJoltages.length; j++) {
        const combo = Number(`${firstDigit}${batteryJoltages[j]}`);
        if (combo > heighestRating) heighestRating = combo;
      }
    }
    result.push(heighestRating);
  }

  return result.reduce((sum, item) => sum + item, 0);
}

function solvePuzzlePart2(puzzleInput) {
  let result = [];
  const batteryBanks = common.getInputLines(puzzleInput);
  const targetLength = 12;

  for (const batteryBank of batteryBanks) {
    const batteryJoltages = String(batteryBank).split("");
    const stack = [];
    const canRemove = batteryJoltages.length - targetLength;
    let removed = 0;

    for (let i = 0; i < batteryJoltages.length; i++) {
      while (
        stack.length > 0 &&
        batteryJoltages[i] > stack[stack.length - 1] &&
        removed < canRemove &&
        stack.length + (batteryJoltages.length - i - 1) >= targetLength
      ) {
        stack.pop();
        removed++;
      }

      if (stack.length < targetLength) {
        stack.push(batteryJoltages[i]);
      }
    }

    while (stack.length > targetLength) {
      stack.pop();
    }

    const resultNumber = Number(stack.join(""));
    result.push(resultNumber);
  }

  return result.reduce((sum, item) => sum + item, 0);
}

function processPuzzle(puzzleInput) {
  const part1 = solvePuzzlePart1(puzzleInput);
  const part2 = solvePuzzlePart2(puzzleInput);

  return { part1, part2 };
}

function main() {
  const start = Date.now();

  const puzzleInput = common.readPuzzleInput(false, __dirname);
  console.log("Puzzle input loaded");

  const { part1, part2 } = processPuzzle(puzzleInput);

  const elapsed = Date.now() - start;
  console.log(`Execution took ${elapsed}ms`);

  console.info("The solution according to part 1 rules is: ", part1);
  console.info("The solution according to part 2 rules is: ", part2);
}

if (require.main === module) {
  main();
}
