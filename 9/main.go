package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

const preamble = 25

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

	lines, err := readFile(f)
	if err != nil {
		panic(err)
	}

	for i := preamble; i < len(lines); i++ {
		_, _, ok := findSum(lines[i], lines[i-preamble:i])
		if !ok {
			fmt.Println(lines[i])
			break
		}
	}
}

func findSum(n int, nums []int) (int, int, bool) {
	ns := make([]int, len(nums))
	copy(ns, nums)
	sort.Ints(ns)
	x, y := 0, len(ns)-1

	for x < y {
		z := ns[x] + ns[y]
		if z > n {
			y--
		} else if z < n {
			x++
		} else {
			return ns[x], ns[y], true
		}
	}

	return 0, 0, false
}

func readFile(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)

	lines := []int{}
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		lines = append(lines, n)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
