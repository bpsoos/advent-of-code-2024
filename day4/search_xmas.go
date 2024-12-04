package day4

import (
	"bytes"
	"io"
	"strings"
)

func SearchXmas(rawInput io.Reader) int {
	bytesInput, err := io.ReadAll(rawInput)
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(bytes.TrimSpace(bytesInput)), "\n")

	count := 0
	count += searchHorizontal(input)
	count += searchHorizontalBackward(input)
	count += searchVertical(input)
	count += searchVerticalBackward(input)
	count += searchDiagonal(input)
	count += searchDiagonalBackward(input)
	count += searchDiagonalFlipped(input)
	count += searchDiagonalFlippedBackward(input)

	return count
}

func searchHorizontal(input []string) int {
	return countOccurences(input)
}
func searchHorizontalBackward(input []string) int {
	return countOccurences(toBackward(input))
}
func searchVertical(input []string) int {
	return countOccurences(transpose(input))
}
func searchVerticalBackward(input []string) int {
	return countOccurences(toBackward(transpose(input)))
}

func searchDiagonal(input []string) int {
	diagonal := toDiagonal(input)
	return countOccurences(diagonal)
}

func searchDiagonalBackward(input []string) int {
	diagonal := toBackward(toDiagonal(input))
	return countOccurences(diagonal)
}
func searchDiagonalFlipped(input []string) int {
	diagonal := toDiagonal(toBackward(input))
	return countOccurences(diagonal)
}
func searchDiagonalFlippedBackward(input []string) int {
	diagonal := toBackward(toDiagonal(toBackward(input)))
	return countOccurences(diagonal)
}

func countOccurences(input []string) int {
	count := 0
	for i := range input {
		count += strings.Count(input[i], "XMAS")
	}

	return count
}

func transpose(input []string) []string {
	transposed := make([][]byte, len(input[0]))
	for i := range transposed {
		transposed[i] = make([]byte, len(input))
	}

	for i := range input {
		for j := range input[i] {
			transposed[j][i] = input[i][j]
		}
	}

	transposedStrings := make([]string, len(input[0]))
	for i := range transposed {
		transposedStrings[i] = string(transposed[i])
	}

	return transposedStrings
}

func toBackward(input []string) []string {
	backwards := make([]string, len(input))
	for i := range input {
		backward := make([]byte, len(input[i]))
		for j := range input[i] {
			backward[len(input[i])-1-j] = input[i][j]
		}
		backwards[i] = string(backward)
	}

	return backwards
}

func toDiagonal(input []string) []string {
	i := 0
	iStart := 0
	j := 0
	jStart := 0
	diagonal := make([]string, 0)
	row := make([]byte, 0)
	for {
		row = append(row, input[j][i])
		if i == len(input[0])-1 {
			diagonal = append(diagonal, string(row))
			if j == len(input)-1 {
				break
			}
			row = make([]byte, 0)
			iStart++
			j = jStart
			i = iStart
			continue
		}
		if j == 0 {
			diagonal = append(diagonal, string(row))
			row = make([]byte, 0)
			jStart++
			i = 0
			j = jStart
			continue
		}
		i++
		j--
	}
	return diagonal
}
