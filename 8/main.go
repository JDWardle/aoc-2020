package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
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

	acc, err := compute(lines)
	if err != nil {
		panic(err)
	}
	fmt.Println("Part 1:", acc)
}

func compute(instructions []string) (accumulator int, err error) {
	m := map[int]struct{}{}

	i := 0

	for i < len(instructions) {
		if _, ok := m[i]; ok {
			return
		}

		cmd := instructions[i][:3]
		v := instructions[i][4:]

		switch cmd {
		case "acc":
			acc, err := strconv.Atoi(v)
			if err != nil {
				return 0, err
			}

			accumulator += acc
		case "jmp":
			jmp, err := strconv.Atoi(v)
			if err != nil {
				return 0, err
			}

			i += jmp
			continue
		}

		m[i] = struct{}{}
		i++
	}

	return
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

func splitByFunc(splitBy []byte) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := bytes.Index(data, splitBy); i >= 0 {
			if i == 0 {
				return splitByFunc(splitBy)(data[len(splitBy):], atEOF)
			}
			return i + len(splitBy), dropNewline(data[:i]), nil
		}

		if atEOF {
			return len(data), dropNewline(data), nil
		}

		return 0, nil, nil
	}
}

func readFileGrouped(r io.Reader, groupBy []byte) ([][]string, error) {
	if groupBy == nil {
		groupBy = []byte{'\n', '\n'}
	}

	scanner := bufio.NewScanner(r)
	scanner.Split(splitByFunc(groupBy))

	var groups [][]string
	for scanner.Scan() {
		groups = append(groups, strings.Split(scanner.Text(), "\n"))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}

func dropNewline(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\n' {
		return data[:len(data)-1]
	}
	return data
}
