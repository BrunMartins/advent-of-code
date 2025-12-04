#!/bin/bash

# Advent of Code Runner Script
# Supports both Go and JavaScript solutions

show_usage() {
    echo "Usage: $0 [OPTIONS]"
    echo ""
    echo "Options:"
    echo "  -d, --day <day>       Puzzle day (1-25)"
    echo "  -y, --year <year>     Puzzle year"
    echo "  -l, --lang <lang>     Language (go|js|javascript). If not specified, auto-detect"
    echo "  -h, --help           Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 --day 1 --year 2024 --lang go"
    echo "  $0 -d 5 -y 2025 -l js"
    echo "  $0 --day 3                    # Auto-detect language for current year"
    echo "  $0                            # Run current day with auto-detect"
    echo ""
    echo "Auto-detection priority: Go > JavaScript"
}

# Default values
DAY=0
YEAR=0
LANG=""

# Parse command line arguments
while [[ $# -gt 0 ]]; do
    case $1 in
        -d|--day)
            DAY="$2"
            shift 2
            ;;
        -y|--year)
            YEAR="$2"
            shift 2
            ;;
        -l|--lang)
            LANG="$2"
            shift 2
            ;;
        -h|--help)
            show_usage
            exit 0
            ;;
        *)
            echo "Unknown option: $1"
            show_usage
            exit 1
            ;;
    esac
done

# Set defaults if not provided
if [ $YEAR -eq 0 ]; then
    YEAR=$(date +%Y)
fi

if [ $DAY -eq 0 ]; then
    DAY=$(date +%d | sed 's/^0*//')  # Remove leading zeros
fi

# Validate inputs
if [ $DAY -lt 1 ] || [ $DAY -gt 25 ]; then
    echo "Error: Day must be between 1 and 25, got $DAY"
    exit 1
fi

if [ $YEAR -lt 2015 ]; then
    echo "Error: Year must be 2015 or later, got $YEAR"
    exit 1
fi

# Build paths
DAY_FOLDER=$(printf "day%02d" $DAY)
YEAR_FOLDER="$YEAR"
GO_FILE="day${DAY}.go"
JS_FILE="day${DAY}.js"
DIR_PATH="${YEAR_FOLDER}/${DAY_FOLDER}"
GO_PATH="${DIR_PATH}/${GO_FILE}"
JS_PATH="${DIR_PATH}/${JS_FILE}"

# Function to run Go solution
run_go() {
    if [ ! -f "$GO_PATH" ]; then
        echo "Error: Go file does not exist at $GO_PATH"
        exit 1
    fi
    
    echo "Running Day $DAY, Year $YEAR (Go)"
    echo "File: $GO_PATH"
    echo "Directory: $DIR_PATH"
    echo "=================================================="
    
    cd "$DIR_PATH" || exit 1
    start_time=$(date +%s%3N)
    
    if go run "$GO_FILE"; then
        end_time=$(date +%s%3N)
        elapsed=$((end_time - start_time))
        echo ""
        echo "Total execution time: ${elapsed}ms"
    else
        echo ""
        echo "Error running Go file"
        exit 1
    fi
}

# Function to run JavaScript solution
run_js() {
    if [ ! -f "$JS_PATH" ]; then
        echo "Error: JavaScript file does not exist at $JS_PATH"
        exit 1
    fi
    
    echo "Running Day $DAY, Year $YEAR (JavaScript)"
    echo "File: $JS_PATH"
    echo "Directory: $DIR_PATH"
    echo "=================================================="
    
    cd "$DIR_PATH" || exit 1
    start_time=$(date +%s%3N)
    
    if node "$JS_FILE"; then
        end_time=$(date +%s%3N)
        elapsed=$((end_time - start_time))
        echo ""
        echo "Total execution time: ${elapsed}ms"
    else
        echo ""
        echo "Error running JavaScript file"
        exit 1
    fi
}

# Auto-detect language if not specified
if [ -z "$LANG" ]; then
    if [ -f "$GO_PATH" ]; then
        LANG="go"
    elif [ -f "$JS_PATH" ]; then
        LANG="js"
    else
        echo "Error: No solution files found at $DIR_PATH"
        echo "Looking for: $GO_FILE or $JS_FILE"
        exit 1
    fi
fi

# Normalize language parameter
case "${LANG,,}" in  # Convert to lowercase
    go|golang)
        run_go
        ;;
    js|javascript|node)
        run_js
        ;;
    *)
        echo "Error: Unsupported language '$LANG'. Use 'go' or 'js'"
        exit 1
        ;;
esac