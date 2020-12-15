package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var memRegex = regexp.MustCompile(`\[(\d+)\] = (\d+)`)

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

	program, err := readFile(f)
	if err != nil {
		panic(err)
	}

	m := map[int]uint64{}
	for _, instruction := range program {
		m[instruction.Address] = instruction.Mask.Apply(instruction.Value)
	}

	var c uint64
	for _, v := range m {
		c += v
	}
	fmt.Println(c)

	mPart2 := map[uint64]int{}
	for _, instruction := range program {
		mask, err := NewMaskPart2(applyMask(instruction.Mask.s, uint64(instruction.Address)))
		if err != nil {
			panic(err)
		}
		for _, mask := range mask.bitMasks {
			mPart2[mask] = instruction.Value
		}
	}

	var cPart2 int

	for _, v := range mPart2 {
		cPart2 += v
	}

	fmt.Println(cPart2)
}

type MaskPart2 struct {
	s       string
	bitMask uint64

	bitMasks []uint64
}

func applyMask(mask string, address uint64) string {
	m := strings.ReplaceAll(mask, "X", "0")
	bitMask, err := strconv.ParseUint(m, 2, 36)
	if err != nil {
		panic(err)
	}

	newMask := fmt.Sprintf("%036b", bitMask|address)

	for i, s := range mask {
		if s == 'X' {
			newMask = newMask[:i] + "X" + newMask[i+1:]
		}
	}
	return newMask
}

func parseMasks(mask string) (addrs []uint64, err error) {
	addr1, err := strconv.ParseUint(strings.ReplaceAll(mask, "X", "0"), 2, 36)
	if err != nil {
		return
	}

	addr2, err := strconv.ParseUint(strings.ReplaceAll(mask, "X", "1"), 2, 36)
	if err != nil {
		return
	}

	addrs = append(addrs, addr1, addr2)

	if i := strings.Index(mask, "X"); i != -1 {
		a, err := parseMasks(mask[:i] + "0" + mask[i+1:])
		if err != nil {
			return nil, err
		}

		addrs = append(addrs, a...)

		a2, err := parseMasks(mask[:i] + "1" + mask[i+1:])
		if err != nil {
			return nil, err
		}

		addrs = append(addrs, a2...)
	}

	u := map[uint64]struct{}{}
	l := []uint64{}

	for _, mask := range addrs {
		if _, ok := u[mask]; !ok {
			u[mask] = struct{}{}
			l = append(l, mask)
		}
	}

	return l, nil
}

func NewMaskPart2(mask string) (*MaskPart2, error) {
	m := &MaskPart2{s: mask}

	var err error
	m.bitMask, err = strconv.ParseUint(strings.ReplaceAll(mask, "X", "0"), 2, 36)
	if err != nil {
		return nil, err
	}

	masks, err := parseMasks(mask)
	if err != nil {
		return nil, err
	}

	m.bitMasks = masks

	return m, nil
}

type Mask struct {
	s      string
	Ones   uint64
	Zeroes uint64
}

func (m Mask) Apply(i int) uint64 {
	v := uint64(i)

	v |= m.Ones
	return m.Zeroes & v
}

func (m Mask) String() string { return m.s }

func NewMask(mask string) (*Mask, error) {
	var err error
	m := &Mask{s: mask}

	m.Ones, err = strconv.ParseUint(strings.ReplaceAll(mask, "X", "0"), 2, 36)
	if err != nil {
		return nil, err
	}

	m.Zeroes, err = strconv.ParseUint(strings.ReplaceAll(mask, "X", "1"), 2, 36)
	if err != nil {
		return nil, err
	}

	return m, nil
}

type Instruction struct {
	Mask      Mask
	MaskPart2 MaskPart2
	Address   int
	Value     int
}

func readFile(r io.Reader) ([]Instruction, error) {
	scanner := bufio.NewScanner(r)

	program := []Instruction{}
	var curMask *Mask
	// var curMaskPart2 *MaskPart2
	for scanner.Scan() {
		txt := scanner.Text()

		if strings.HasPrefix(txt, "mask") {
			mask, err := NewMask(txt[7:])
			if err != nil {
				return nil, err
			}

			// maskPart2, err := NewMaskPart2(txt[7:])
			// if err != nil {
			// 	return nil, err
			// }

			curMask = mask
			// curMaskPart2 = maskPart2
		} else {
			matches := memRegex.FindStringSubmatch(txt)

			address, err := strconv.Atoi(matches[1])
			if err != nil {
				return nil, err
			}

			value, err := strconv.Atoi(matches[2])
			if err != nil {
				return nil, err
			}

			program = append(program, Instruction{
				Mask: *curMask,
				// MaskPart2: *curMaskPart2,
				Address: address,
				Value:   value,
			})
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return program, nil
}
