const fs = require('fs');
const path = require('path');
const { spawn } = require('child_process');

function parseArgs() {
    const args = process.argv.slice(2);
    let day = 0;
    let year = 0;

    for (let i = 0; i < args.length; i++) {
        if (args[i] === '--day' && i + 1 < args.length) {
            day = parseInt(args[i + 1]);
            i++;
        } else if (args[i] === '--year' && i + 1 < args.length) {
            year = parseInt(args[i + 1]);
            i++;
        } else if (args[i] === '--help' || args[i] === '-h') {
            console.log('Usage: node main.js [--day <day>] [--year <year>]');
            console.log('  --day   Puzzle day (1-25)');
            console.log('  --year  Puzzle year');
            process.exit(0);
        }
    }

    return { day, year };
}

function main() {
    const { day: inputDay, year: inputYear } = parseArgs();
    
    const today = new Date();
    let currentYear = inputYear || today.getFullYear();
    let currentDay = inputDay || today.getDate();

    // Validate inputs
    if (currentDay < 1 || currentDay > 25) {
        console.error(`Error: Day must be between 1 and 25, got ${currentDay}`);
        process.exit(1);
    }

    if (currentYear < 2015) {
        console.error(`Error: Year must be 2015 or later, got ${currentYear}`);
        process.exit(1);
    }

    // Build the path to the JavaScript file (relative to root directory)
    const dayFolder = `day${currentDay.toString().padStart(2, '0')}`;
    const yearFolder = currentYear.toString();
    const jsFile = `day${currentDay}.js`;
    const filePath = path.join('..', '..', yearFolder, dayFolder, jsFile);
    const dir = path.join('..', '..', yearFolder, dayFolder);

    // Check if the file exists
    if (!fs.existsSync(filePath)) {
        console.error(`Error: JavaScript file does not exist at ${filePath}`);
        process.exit(1);
    }

    console.log(`Running Day ${currentDay}, Year ${currentYear} (JavaScript)`);
    console.log(`File: ${filePath}`);
    console.log(`Directory: ${dir}`);
    console.log('=' + '='.repeat(50));

    // Run the JavaScript file
    const start = Date.now();
    const child = spawn('node', [jsFile], {
        cwd: dir,
        stdio: 'inherit'
    });

    child.on('close', (code) => {
        const elapsed = Date.now() - start;
        
        if (code !== 0) {
            console.error(`\nError running JavaScript file: Process exited with code ${code}`);
            process.exit(code);
        }

        console.log(`\nTotal execution time: ${elapsed}ms`);
    });

    child.on('error', (err) => {
        console.error(`\nError running JavaScript file: ${err.message}`);
        process.exit(1);
    });
}

if (require.main === module) {
    main();
}