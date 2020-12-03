package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
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

	geology, err := readFile(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part one:")
	fmt.Println(treesInPath(geology, 3, 1))

	fmt.Println("Part two:")
	fmt.Println(
		treesInPath(geology, 1, 1) *
			treesInPath(geology, 3, 1) *
			treesInPath(geology, 5, 1) *
			treesInPath(geology, 7, 1) *
			treesInPath(geology, 1, 2),
	)
}

func treesInPath(geology [][]string, stepX int, stepY int) (trees int) {
	x := stepX
	for y := stepY; y < len(geology); y += stepY {
		location := geology[y][x]

		if location == "#" {
			trees++
		}

		x += stepX

		if x >= len(geology[y]) {
			x -= len(geology[y])
		}
	}

	return trees
}

func readFile(r io.Reader) ([][]string, error) {
	scanner := bufio.NewScanner(r)

	geology := make([][]string, 0)
	for scanner.Scan() {
		geology = append(geology, strings.Split(scanner.Text(), ""))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return geology, nil
}
