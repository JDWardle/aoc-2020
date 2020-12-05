package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

var input string

func main() {
	flag.StringVar(&input, "input", "input.txt", "Sets the input file to load")
	flag.Parse()

	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	passes, err := readFile(f)
	if err != nil {
		panic(err)
	}

	highest := 0
	ids := []int{}
	for _, pass := range passes {
		row := findNum(pass[:len(pass)-3], 0, 127)
		column := findNum(pass[len(pass)-3:], 0, 7)

		id := seatID(row, column)
		if id > highest {
			highest = id
		}
		ids = append(ids, id)
	}
	fmt.Println(highest)

	sort.Ints(ids)

	for i, id := range ids {
		if ids[i+1] != id+1 {
			fmt.Println(id + 1)
			return
		}
	}
}

func findNum(pass string, min, max float64) int {
	for _, p := range pass {
		mid := (max - min) / 2
		switch p {
		case 'F', 'L':
			max = math.Floor(max - mid)
		case 'B', 'R':
			min = math.Ceil(mid + min)
		}

		if min == max {
			return int(min)
		}
	}

	return -1
}

func seatID(row, column int) int {
	return row*8 + column
}

func readFile(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
