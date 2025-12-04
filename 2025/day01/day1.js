// Day 1
const { readPuzzleInput } = require("../../common");

const directions = Object.freeze({
  L: -1,
  R: +1,
});

const splitterReg = new RegExp(/([R|L]{1})([\d]{1,})/);

function getInputLines(puzzleInput) {
  return puzzleInput.split("\n");
}

function solvePuzzlePart1(puzzleInput) {
  let result = 0;
  const actions = getInputLines(puzzleInput).map((puzzleItem) => {
    const itemValues = splitterReg.exec(puzzleItem);

    return { dir: itemValues[1], steps: parseInt(itemValues[2], 10) };
  });

  let step = 50;
  for (const action of actions) {
    const directionOp = directions[action.dir];

    for (let i = 1; i <= action.steps; i++) {
      step += directionOp;
      if (step == -1) {
        step = 99;
      } else if (step === 100) {
        step = 0;
      }

      if (i == action.steps && step === 0) result++;
    }
  }

  return result;
}

function solvePuzzlePart2(puzzleInput) {
  let result = 0;
  const actions = getInputLines(puzzleInput).map((puzzleItem) => {
    const itemValues = splitterReg.exec(puzzleItem);

    return { dir: itemValues[1], steps: parseInt(itemValues[2], 10) };
  });

  let step = 50;
  for (const action of actions) {
    const directionOp = directions[action.dir];

    for (let i = 1; i <= action.steps; i++) {
      step += directionOp;
      if (step === -1) {
        step = 99;
      } else if (step === 100) {
        step = 0;
      }

      if (step === 0) result++;
    }
  }

  return result;
}

function main() {
  const start = Date.now();

  const puzzleInput = readPuzzleInput();
  console.log("Puzzle input loaded");

  const puzzleSolution1 = solvePuzzlePart1(puzzleInput);
  const puzzleSolution2 = solvePuzzlePart2(puzzleInput);

  const elapsed = Date.now() - start;
  console.log(`Execution took ${elapsed}ms`);

  console.info(
    "The safe combination according to part 1 rules is: ",
    puzzleSolution1
  );
  console.info(
    "The safe combination according to part 2 rules is: ",
    puzzleSolution2
  );
}

if (require.main === module) {
  main();
}
