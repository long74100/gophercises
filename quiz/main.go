package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problem.csv", "a csv file in the format of 'questions,answers")
	flag.Parse()

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open csv: %s", *csvFilename), 1)
	}

	csvReader := csv.NewReader(file)

	lines, err := csvReader.ReadAll()

	if err != nil {
		exit(fmt.Sprintf("Failed to parse csv: %s", *csvFilename), 1)
	}

	problems := linesToProblems(lines)

	numCorrect := 0

	for i, problem := range problems {
		fmt.Printf("Problem %d: %s = \n", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == problem.answer {
			numCorrect++
		}
	}

	fmt.Printf("You got %d out of %d problems correct!", numCorrect, len(problems))
}

func linesToProblems(lines [][]string) []problem {
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}

	return problems
}

type problem struct {
	question string
	answer   string
}

func exit(msg string, code int) {
	fmt.Println(msg)
	os.Exit(code)
}
