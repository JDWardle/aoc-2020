package main

import (
	"bytes"
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

// Valid returns true if the password contains at least Min number of Character
// but no more then Max.
func (p *Policy) Valid(password string) bool {
	var n int
	for _, r := range password {
		if string(r) == p.Character {
			n++
		}
	}
	return n >= p.Min && n <= p.Max
}

type PolicyPart2 struct {
	Pos       []int
	Character string
}

// Valid returns true if the password contains the Character only once at two
// positions.
func (p *PolicyPart2) Valid(password string) bool {
	pos1 := password[p.Pos[0]-1:p.Pos[0]] == p.Character
	pos2 := password[p.Pos[1]-1:p.Pos[1]] == p.Character

	if pos1 && pos2 {
		return false
	}

	return pos1 || pos2
}

func main() {
	flag.StringVar(&input, "input", "input.txt", "Sets the input file to load")
	flag.Parse()

	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(f, &buf)

	valid, err := readInput(tee)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part one: %d\n", valid)

	valid, err = readInputPart2(&buf)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part two: %d\n", valid)
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

func readInputPart2(r io.Reader) (valid int, err error) {
	var s scanner.Scanner
	s.Init(r)

	policy := &PolicyPart2{
		Pos: make([]int, 2),
	}

	mode := "pos1"

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		txt := s.TokenText()
		if txt == "-" || txt == ":" {
			continue
		}

		switch mode {
		case "pos1":
			n, err := strconv.Atoi(txt)
			if err != nil {
				return 0, err
			}

			policy.Pos[0] = n

			mode = "pos2"
		case "pos2":
			n, err := strconv.Atoi(txt)
			if err != nil {
				return 0, err
			}

			policy.Pos[1] = n

			mode = "char"
		case "char":
			policy.Character = txt
			mode = "pass"
		case "pass":
			if policy.Valid(txt) {
				valid++
			}
			mode = "pos1"
		}
	}

	return
}
