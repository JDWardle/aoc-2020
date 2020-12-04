package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
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

var (
	hairColorRegex  = regexp.MustCompile(`#[0-9a-f]{6}`)
	eyeColorRegex   = regexp.MustCompile(`amb|blu|brn|gry|grn|hzl|oth`)
	passportIDRegex = regexp.MustCompile(`[0-9]{9}`)

	input string
)

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
	fmt.Println("Valid part one: ", validPassports)

	validPassports = 0
	for _, passport := range passports {
		ok, err := validatePassportPart2(passport)
		if err != nil {
			panic(err)
		}
		if ok {
			validPassports++
		}
	}

	fmt.Println("Valid part two: ", validPassports)
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

func validatePassportPart2(passport string) (bool, error) {
	s := strings.Split(passport, " ")
	m := map[string]string{}

	for i := 0; i < len(s); i++ {
		kv := strings.Split(s[i], ":")

		m[kv[0]] = kv[1]
	}

	for _, field := range requiredFields {
		if _, ok := m[field]; !ok {
			return false, nil
		}

		switch field {
		case birthYear:
			byr := m[field]

			n, err := strconv.Atoi(byr)
			if err != nil {
				return false, err
			}

			if n < 1920 || n > 2002 {
				return false, nil
			}
		case issueYear:
			iyr := m[field]

			n, err := strconv.Atoi(iyr)
			if err != nil {
				return false, err
			}

			if n < 2010 || n > 2020 {
				return false, nil
			}
		case expirationYear:
			eyr := m[field]

			n, err := strconv.Atoi(eyr)
			if err != nil {
				return false, err
			}

			if n < 2020 || n > 2030 {
				return false, nil
			}
		case height:
			hgt := m[field]

			if strings.HasSuffix(hgt, "cm") {
				n, err := strconv.Atoi(hgt[:len(hgt)-2])
				if err != nil {
					return false, err
				}

				if n < 150 || n > 193 {
					return false, err
				}
			} else if strings.HasSuffix(hgt, "in") {
				n, err := strconv.Atoi(hgt[:len(hgt)-2])
				if err != nil {
					return false, err
				}

				if n < 59 || n > 76 {
					return false, err
				}
			} else {
				return false, nil
			}
		case hairColor:
			hcl := m[field]

			if len(hcl) > 7 {
				return false, nil
			}

			if !hairColorRegex.MatchString(hcl) {
				return false, nil
			}
		case eyeColor:
			ecl := m[field]

			if !eyeColorRegex.MatchString(ecl) {
				return false, nil
			}
		case passportID:
			pid := m[field]

			if len(pid) > 9 {
				return false, nil
			}

			if !passportIDRegex.MatchString(pid) {
				return false, nil
			}
		}
	}

	return true, nil
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
