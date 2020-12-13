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

	// schedules := map[int]int{}
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
		fmt.Println(bus, c)
		if closest == 0 || c < closest {
			closestBus = n
			closest = c
		}
	}

	fmt.Println(closest)
	fmt.Println(time)
	fmt.Println(closestBus)

	fmt.Println((closest - time) * closestBus)
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
