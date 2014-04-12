package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func readCases(input *bufio.Scanner) (int, error) {
	if !input.Scan() {
		return 0, io.EOF
	}
	line := input.Text()
	cases, err := strconv.ParseInt(line, 10, 8)
	if err != nil {
		return 0, err
	}
	return int(cases), nil
}

type row struct {
	cols [4]int
}

type board struct {
	rows [4]row
}

var columnError = errors.New("Invalid column count!")

func readBoard(input *bufio.Scanner) (int, board, error) {
	if !input.Scan() {
		return 0, board{}, io.EOF
	}
	line := input.Text()
	answer, err := strconv.ParseInt(line, 10, 8)
	if err != nil {
		return 0, board{}, err
	}
	board := board{}
	for r := 0; r < 4; r++ {
		if !input.Scan() {
			return int(answer), board, io.EOF
		}
		line = input.Text()
		cols := strings.Split(line, " ")
		if len(cols) != 4 {
			return int(answer), board, columnError
		}
		for c := 0; c < 4; c++ {
			val, err := strconv.ParseInt(cols[c], 10, 8)
			if err != nil {
				return int(answer), board, err
			}
			board.rows[r].cols[c] = int(val)
		}
	}
	return int(answer), board, nil
}

var magicianError = errors.New("Bad magician!")
var volunteerError = errors.New("Volunteer cheated!")

func findSolution(answer1 int, board1 board, answer2 int, board2 board) (int, error) {
	board1Row := board1.rows[answer1-1]
	log.Printf("Board1 row: %v", board1Row)
	board2Row := board2.rows[answer2-1]
	log.Printf("Board2 row: %v", board2Row)
	solutions := make([]int, 0, 4)
	for i := 0; i < 4; i++ {
		v := board1Row.cols[i]
		for j := 0; j < 4; j++ {
			if board2Row.cols[j] == v {
				solutions = append(solutions, v)
			}
		}
	}
	log.Printf("solutions: %v", solutions)
	solutionCount := len(solutions)
	var solution int
	if solutionCount == 0 {
		return 0, volunteerError
	} else if solutionCount > 1 {
		return 0, magicianError
	} else {
		solution = solutions[0]
	}
	return solution, nil
}

func main() {
	flag.Parse()
	inputName := flag.Arg(0)
	log.Printf("Input file: %s", inputName)
	inputRaw, err := os.Open(inputName)
	if err != nil {
		log.Fatal(err)
	}

	input := bufio.NewScanner(inputRaw)
	cases, err := readCases(input)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Test cases: %d", cases)

	for i := 0; i < cases; i++ {
		fmt.Printf("Case #%d: ", i+1)
		answer1, board1, err := readBoard(input)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Answer 1: %d Board 1: %v", answer1, board1)
		answer2, board2, err := readBoard(input)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Answer 2: %d Board 2: %v", answer1, board1)
		solution, err := findSolution(answer1, board1, answer2, board2)
		log.Printf("Solution: %d Error: %v", solution, err)
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Printf("%d", solution)
		}
		fmt.Println("")
	}
}
