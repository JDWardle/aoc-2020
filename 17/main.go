package main

// During a cycle, all cubes simultaneously change their state according to the
// following rules:

// If a cube is active and exactly 2 or 3 of its neighbors are also active, the
// cube remains active. Otherwise, the cube becomes inactive. If a cube is
// inactive but exactly 3 of its neighbors are active, the cube becomes active.
// Otherwise, the cube remains inactive.

// The engineers responsible for this experimental energy source would like you
// to simulate the pocket dimension and determine what the configuration of
// cubes should be at the end of the six-cycle boot process.

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Vector struct {
	X int
	Y int
	Z int
}

func NewVector(x, y, z int) Vector { return Vector{X: x, Y: y, Z: z} }

func (v Vector) String() string { return fmt.Sprintf("(%d, %d, %d)", v.X, v.Y, v.Z) }

type Cube struct {
	Active   bool
	Position Vector
}

func NewCube(active bool, pos Vector) Cube {
	return Cube{
		Active:   active,
		Position: pos,
	}
}

func (c Cube) String() string {
	if c.Active {
		return "#"
	} else {
		return "."
	}
}

type Map struct {
	m      map[Vector]Cube
	MaxPos Vector
	MinPos Vector
}

func NewMap(lines []string) Map {
	m := Map{
		m: make(map[Vector]Cube),
	}

	y := 0
	x := 0
	for _, line := range lines {
		s := strings.Split(line, "")
		x = 0

		for _, c := range s {
			v := NewVector(x, y, 0)

			switch c {
			case "#":
				m.m[v] = NewCube(true, v)
			case ".":
				m.m[v] = NewCube(false, v)
			}
			x++
		}
		y++
	}

	m.MinPos = NewVector(0, 0, -1)
	m.MaxPos = NewVector(x-1, y-1, 1)

	return m
}

func (m Map) String() string {
	s := []string{}
	for z := m.MinPos.Z; z <= m.MaxPos.Z; z++ {
		s = append(s, fmt.Sprintf("z=%d", z))

		for y := m.MinPos.Y; y <= m.MaxPos.Y; y++ {
			ys := ""
			for x := m.MinPos.X; x <= m.MaxPos.X; x++ {
				ys += m.m[NewVector(x, y, z)].String()
			}
			s = append(s, ys)
		}

		s = append(s, "")
	}

	return strings.Join(s, "\n")
}

func (m Map) Cycle() Map {
	newMap := Map{
		m:      make(map[Vector]Cube),
		MinPos: NewVector(m.MinPos.X-1, m.MinPos.Y-1, m.MinPos.Z-1),
		MaxPos: NewVector(m.MaxPos.X+1, m.MaxPos.Y+1, m.MaxPos.Z+1),
	}

	for z := m.MinPos.Z; z <= m.MaxPos.Z; z++ {
		for y := m.MinPos.Y; y <= m.MaxPos.Y; y++ {
			for x := m.MinPos.X; x <= m.MaxPos.X; x++ {
				v := NewVector(x, y, z)
				cube, ok := m.m[v]
				if !ok {
					cube.Position = v
				}

				activeCubes := len(m.ActiveAdjacentCubes(cube))

				active := false

				if cube.Active {
					if activeCubes == 2 || activeCubes == 3 {
						active = true
					}
				} else {
					if activeCubes == 3 {
						active = true
					}
				}

				newMap.m[v] = NewCube(active, v)
			}
		}
	}

	return newMap
}

func (m Map) ActiveAdjacentCubes(cube Cube) []Cube {
	adjacentCubes := m.AdjacentCubes(cube)

	n := 0

	for _, cube := range adjacentCubes {
		if cube.Active {
			adjacentCubes[n] = cube
			n++
		}
	}

	return adjacentCubes[:n]
}

func (m Map) AdjacentCubes(cube Cube) []Cube {
	positions := m.AdjacentVectors(cube.Position)

	cubes := []Cube{}
	for _, pos := range positions {
		if pos.X < m.MinPos.X || pos.Y < m.MinPos.Y || pos.Z < m.MinPos.Z {
			continue
		}

		if pos.X > m.MaxPos.X || pos.Y > m.MaxPos.Y || pos.Z > m.MaxPos.Z {
			continue
		}

		cubes = append(cubes, m.m[pos])
	}

	return cubes
}

func (m Map) AdjacentVectors(pos Vector) [26]Vector {
	minPos := NewVector(pos.X-1, pos.Y-1, pos.Z-1)
	maxPos := NewVector(pos.X+1, pos.Y+1, pos.Z+1)

	adjacent := [26]Vector{}

	i := 0
	for z := minPos.Z; z <= maxPos.Z; z++ {
		for y := minPos.Y; y <= maxPos.Y; y++ {
			for x := minPos.X; x <= maxPos.X; x++ {
				v := NewVector(x, y, z)

				if v == pos {
					continue
				}

				adjacent[i] = v
				i++
			}
		}
	}

	return adjacent
}

func (m Map) ActiveCubes() (count int) {
	for _, cube := range m.m {
		if cube.Active {
			count++
		}
	}
	return
}

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

	m := NewMap(lines)

	for cycle := 0; cycle < 6; cycle++ {
		fmt.Printf("------------------------ Cycle %d ------------------------\n", cycle)
		m = m.Cycle()
		fmt.Println(m)
	}

	fmt.Println("Part 1:", m.ActiveCubes())
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
