package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Ranges [][]int
}

func NewRule(line string) (*Rule, error) {
	rule := &Rule{}

	r := strings.Split(line, " or ")

	for _, rr := range r {
		rng := make([]int, 2)

		nums := strings.Split(rr, "-")
		n, err := strconv.Atoi(nums[0])
		if err != nil {
			return nil, err
		}
		rng[0] = n

		n, err = strconv.Atoi(nums[1])
		if err != nil {
			return nil, err
		}
		rng[1] = n

		rule.Ranges = append(rule.Ranges, rng)
	}

	return rule, nil
}

func (r *Rule) Valid(n int) bool {
	valid := false
	for _, rng := range r.Ranges {
		if rng[0] <= n && n <= rng[1] {
			valid = true
		}
	}
	return valid
}

type Ticket struct {
	Values []int
	Yours  bool
}

func NewTicket(line string) (*Ticket, error) {
	vs := strings.Split(line, ",")

	ticket := &Ticket{
		Values: make([]int, len(vs)),
	}
	for i, v := range vs {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		ticket.Values[i] = n
	}

	return ticket, nil
}

func (t *Ticket) Validate(rules map[string]Rule) []int {
	invalids := []int{}
	for _, value := range t.Values {
		valid := false
		for _, rule := range rules {
			if rule.Valid(value) {
				valid = true
			}
		}

		if !valid {
			invalids = append(invalids, value)
		}
	}

	return invalids
}

func main() {
	file := "input2.txt"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tickets, rules, err := readFileGrouped(f, []byte{'\n', '\n'})
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1:", partA(tickets, rules))
}

func partA(tickets []Ticket, rules map[string]Rule) int {
	invalids := []int{}

	for _, ticket := range tickets {
		if ticket.Yours {
			continue
		}

		invalid := ticket.Validate(rules)

		if len(invalid) > 0 {
			invalids = append(invalids, invalid...)
		}
	}

	c := 0
	for _, invalid := range invalids {
		c += invalid
	}

	return c
}

func splitByFunc(splitBy []byte) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF && len(data) == 0 {
			return 0, nil, nil
		}

		if i := bytes.Index(data, splitBy); i >= 0 {
			if i == 0 {
				return splitByFunc(splitBy)(data[len(splitBy):], atEOF)
			}
			return i + len(splitBy), dropNewline(data[:i]), nil
		}

		if atEOF {
			return len(data), dropNewline(data), nil
		}

		return 0, nil, nil
	}
}

func parseRules(rules string) (map[string]Rule, error) {
	r := strings.Split(rules, "\n")

	m := map[string]Rule{}
	for _, rr := range r {
		v := strings.Split(rr, ": ")

		rule, err := NewRule(v[1])
		if err != nil {
			return nil, err
		}
		m[v[0]] = *rule
	}
	return m, nil
}

func parseTickets(tickets string) ([]Ticket, error) {
	s := strings.ReplaceAll(tickets, "nearby tickets:\n", "")

	t := strings.Split(s, "\n")

	ts := []Ticket{}
	for _, tt := range t {
		ticket, err := NewTicket(tt)
		if err != nil {
			return nil, err
		}
		ts = append(ts, *ticket)
	}

	return ts, nil
}

func readFileGrouped(r io.Reader, groupBy []byte) ([]Ticket, map[string]Rule, error) {
	if groupBy == nil {
		groupBy = []byte{'\n', '\n'}
	}

	scanner := bufio.NewScanner(r)
	scanner.Split(splitByFunc(groupBy))

	section := "rules"

	var tickets []Ticket
	var rules map[string]Rule
	var err error
	for scanner.Scan() {
		txt := scanner.Text()

		switch section {
		case "rules":
			rules, err = parseRules(txt)
			if err != nil {
				return nil, nil, err
			}
			section = "yours"
		case "yours":
			ticket, err := NewTicket(strings.ReplaceAll(txt, "your ticket:\n", ""))
			if err != nil {
				return nil, nil, err
			}

			ticket.Yours = true

			tickets = append(tickets, *ticket)
			section = "tickets"
		case "tickets":
			ts, err := parseTickets(txt)
			if err != nil {
				return nil, nil, err
			}

			tickets = append(tickets, ts...)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return tickets, rules, nil
}

func dropNewline(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\n' {
		return data[:len(data)-1]
	}
	return data
}
