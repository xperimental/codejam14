package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

type lotteryConfig struct {
	A int
	B int
	K int
}

func readLottery(input *bufio.Scanner) (lotteryConfig, error) {
	var config lotteryConfig
	if !input.Scan() {
		return config, io.EOF
	}
	line := input.Text()
	tokens := strings.Split(line, " ")
	if len(tokens) != 3 {
		return config, columnError
	}
	var err error
	config.A, err = convInt(tokens[0])
	if err != nil {
		return config, err
	}
	config.B, err = convInt(tokens[1])
	if err != nil {
		return config, err
	}
	config.K, err = convInt(tokens[2])
	if err != nil {
		return config, err
	}
	return config, nil
}

func solveLottery(config lotteryConfig) int {
	var combinations int = 0
	for a := 0; a < config.A; a++ {
		for b := 0; b < config.B; b++ {
			v := a & b
			if v < config.K {
				combinations++
			}
		}
	}
	return combinations
}

func main() {
	input, cases := initCases()

	for i := 0; i < cases; i++ {
		fmt.Printf("Case #%d: ", i+1)
		lotteryConfig, err := readLottery(input)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("config: %+v", lotteryConfig)
		combinations := solveLottery(lotteryConfig)
		fmt.Printf("%d\n", combinations)
	}
}
