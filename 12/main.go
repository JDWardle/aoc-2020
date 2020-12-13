package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

type Ship struct {
	Dir  int
	PosX int
	PosY int
}

func NewShip() *Ship {
	return &Ship{Dir: 90}
}

func (s *Ship) ManhattanDistance() int {
	x, y := float64(s.PosX), float64(s.PosY)

	dist := math.Abs(x) + math.Abs(y)
	return int(dist)
}

func (s *Ship) Rotate(deg int) {
	dir := float64(s.Dir + deg)

	if dir > 360 {
		dir -= 360
	}

	if dir <= 0 {
		dir += 360
	}

	s.Dir = int(dir)
}

func (s *Ship) Direction() string {
	switch s.Dir {
	case 90:
		return "E"
	case 180:
		return "S"
	case 270:
		return "W"
	case 360:
		return "N"
	default:
		return ""
	}
}

func (s *Ship) Navigate(action string, units int) {
	switch action {
	case "N":
		s.PosY += units
	case "S":
		s.PosY -= units
	case "E":
		s.PosX += units
	case "W":
		s.PosX -= units
	case "L":
		s.Rotate(-units)
	case "R":
		s.Rotate(units)
	case "F":
		s.Navigate(s.Direction(), units)
	}
}

type ShipPart2 struct {
	PosX      int
	PosY      int
	WaypointX int
	WaypointY int
}

func NewShipPart2() *ShipPart2 {
	return &ShipPart2{
		WaypointX: 10,
		WaypointY: 1,
	}
}

func degreeToRad(deg int) float64 {
	return float64(deg) * math.Pi / 180
}

func (s *ShipPart2) MoveTowardsWaypoint(units int) {
	s.PosX += s.WaypointX * units
	s.PosY += s.WaypointY * units
}

func (s *ShipPart2) Navigate(action string, units int) {
	switch action {
	case "N":
		s.WaypointY += units
	case "S":
		s.WaypointY -= units
	case "E":
		s.WaypointX += units
	case "W":
		s.WaypointX -= units
	case "L":
		s.RotateWaypoint(units)
	case "R":
		s.RotateWaypoint(-units)
	case "F":
		s.MoveTowardsWaypoint(units)
	}
}

func (s *ShipPart2) RotateWaypoint(deg int) {
	rad := degreeToRad(deg)
	sine := math.Sin(rad)
	cos := math.Cos(rad)

	x := float64(s.WaypointX)
	y := float64(s.WaypointY)

	s.WaypointX = int(math.Round(x*cos - y*sine))
	s.WaypointY = int(math.Round(x*sine + y*cos))
}

func (s *ShipPart2) ManhattanDistance() int {
	x, y := float64(s.PosX), float64(s.PosY)

	dist := math.Abs(x) + math.Abs(y)
	return int(dist)
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

	s, err := navigate(lines)
	if err != nil {
		panic(err)
	}
	fmt.Println("Manhattan Distance Part 1:", s.ManhattanDistance())

	sh, err := navigatePart2(lines)
	if err != nil {
		panic(err)
	}
	fmt.Println("Manhattan Distance Part 2:", sh.ManhattanDistance())
}

func navigate(instructions []string) (*Ship, error) {
	ship := NewShip()
	for _, instruction := range instructions {
		action := instruction[:1]
		units, err := strconv.Atoi(instruction[1:])
		if err != nil {
			return nil, err
		}
		ship.Navigate(action, units)

		fmt.Printf("%s\tPos: %d,%d\t Dir: %d %s\n", instruction, ship.PosX, ship.PosY, ship.Dir, ship.Direction())
	}

	return ship, nil
}

func navigatePart2(instructions []string) (*ShipPart2, error) {
	ship := NewShipPart2()
	for _, instruction := range instructions {
		action := instruction[:1]
		units, err := strconv.Atoi(instruction[1:])
		if err != nil {
			return nil, err
		}
		ship.Navigate(action, units)

		fmt.Printf("%s\tPos: %d,%d\tWpt: %d,%d\n", instruction, ship.PosX, ship.PosY, ship.WaypointX, ship.WaypointY)
	}

	return ship, nil
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
