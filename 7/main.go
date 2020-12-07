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

	m := map[string]string{}
	for _, line := range lines {
		s := strings.Split(line, " bags contain")
		m[strings.TrimSpace(s[0])] = s[1]
	}

	fmt.Println(contains(m, "shiny gold"))
}

func contains(m map[string]string, s string) int {
	c := 0
	for k, v := range m {
		if strings.Contains(v, s) {
			delete(m, k)
			c++

			c += contains(m, k)
		}
	}

	return c
}

func readFile(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)

	lines := []string{}
	for scanner.Scan() {
		txt := scanner.Text()

		lines = append(lines, txt)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
