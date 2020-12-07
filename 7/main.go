package main

import (
	"bufio"
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

	m := map[string]string{}
	for _, line := range lines {
		s := strings.Split(line, " bags contain")
		m[strings.TrimSpace(s[0])] = s[1]
	}

	fmt.Println(contains(m, "shiny gold"))

	mPart2 := map[string]map[string]int{}
	for _, line := range lines {
		name, bags := bag(line)
		mPart2[name] = bags
	}

	fmt.Println(containsPart2(mPart2, mPart2["shiny gold"]))
}

func bag(s string) (string, map[string]int) {
	ss := strings.Split(s, " bags contain ")
	name := ss[0]
	m := map[string]int{}

	for _, sss := range strings.Split(ss[1], ", ") {
		if strings.Contains(sss, "no") {
			continue
		}

		ssss := strings.Split(sss, " ")
		n, err := strconv.Atoi(ssss[0])
		if err != nil {
			panic(err)
		}

		name := ssss[1] + " " + ssss[2]

		m[name] = n
	}

	return name, m
}

func containedBagCount(bags map[string]int) int {
	var c int
	for _, n := range bags {
		c += n
	}
	return c
}

func containsPart2(m map[string]map[string]int, containingBag map[string]int) int {
	c := 0

	for bag, count := range containingBag {
		c += count * (1 + containsPart2(m, m[bag]))
	}

	return c
}

func contains(m map[string]string, s string) int {
	c := 0
	for k, v := range m {
		if strings.Contains(v, s) {
			delete(m, k)
			c++

			c += contains(m, k)
		}
	}

	return c
}

func readFile(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)

	lines := []string{}
	for scanner.Scan() {
		txt := scanner.Text()

		lines = append(lines, txt)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
