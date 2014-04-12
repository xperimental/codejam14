package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type mineConfig struct {
	rows  int
	cols  int
	mines int
}

func convInt(input string) (int, error) {
	val, err := strconv.ParseInt(input, 10, 8)
	if err != nil {
		return 0, err
	} else {
		return int(val), nil
	}
}

func readConfig(input *bufio.Scanner) (mineConfig, error) {
	var config mineConfig
	if !input.Scan() {
		return config, io.EOF
	}
	line := input.Text()
	tokens := strings.Split(line, " ")
	if len(tokens) != 3 {
		return config, columnError
	}
	var err error
	config.rows, err = convInt(tokens[0])
	if err != nil {
		return config, err
	}
	config.cols, err = convInt(tokens[1])
	if err != nil {
		return config, err
	}
	config.mines, err = convInt(tokens[2])
	if err != nil {
		return config, err
	}
	return config, nil
}

var impossibleError = errors.New("Impossible!")

func solve(config mineConfig) (string, error) {
	return "", impossibleError
}

func main() {
	input, cases := initCases()

	for i := 0; i < cases; i++ {
		fmt.Printf("Case #%d:\n", i+1)
		config, err := readConfig(input)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("config: %+v", config)
		solution, err := solve(config)
		switch {
		case err == impossibleError:
			fmt.Println(err)
		case err != nil:
			log.Fatal(err)
		default:
			fmt.Println(solution)
		}
	}
}
