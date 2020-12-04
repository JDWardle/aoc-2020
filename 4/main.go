package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	birthYear      = "byr"
	issueYear      = "iyr"
	expirationYear = "eyr"
	height         = "hgt"
	hairColor      = "hcl"
	eyeColor       = "ecl"
	passportID     = "pid"
	countryID      = "cid"
)

var requiredFields = []string{
	birthYear,
	issueYear,
	expirationYear,
	height,
	hairColor,
	eyeColor,
	passportID,
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

	passports, err := readFile(f)
	if err != nil {
		panic(err)
	}

	validPassports := 0
	for _, passport := range passports {
		if validatePassport(passport) {
			validPassports++
		}
	}

	fmt.Println(validPassports)
}

func validatePassport(passport string) bool {
	s := strings.Split(passport, " ")
	m := map[string]string{}

	for i := 0; i < len(s); i++ {
		kv := strings.Split(s[i], ":")

		m[kv[0]] = kv[1]
	}

	for _, required := range requiredFields {
		if _, ok := m[required]; !ok {
			return false
		}
	}

	return true
}

func readFile(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)

	passports := []string{}
	curPassport := ""
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			passports = append(passports, strings.TrimSpace(curPassport))
			curPassport = ""
		}

		curPassport += " " + scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	passports = append(passports, strings.TrimSpace(curPassport))

	return passports, nil
}
