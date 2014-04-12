package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type testCase struct {
	farmPrice      float64
	farmProduction float64
	targetScore    float64
}

func readTestCase(input *bufio.Scanner) (testCase, error) {
	var testCase testCase
	if !input.Scan() {
		return testCase, io.EOF
	}
	line := input.Text()
	tokens := strings.Split(line, " ")
	if len(tokens) != 3 {
		return testCase, columnError
	}
	var err error
	testCase.farmPrice, err = strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return testCase, err
	}
	testCase.farmProduction, err = strconv.ParseFloat(tokens[1], 64)
	if err != nil {
		return testCase, err
	}
	testCase.targetScore, err = strconv.ParseFloat(tokens[2], 64)
	if err != nil {
		return testCase, err
	}
	return testCase, nil
}

func calculateTime(testCase testCase, production float64) float64 {
	log.Printf("New step: production: %f", production)
	timeToFarm := testCase.farmPrice / production
	timeToTarget := testCase.targetScore / production
	log.Printf("TTF: %f TTT: %f", timeToFarm, timeToTarget)
	newProduction := production + testCase.farmProduction
	timeWithFarm := timeToFarm + (testCase.targetScore / newProduction)
	if timeToTarget <= timeWithFarm {
		return timeToTarget
	} else {
		return timeToFarm + calculateTime(testCase, newProduction)
	}
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
		testCase, err := readTestCase(input)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Case data: %+v", testCase)
		time := calculateTime(testCase, 2)
		fmt.Printf("%0.7f\n", time)
	}
}
