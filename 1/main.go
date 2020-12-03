package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Entry struct {
	Num1 int
	Num2 int
}

var input string

func main() {
	flag.StringVar(&input, "input", "input.txt", "Sets the input file to load")
	flag.Parse()

	f, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	nums, err := readFile(f)
	if err != nil {
		panic(err)
	}

	if entry := find2020(nums); entry != nil {
		fmt.Printf("%d * %d = %d\n", entry.Num1, entry.Num2, entry.Num1*entry.Num2)
	}
}

func find2020(nums []int) *Entry {
	for _, i := range nums {
		for _, j := range nums {
			if i+j == 2020 {
				return &Entry{
					Num1: i,
					Num2: j,
				}
			}
		}
	}
	return nil
}

func readFile(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)

	nums := []int{}
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}

		nums = append(nums, n)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nums, nil
}
