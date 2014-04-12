package main

import (
	"bufio"
	"errors"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
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

func initCases() (*bufio.Scanner, int) {
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
	return input, cases
}

var columnError = errors.New("Invalid column count!")
