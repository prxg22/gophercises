package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

const default_path = "problems.csv"
const default_time = 5

func readCSV(file *os.File) ([][]string, error) {
	reader := csv.NewReader(file)

	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, nil
}

func run(file *os.File, t *int) (int, int, error) {
	lines, err := readCSV(file)
	if err != nil {
		return 0, 0, err
	}

	ans := make(chan string)
	len_lines := len(lines)
	for i, line := range lines {
		question, answer, timer := line[0], line[1], time.NewTimer(time.Duration(*t)*time.Second)

		fmt.Printf("Question: %v ", question)
		go getAnswer(&ans)

		select {
		case a := <-ans:
			if a == answer {
				fmt.Println("Correct answer!")
			} else {
				return i, len_lines, fmt.Errorf("wrong answer")
			}
		case <-timer.C:
			return i, len_lines, fmt.Errorf("time's up")
		}
	}

	return len_lines, len_lines, nil
}

func getAnswer(ans *chan string) {
	var input string
	fmt.Scanln(&input)
	*ans <- input
}

func main() {
	path := flag.String("path", default_path, "a csv file in the format of 'question,answer'")
	t := flag.Int("time", default_time, "time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*path)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	correct, total, err := run(file, t)
	if err != nil {
		fmt.Printf("\n%v! You got %v correct answers out of %d\n", err, correct, total)
	} else {
		fmt.Printf("\nYou got %v correct answers out of %d\n", correct, total)
	}
}
