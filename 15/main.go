package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	maxTurnPartA = 2_000
	maxTurnPartB = 30_000_000
)

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

	nums, err := readFile(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(partA(nums, maxTurnPartA))
	fmt.Println(partA(nums, maxTurnPartB))
}

func partA(start []int, maxTurn int) int {
	turns := make([]int, len(start))
	copy(turns, start)
	m := map[int][]int{}

	for i, n := range start {
		m[n] = []int{i, -1}
	}

	for i := len(start); i < maxTurn; i++ {
		prev := turns[i-1]

		prevTurns := m[prev]

		x, y := prevTurns[0], prevTurns[1]

		if y == -1 {
			if _, ok := m[0]; ok {
				m[0][1] = m[0][0]
				m[0][0] = i
			} else {
				m[0] = []int{i, -1}
			}

			turns = append(turns, 0)
			continue
		}

		diff := (x + 1) - (y + 1)

		if _, ok := m[diff]; ok {
			m[diff][1] = m[diff][0]
			m[diff][0] = i
		} else {
			m[diff] = []int{i, -1}
		}

		turns = append(turns, diff)
	}

	return turns[maxTurn-1]
}

func readFile(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)

	nums := []int{}
	for scanner.Scan() {
		for _, s := range strings.Split(scanner.Text(), ",") {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}

			nums = append(nums, n)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nums, nil
}
