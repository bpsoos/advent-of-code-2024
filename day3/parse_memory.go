package day3

import (
	"bytes"
	"io"
	"strconv"
	"strings"
)

type Mul [2]int

func ParseMemory(rawMemory io.Reader) []Mul {
	muls := make([]Mul, 0)
	memorySectionBytes, err := io.ReadAll(rawMemory)
	if err != nil {
		panic(err)
	}
	remaining := string(bytes.TrimSpace(memorySectionBytes))
	var (
		contains                  bool
		num1String, num2String, r string
	)
	for {
		_, remaining, contains = strings.Cut(remaining, "mul(")
		if !contains {
			break
		}
		num1String, r, contains = strings.Cut(remaining, ",")
		if !contains {
			continue
		}
		num1, err := strconv.Atoi(num1String)
		if err != nil {
			continue
		}
		num2String, r, contains = strings.Cut(r, ")")
		if !contains {
			continue
		}
		num2, err := strconv.Atoi(num2String)
		if err != nil {
			continue
		}
		remaining = r
		muls = append(muls, Mul{num1, num2})
	}

	return muls
}
