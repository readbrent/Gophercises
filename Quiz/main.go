package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFile := flag.String("csv", "problems.csv", "CSV separating question,answer")
	flag.Parse()

	file, err := os.Open(*csvFile)
	if err != nil {
		exit(fmt.Sprintf("Could not open CSV file: %s\n", *csvFile))
	}

	r := csv.NewReader(file)

	// Parse the file into a slice.
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to read the CSV")
	}

	problems := parseLines(lines)
	correct := 0

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, problem.question)

		var answer string

		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			fmt.Println("Correct")
			correct++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {

	// Create return slice. Same length input.
	ret := make([]problem, len(lines))

	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
