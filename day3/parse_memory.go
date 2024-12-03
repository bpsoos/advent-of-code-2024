package day3

import (
	"bytes"
	"io"
	"strconv"
	"strings"
)

type Mul [2]int

func ParseMemory(rawMemory io.Reader) []Mul {
	memorySectionBytes, err := io.ReadAll(rawMemory)
	if err != nil {
		panic(err)
	}

	return getMuls(string(bytes.TrimSpace(memorySectionBytes)))
}

func ParseMemoryWithDosAndDonts(rawMemory io.Reader) []Mul {
	memorySectionBytes, err := io.ReadAll(rawMemory)
	if err != nil {
		panic(err)
	}
	unstripped := string(bytes.TrimSpace(memorySectionBytes))
	stripped := stripDonts(unstripped)

	return getMuls(stripped)
}

func stripDonts(rawMemory string) string {
	stripped := ""
	remaining := rawMemory
	for {
		before, after, contains := strings.Cut(remaining, "don't()")
		if !contains {
			stripped += remaining
			break
		}
		stripped += before
		_, after, contains = strings.Cut(after, "do()")
		if !contains {
			break
		}
		remaining = after
	}

	return stripped
}

func getMuls(rawMemory string) []Mul {
	muls := make([]Mul, 0)
	var (
		contains                  bool
		num1String, num2String, r string
	)
	remaining := rawMemory
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
