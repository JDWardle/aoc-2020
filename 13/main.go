package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

	time, buses, err := readFile(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", part1(time, buses))
	fmt.Println("Part 2:", part2(buses))
}

func part1(time int, buses []string) int {
	closest := 0
	closestBus := 0
	for _, bus := range buses {
		if bus == "x" {
			continue
		}

		n, err := strconv.Atoi(bus)
		if err != nil {
			panic(err)
		}

		c := closestNumber(time, n)
		if closest == 0 || c < closest {
			closestBus = n
			closest = c
		}
	}

	return (closest - time) * closestBus
}

func parse64(s string) int64 {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return int64(i)
}

func part2(buses []string) int64 {
	k := 0
	busInts := make([]int64, len(buses))
	for i, bus := range buses {
		if bus != "x" {
			k++
			busInts[i] = parse64(bus)
		}
	}

	t := busInts[0]

	ms := map[int64]struct{}{
		t: struct{}{},
	}

	for i := t; i < math.MaxInt64; i += t {
		for offset, id := range busInts {
			if id == 0 {
				continue
			}

			if (i+int64(offset))%id != 0 {
				break
			}

			if _, ok := ms[id]; !ok {
				t = t * id

				ms[id] = struct{}{}
			}
		}

		if len(ms) == k {
			return i
		}
	}

	return -1
}

func offset(time int64, m map[int64]int64) bool {
	for i, v := range m {
		if (time+i)%v != 0 {
			return false
		}
	}

	return true
}

func closestNumber(n, m int) int {
	n2 := 0
	q := n / m
	n1 := m * q

	if n*m > 0 {
		n2 = m * (q + 1)
	} else {
		n2 = m * (q - 1)
	}

	num := 0
	if math.Abs(float64(n-n1)) < math.Abs(float64(n-n2)) {
		num = n1
	} else {
		num = n2
	}

	if num < n {
		num += m
	}

	return num
}

func readFile(r io.Reader) (int, []string, error) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	time, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, nil, err
	}

	scanner.Scan()
	buses := strings.Split(scanner.Text(), ",")

	if err := scanner.Err(); err != nil {
		return 0, nil, err
	}

	return time, buses, nil
}
