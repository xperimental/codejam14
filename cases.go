package main

import (
	"bufio"
	"io"
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
