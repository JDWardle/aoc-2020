package main

// NOTE: Test input gives incorrect result but works for actual input...

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Vector4 struct {
	X, Y, Z, W int
}

func NewVector4(x, y, z, w int) Vector4 { return Vector4{X: x, Y: y, Z: z, W: w} }

func (v Vector4) String() string { return fmt.Sprintf("(%d, %d, %d, %d)", v.X, v.Y, v.Z, v.W) }

type CubePart2 struct {
	Active   bool
	Position Vector4
}

func NewCubePart2(active bool, pos Vector4) CubePart2 {
	return CubePart2{
		Active:   active,
		Position: pos,
	}
}

func (c CubePart2) String() string {
	if c.Active {
		return "#"
	}
	return "."
}

type MapPart2 struct {
	m      map[Vector4]CubePart2
	MaxPos Vector4
	MinPos Vector4
}

func NewMapPart2(lines []string) MapPart2 {
	m := MapPart2{
		m: make(map[Vector4]CubePart2),
	}

	y := 0
	x := 0
	for _, line := range lines {
		s := strings.Split(line, "")
		x = 0

		for _, c := range s {
			v := NewVector4(x, y, 0, 0)

			switch c {
			case "#":
				m.m[v] = NewCubePart2(true, v)
			case ".":
				m.m[v] = NewCubePart2(false, v)
			}
			x++
		}
		y++
	}

	m.MinPos = NewVector4(-1, -1, -1, -1)
	m.MaxPos = NewVector4(x-1, y-1, 1, 1)

	return m
}

func (m MapPart2) String() string {
	s := []string{}

	for w := m.MinPos.W; w <= m.MaxPos.W; w++ {
		for z := m.MinPos.Z; z <= m.MaxPos.Z; z++ {
			s = append(s, fmt.Sprintf("z=%d, w=%d", z, w))

			for y := m.MinPos.Y; y <= m.MaxPos.Y; y++ {
				ys := ""
				for x := m.MinPos.X; x <= m.MaxPos.X; x++ {
					ys += m.m[NewVector4(x, y, z, w)].String()
				}
				s = append(s, ys)
			}

			s = append(s, "")
		}
	}

	return strings.Join(s, "\n")
}

func (m MapPart2) Cycle() MapPart2 {
	newMap := MapPart2{
		m:      make(map[Vector4]CubePart2),
		MinPos: NewVector4(m.MinPos.X-1, m.MinPos.Y-1, m.MinPos.Z-1, m.MinPos.W-1),
		MaxPos: NewVector4(m.MaxPos.X+1, m.MaxPos.Y+1, m.MaxPos.Z+1, m.MaxPos.W+1),
	}

	for w := m.MinPos.W; w <= m.MaxPos.W; w++ {
		for z := m.MinPos.Z; z <= m.MaxPos.Z; z++ {
			for y := m.MinPos.Y; y <= m.MaxPos.Y; y++ {
				for x := m.MinPos.X; x <= m.MaxPos.X; x++ {
					v := NewVector4(x, y, z, w)
					cube := m.Cube(v)

					activeCubes := len(m.ActiveAdjacentCubes(cube))
					if activeCubes == 0 {
						continue
					}

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

					newMap.m[v] = NewCubePart2(active, v)
				}
			}
		}
	}

	return newMap
}

func (m MapPart2) ActiveAdjacentCubes(cube CubePart2) []CubePart2 {
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

func (m MapPart2) Cube(pos Vector4) CubePart2 {
	cube, ok := m.m[pos]
	if !ok {
		cube.Position = pos
	}

	return cube
}

func (m MapPart2) AdjacentCubes(cube CubePart2) []CubePart2 {
	positions := m.AdjacentVectors(cube.Position)

	cubes := []CubePart2{}
	for _, pos := range positions {
		if pos.X < m.MinPos.X || pos.Y < m.MinPos.Y || pos.Z < m.MinPos.Z || pos.W < m.MinPos.W {
			continue
		}

		if pos.X > m.MaxPos.X || pos.Y > m.MaxPos.Y || pos.Z > m.MaxPos.Z || pos.W > m.MaxPos.W {
			continue
		}

		cubes = append(cubes, m.Cube(pos))
	}

	return cubes
}

func (m MapPart2) AdjacentVectors(pos Vector4) [80]Vector4 {
	minPos := NewVector4(pos.X-1, pos.Y-1, pos.Z-1, pos.W-1)
	maxPos := NewVector4(pos.X+1, pos.Y+1, pos.Z+1, pos.W+1)

	adjacent := [80]Vector4{}

	i := 0
	for w := minPos.W; w <= maxPos.W; w++ {
		for z := minPos.Z; z <= maxPos.Z; z++ {
			for y := minPos.Y; y <= maxPos.Y; y++ {
				for x := minPos.X; x <= maxPos.X; x++ {
					v := NewVector4(x, y, z, w)

					if v == pos {
						continue
					}

					adjacent[i] = v
					i++
				}
			}
		}
	}

	return adjacent
}

func (m MapPart2) ActiveCubes() (count int) {
	for _, cube := range m.m {
		if cube.Active {
			count++
		}
	}
	return
}

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
	}
	return "."
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

	m.MinPos = NewVector(-1, -1, -1)
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

				if activeCubes == 0 {
					continue
				}

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
	mPart2 := NewMapPart2(lines)

	for cycle := 0; cycle < 6; cycle++ {
		m = m.Cycle()
		mPart2 = mPart2.Cycle()
	}

	fmt.Println("Part 1:", m.ActiveCubes())
	fmt.Println("Part 2:", mPart2.ActiveCubes())
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
