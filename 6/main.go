package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
	for _, line := range lines {
		count += uniqueAnsers(line)
	}
	fmt.Println(count)
}

func uniqueAnsers(group string) int {
	m := map[rune]struct{}{}
	for _, s := range group {
		m[s] = struct{}{}
	}

	return len(m)
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

		group += scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	groups = append(groups, group)

	return groups, nil
}
