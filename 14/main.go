package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var memRegex = regexp.MustCompile(`\[(\d+)\] = (\d+)`)

func main() {
	file := "input.txt"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	program, err := readFile(f)
	if err != nil {
		panic(err)
	}

	m := map[int]uint64{}
	for _, instruction := range program {
		m[instruction.Address] = instruction.Mask.Apply(instruction.Value)
	}

	var c uint64
	for _, v := range m {
		c += v
	}
	fmt.Println(c)
}

type Mask struct {
	s      string
	Ones   uint64
	Zeroes uint64
}

func (m Mask) Apply(i int) uint64 {
	v := uint64(i)

	v |= m.Ones
	return m.Zeroes & v
}

func (m Mask) String() string { return m.s }

func NewMask(mask string) (*Mask, error) {
	var err error
	m := &Mask{s: mask}

	m.Ones, err = strconv.ParseUint(strings.ReplaceAll(mask, "X", "0"), 2, 36)
	if err != nil {
		return nil, err
	}

	m.Zeroes, err = strconv.ParseUint(strings.ReplaceAll(mask, "X", "1"), 2, 36)
	if err != nil {
		return nil, err
	}

	return m, nil
}

type Instruction struct {
	Mask    Mask
	Address int
	Value   int
}

func readFile(r io.Reader) ([]Instruction, error) {
	scanner := bufio.NewScanner(r)

	program := []Instruction{}
	var curMask *Mask
	for scanner.Scan() {
		txt := scanner.Text()

		if strings.HasPrefix(txt, "mask") {
			mask, err := NewMask(txt[7:])
			if err != nil {
				return nil, err
			}
			curMask = mask
		} else {
			matches := memRegex.FindStringSubmatch(txt)

			address, err := strconv.Atoi(matches[1])
			if err != nil {
				return nil, err
			}

			value, err := strconv.Atoi(matches[2])
			if err != nil {
				return nil, err
			}

			program = append(program, Instruction{
				Mask:    *curMask,
				Address: address,
				Value:   value,
			})
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return program, nil
}
