package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	envKey                   = "AOC_TOKEN"
	envFile                  = "." + envKey
	aocURL                   = "https://adventofcode.com/%d%s"
	puzzleInputFile          = "puzzleinput.txt"
	nicenessSleep            = time.Second
	firstAOCYear             = 2015
	christmasDay             = 25
	yearFolderName           = "%d"
	dayFolderName            = "day%02d"
	progressBarChar          = "■"
	progressInProgChar       = "▪"
	progressEmptyChar        = "□"
	scriptFileInitialContent = "//Day%d\npackage main\n\nimport (\n\t\"advent-of-code/common\"\n\t\"bufio\"\n\t\"os\"\n)\n\nvar (\n\tpuzzleInput *os.File\n)\n\nfunc getInputLineScanner() *bufio.Scanner {\n\tfileScanner := bufio.NewScanner(puzzleInput)\n\tfileScanner.Split(bufio.ScanLines)\n\n\treturn fileScanner\n}\n\nfunc main() {\n\tvar err error\n\tpuzzleInput, err = common.OpenPuzzleInput(nil)\n\n\tif err != nil {\n\t\tpanic(err)\n\t}\n\n\tprintln(puzzleInput)\n\n\t//Content here\n}"
)

var (
	today     = time.Now()
	christmas = 25
)

// func ordinal(n int) string {
// 	suffix := "th"
// 	if n%10 == 1 && n%100 != 11 {
// 		suffix = "st"
// 	} else if n%10 == 2 && n%100 != 12 {
// 		suffix = "nd"
// 	} else if n%10 == 3 && n%100 != 13 {
// 		suffix = "rd"
// 	}
// 	return fmt.Sprintf("%d%s", n, suffix)
// }

func getToken(argumentInput string) (string, error) {
	if argumentInput != "" {
		fmt.Println("Using token passed via --token")
		return argumentInput, nil
	}

	if envVar, ok := os.LookupEnv(envKey); ok {
		fmt.Printf("Using token found in environment variable %s\n", envKey)
		return envVar, nil
	}

	fmt.Printf("Using token found in file %s\n", envFile)
	file, err := os.Open(envFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text(), nil
	}

	return "", errors.New("token file is empty")
}

func createFolder(day, year int) (string, error) {
	path := filepath.Join(".", fmt.Sprintf(yearFolderName, year), fmt.Sprintf(dayFolderName, day))

	if exists, _ := dirExists(path); exists {
		return path, nil
	}

	err := os.MkdirAll(path, os.ModePerm)
	return path, err
}

func dirExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
func createInputFile(path, data, filename string, overwrite bool) error {
	mode := os.O_WRONLY | os.O_CREATE
	if !overwrite {
		mode |= os.O_EXCL
	}

	file, err := os.OpenFile(filepath.Join(path, filename), mode, 0644)
	if err != nil {
		if os.IsExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	return err
}

func createScript(path, data, filename string, overwrite bool) error {
	mode := os.O_WRONLY | os.O_CREATE
	if !overwrite {
		mode |= os.O_EXCL
	}

	file, err := os.OpenFile(filepath.Join(path, filename), mode, 0644)
	if err != nil {
		if os.IsExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	return err
}

func aocInputRequest(token string, day, year int) (string, error) {
	url := fmt.Sprintf(aocURL, year, fmt.Sprintf("/day/%d/input", day))
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: token})
	req.Header.Set("User-Agent", "https://github.com/brunMartins/advent-of-code")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return string(body), err
}

func aocPuzzleNameRequest(token string, day, year int) (string, error) {
	url := fmt.Sprintf(aocURL, year, fmt.Sprintf("/day/%d", day))
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: token})
	req.Header.Set("User-Agent", "https://github.com/brunMartins/advent-of-code")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	bodyText := string(body)

	bodyRegex := regexp.MustCompile(`<h2>-{3}(.+)-{3}<\/h2>`)

	matches := bodyRegex.FindAllStringSubmatch(bodyText, 1)

	if len(matches) == 0 {
		panic("Could not find title!")
	}

	title := strings.TrimSpace(matches[0][1])
	return title, err
}

func aocPuzzleStarsRequest(token string, day, year int) (string, error) {
	url := fmt.Sprintf(aocURL, year, "")
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "0", err
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: token})
	req.Header.Set("User-Agent", "https://github.com/brunMartins/advent-of-code")
	resp, err := client.Do(req)
	if err != nil {
		return "0", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	bodyText := string(body)

	bodyRegex := regexp.MustCompile(fmt.Sprintf(`Day %d,\s(\w+\s\w+)`, day))

	matches := bodyRegex.FindAllStringSubmatch(bodyText, 1)
	if len(matches) > 0 {
		if strings.Contains(matches[0][1], "two") {
			return "2", err
		} else {
			return "1", err
		}
	}
	fmt.Println(matches)
	return "0", err
}

func handleDay(token string, day, year int, placeholder bool) error {
	path, err := createFolder(day, year)
	if err != nil {
		return err
	}

	var data = ""
	if !placeholder {
		data, err = aocInputRequest(token, day, year)
	}

	if err != nil {
		return err
	}

	puzzleName := ""
	puzzleName, err = aocPuzzleNameRequest(token, day, year)

	if err != nil {
		return err
	}

	puzzleStars := "0"
	puzzleStars, err = aocPuzzleStarsRequest(token, day, year)

	if err != nil {
		return err
	}

	if err = createScript(path, fmt.Sprintf(scriptFileInitialContent, day), fmt.Sprintf("day%d.go", day), false); err != nil {
		return err
	}

	if err = createInputFile(path, puzzleName, "puzzleName", true); err != nil {
		return err
	}

	if err = createInputFile(path, puzzleStars, "completed", true); err != nil {
		return err
	}

	if err = createInputFile(path, data, puzzleInputFile, true); err != nil {
		return err
	}

	return nil
}

func handlePopulate(token string, year int, placeholder bool) {
	lastDay := christmas
	if year == today.Year() && today.Day() < christmas {
		lastDay = today.Day()
	}

	for i := 1; i <= lastDay; i++ {
		fmt.Printf("%02d/%02d : %s : Working\n", i, lastDay, progressVisualAsString(lastDay, i-1))
		err := handleDay(token, i, year, placeholder)
		if err != nil {
			fmt.Printf("Error handling day %d: %v\n", i, err)
		}

		// if !placeholder && i != lastDay {
		// 	fmt.Printf("%02d/%02d : %s : Sleeping\n", i, lastDay, progressVisualAsString(lastDay, i))
		// 	time.Sleep(nicenessSleep)
		// }
	}
	fmt.Printf("%02d/%02d : %s : Done\n", lastDay, lastDay, progressVisualAsString(lastDay, lastDay))
}

func progressVisualAsString(total, done int) string {
	if done >= total {
		return strings.Repeat(progressBarChar, total)
	}
	return strings.Repeat(progressBarChar, done) +
		progressInProgChar +
		strings.Repeat(progressEmptyChar, total-done-1)
}

func main() {
	day := flag.Int("day", 0, "Puzzle day")
	year := flag.Int("year", 0, "Puzzle year")
	populate := flag.Bool("populate", false, "Use to populate a whole year")
	placeholder := flag.Bool("placeholder", false, "Create the file structure without requests to AOC")
	tokenArg := flag.String("token", "", "AOC session token (priority over environment variable and local dotfile)")
	flag.Parse()

	token, err := getToken(*tokenArg)
	if err != nil {
		fmt.Printf("Error retrieving token: %v\n", err)
		os.Exit(1)
	}

	currentYear := today.Year()
	if *year != 0 && *year >= firstAOCYear && *year <= currentYear {
		currentYear = *year
	}

	if *populate {
		fmt.Printf("Populating %d with %s\n", currentYear, map[bool]string{true: "placeholders", false: "puzzle inputs"}[*placeholder])
		handlePopulate(token, currentYear, *placeholder)
	} else {
		currentDay := today.Day()
		if *day != 0 && *day <= 25 {
			currentDay = *day
		}
		fmt.Printf("Fetching puzzle input for %02d/%d\n", currentDay, currentYear)
		err := handleDay(token, currentDay, currentYear, *placeholder)
		if err != nil {
			fmt.Printf("Error handling day: %v\n", err)
		}
	}
}
