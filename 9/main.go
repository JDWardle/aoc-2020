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

	var n int

	for i := preamble; i < len(lines); i++ {
		_, _, ok := findSum(lines[i], lines[i-preamble:i])
		if !ok {
			n = lines[i]
			break
		}
	}

	fmt.Println("Part 1:", n)
	fmt.Println("Part 2:", findWeakness(n, lines))
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

func sum(nums []int) (n int) {
	for _, num := range nums {
		n += num
	}
	return
}

func findWeakness(n int, nums []int) int {
	l := 2

	i := 0
	for {
		if l > 50 {
			panic("too big?")
		}

		if i+l > len(nums) {
			i = 0
			l++
			continue
		}

		if n == sum(nums[i:l+i]) {
			ns := make([]int, l)
			copy(ns, nums[i:l+i])

			sort.Ints(ns)

			return ns[0] + ns[len(ns)-1]
		}
		i++
	}
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
