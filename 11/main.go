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

	layout, err := readFile(f)
	if err != nil {
		panic(err)
	}

	for {
		var changed bool
		layout, changed = checkSeating(layout)
		if !changed {
			break
		}
	}

	occupied := 0
	for row := range layout {
		for col := range layout[row] {
			if layout[row][col] == "#" {
				occupied++
			}
		}
	}

	fmt.Println(occupied)
	// newLayout, changed := checkSeating(layout)
	// fmt.Println(changed)

}

func checkSeating(layout [][]string) ([][]string, bool) {
	changed := false
	newLayout := make([][]string, len(layout))

	for i := range layout {
		newLayout[i] = make([]string, len(layout[i]))
		copy(newLayout[i], layout[i])
	}

	for row := 0; row < len(layout); row++ {
		topRow := row-1 >= 0
		bottomRow := row+1 < len(layout)

		for col := 0; col < len(layout[row]); col++ {
			if layout[row][col] == "." {
				continue
			}

			rightColumn := col+1 < len(layout[row])
			leftColumn := col-1 >= 0
			adjacentSeats := 0

			if topRow {
				if layout[row-1][col] == "#" {
					adjacentSeats++
				}

				if rightColumn {
					if layout[row-1][col+1] == "#" {
						adjacentSeats++
					}
				}

				if leftColumn {
					if layout[row-1][col-1] == "#" {
						adjacentSeats++
					}
				}
			}

			if bottomRow {
				if layout[row+1][col] == "#" {
					adjacentSeats++
				}

				if rightColumn {
					if layout[row+1][col+1] == "#" {
						adjacentSeats++
					}
				}

				if leftColumn {
					if layout[row+1][col-1] == "#" {
						adjacentSeats++
					}
				}
			}

			if rightColumn {
				if layout[row][col+1] == "#" {
					adjacentSeats++
				}
			}

			if leftColumn {
				if layout[row][col-1] == "#" {
					adjacentSeats++
				}
			}

			if adjacentSeats >= 4 && layout[row][col] == "#" {
				changed = true
				newLayout[row][col] = "L"
			} else if adjacentSeats == 0 && layout[row][col] == "L" {
				changed = true
				newLayout[row][col] = "#"
			}
		}
	}

	return newLayout, changed
}

func readFile(r io.Reader) ([][]string, error) {
	scanner := bufio.NewScanner(r)

	lines := [][]string{}
	for scanner.Scan() {
		lines = append(lines, strings.Split(scanner.Text(), ""))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
