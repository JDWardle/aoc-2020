package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"text/scanner"
)

var input string

type Policy struct {
	Min       int
	Max       int
	Character string
}

func (p *Policy) Valid(password string) bool {
	var n int
	for _, r := range password {
		if string(r) == p.Character {
			n++
		}
	}
	return n >= p.Min && n <= p.Max
}

func main() {
	flag.StringVar(&input, "input", "input.txt", "Sets the input file to load")
	flag.Parse()

	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	valid, err := readInput(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(valid)
}

func readInput(r io.Reader) (valid int, err error) {
	var s scanner.Scanner
	s.Init(r)

	policy := &Policy{}

	mode := "min"

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		txt := s.TokenText()
		if txt == "-" || txt == ":" {
			continue
		}

		switch mode {
		case "min":
			policy.Min, err = strconv.Atoi(txt)
			if err != nil {
				return 0, err
			}

			mode = "max"
		case "max":
			policy.Max, err = strconv.Atoi(txt)
			if err != nil {
				return 0, err
			}

			mode = "char"
		case "char":
			policy.Character = txt
			mode = "pass"
		case "pass":
			if policy.Valid(txt) {
				valid++
			}
			mode = "min"
		}
	}

	return
}
