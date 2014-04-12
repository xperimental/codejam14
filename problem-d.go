package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type gameConfig struct {
	blocks int
	naomi  []float64
	ken    []float64
}

func readBlocks(line string, blockCount int) ([]float64, error) {
	blocks := make([]float64, blockCount)
	tokens := strings.Split(line, " ")
	if len(tokens) != blockCount {
		return blocks, columnError
	}
	for i := 0; i < blockCount; i++ {
		var err error
		blocks[i], err = strconv.ParseFloat(tokens[i], 64)
		if err != nil {
			return blocks, err
		}
	}
	return blocks, nil
}

func readGameConfig(input *bufio.Scanner) (gameConfig, error) {
	var config gameConfig
	if !input.Scan() {
		return config, io.EOF
	}
	line := input.Text()
	blockCount, err := convInt(line)
	if err != nil {
		return config, err
	}
	config.blocks = blockCount
	if !input.Scan() {
		return config, io.EOF
	}
	line = input.Text()
	config.naomi, err = readBlocks(line, blockCount)
	if err != nil {
		return config, err
	}
	if !input.Scan() {
		return config, io.EOF
	}
	line = input.Text()
	config.ken, err = readBlocks(line, blockCount)
	if err != nil {
		return config, err
	}
	return config, nil
}

func playGame(config gameConfig) (int, int) {
	return 0, 0
}

func main() {
	input, cases := initCases()

	for i := 0; i < cases; i++ {
		fmt.Printf("Case #%d: ", i+1)
		config, err := readGameConfig(input)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("config: %+v", config)
		scoreCheat, scoreNormal := playGame(config)
		fmt.Printf("%d %d\n", scoreCheat, scoreNormal)
	}
}
