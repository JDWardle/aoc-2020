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

	layoutPart1, changed := checkSeating(layout)
	for {
		layoutPart1, changed = checkSeating(layoutPart1)
		if !changed {
			break
		}
	}

	occupied := 0
	for row := range layoutPart1 {
		for col := range layoutPart1[row] {
			if layoutPart1[row][col] == "#" {
				occupied++
			}
		}
	}

	fmt.Println("Part 1:", occupied)

	layoutPart2, changed := checkSeatingPart2(layout)
	for {
		layoutPart2, changed = checkSeatingPart2(layoutPart2)
		if !changed {
			break
		}
	}

	occupied = 0
	for row := range layoutPart2 {
		for col := range layoutPart2[row] {
			if layoutPart2[row][col] == "#" {
				occupied++
			}
		}
	}

	fmt.Println("Part 2:", occupied)

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

func trace(dirRow, dirCol, startRow, startCol int, layout [][]string) bool {
	row := startRow
	col := startCol

	for {
		row += dirRow
		col += dirCol

		if row < 0 || col < 0 {
			break
		}

		if row > len(layout)-1 || col > len(layout[row])-1 {
			break
		}

		if layout[row][col] == "L" {
			return false
		}

		if layout[row][col] == "#" {
			return true
		}
	}

	return false
}

func checkSeatingPart2(layout [][]string) ([][]string, bool) {
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
			visibleSeats := 0

			if topRow {
				if trace(-1, 0, row, col, layout) {
					visibleSeats++
				}

				if rightColumn {
					if trace(-1, +1, row, col, layout) {
						visibleSeats++
					}
				}

				if leftColumn {
					if trace(-1, -1, row, col, layout) {
						visibleSeats++
					}
				}
			}

			if bottomRow {
				if trace(+1, 0, row, col, layout) {
					visibleSeats++
				}

				if rightColumn {
					if trace(+1, +1, row, col, layout) {
						visibleSeats++
					}
				}

				if leftColumn {
					if trace(+1, -1, row, col, layout) {
						visibleSeats++
					}
				}
			}

			if rightColumn {
				if trace(0, +1, row, col, layout) {
					visibleSeats++
				}
			}

			if leftColumn {
				if trace(0, -1, row, col, layout) {
					visibleSeats++
				}
			}

			if visibleSeats >= 5 && layout[row][col] == "#" {
				changed = true
				newLayout[row][col] = "L"
			} else if visibleSeats == 0 && layout[row][col] == "L" {
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
