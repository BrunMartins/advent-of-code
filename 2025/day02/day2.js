// Day 2
const common = require("../../common");

function getInputLines(puzzleInput) {
  return puzzleInput.split("\n");
}

function solvePuzzlePart1(puzzleInput) {
  let result = [];
  const puzzleIds = puzzleInput.trim().split(",");
  for (const ids of puzzleIds) {
    const [firstId, lastId] = ids.split("-");
    for (let pointer = parseInt(firstId); pointer <= lastId; pointer++) {
      const idString = String(pointer);
      if (idString.length % 2 !== 0) continue;
      const midPoint = idString.length / 2;
      const [idLeftPart, idRightPart] = [
        idString.substring(0, midPoint),
        idString.substring(midPoint),
      ];
      if (idLeftPart == idRightPart) result.push(pointer);
    }
  }

  return result.reduce((sum, item) => sum + item, 0);
}

function solvePuzzlePart2(puzzleInput) {
  let result = [];
  const puzzleIds = puzzleInput.split(",");
  for (const ids of puzzleIds) {
    const [firstId, lastId] = ids.split("-");
    for (let pointer = parseInt(firstId); pointer <= lastId; pointer++) {
      const idString = String(pointer);
      for (let i = 1; i <= idString.length; i++) {
        const targetCombo = idString.substring(0, i);
        if (idString.length % i !== 0) continue;
        const timesRepeated = idString.length / i;

        if (
          timesRepeated >= 2 &&
          targetCombo.repeat(timesRepeated) === idString
        ) {
          result.push(pointer);
          break;
        }
      }
    }
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
