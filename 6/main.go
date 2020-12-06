package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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

	count := 0
	count2 := 0
	for _, line := range lines {
		count += uniqueAnsers(line)
		count2 += uniqueAnswersPart2(line)
	}
	fmt.Println("Part 1: ", count)
	fmt.Println("Part 2: ", count2)

}

func uniqueAnsers(group string) int {
	m := map[rune]struct{}{}
	for _, s := range group {
		if s == ',' {
			continue
		}
		m[s] = struct{}{}
	}

	return len(m)
}

func uniqueAnswersPart2(group string) int {
	g := strings.Split(group, ",")

	m := map[rune]int{}
	for _, s := range g {
		for _, ss := range s {
			if _, ok := m[ss]; !ok {
				m[ss] = 1
			} else {
				m[ss]++
			}
		}
	}

	c := 0
	for _, i := range m {
		if i == len(g) {
			c++
		}
	}

	return c
}

func readFile(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)

	groups := []string{}
	group := ""
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			groups = append(groups, group)
			group = ""
		}

		if group == "" {
			group += scanner.Text()
		} else {
			group += "," + scanner.Text()
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	groups = append(groups, group)

	return groups, nil
}
