package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
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

	lines, err := readFile(f)
	if err != nil {
		panic(err)
	}
	sort.Ints(lines)

	fmt.Println("Part 1:", part1(lines))

	fmt.Println(part2(lines))

}

func part1(nums []int) int {
	m := map[int]int{
		1: 0,
		2: 0,
		3: 1,
	}

	i := 0
	for _, n := range nums {
		m[n-i]++
		i = n
	}

	return m[1] * m[3]
}

func diffs(nums []int) []int {
	diff := []int{0}

	prev := 0
	for i := 0; i < len(nums); i++ {
		diff = append(diff, nums[i]-prev)
		prev = nums[i]
	}

	diff = append(diff, 3)

	return diff
}

func part2(nums []int) int {
	nums = append([]int{0}, nums...)

	pathsToAdapter := map[int]int{
		0: 1,
	}
	for i := 0; i < len(nums); i++ {
		currentJoltage := nums[i]
		for j := i + 1; j < len(nums) && currentJoltage+3 >= nums[j]; j++ {
			nextJoltage := nums[j]
			pathsToAdapter[nextJoltage] += pathsToAdapter[currentJoltage]
		}
	}
	last := nums[len(nums)-1]

	return pathsToAdapter[last]
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

	sort.Ints(lines)

	return lines, nil
}
