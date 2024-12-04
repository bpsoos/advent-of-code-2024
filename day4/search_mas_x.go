package day4

import (
	"bytes"
	"io"
	"strings"
)

func SearchMasX(rawInput io.Reader) int {
	bytesInput, err := io.ReadAll(rawInput)
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(bytes.TrimSpace(bytesInput)), "\n")
	count := 0
	for i := 0; i < len(input)-2; i++ {
		for j := 0; j < len(input[0])-2; j++ {
			masArea := getMasArea(input, i, j)
			if isMasX(masArea) {
				count++
			}
		}
	}

	return count
}

func isMasX(masArea []string) bool {
	return isUpwardMasX(masArea) || isDownwardMasX(masArea) || isLeftMasX(masArea) || isRightMasX(masArea)
}

func isUpwardMasX(masArea []string) bool {
	diagonal := toDiagonal(masArea)
	diagonalFlipped := toDiagonal(toBackward(masArea))
	if diagonal[2] == "MAS" && diagonalFlipped[2] == "MAS" {
		return true
	}
	return false
}

func isRightMasX(masArea []string) bool {
	diagonal := toDiagonal(transpose(masArea))
	diagonalFlipped := toDiagonal(toBackward(transpose(masArea)))
	if diagonal[2] == "SAM" && diagonalFlipped[2] == "SAM" {
		return true
	}
	return false
}

func isDownwardMasX(masArea []string) bool {
	diagonal := toDiagonal(masArea)
	diagonalFlipped := toDiagonal(toBackward(masArea))
	if diagonal[2] == "SAM" && diagonalFlipped[2] == "SAM" {
		return true
	}
	return false
}
func isLeftMasX(masArea []string) bool {
	diagonal := toDiagonal(transpose(masArea))
	diagonalFlipped := toDiagonal(toBackward(transpose(masArea)))
	if diagonal[2] == "MAS" && diagonalFlipped[2] == "MAS" {
		return true
	}
	return false
}

func getMasArea(input []string, fromI, fromJ int) []string {
	var masAreaBytes [3][3]byte
	for i := range 3 {
		for j := range 3 {
			masAreaBytes[j][i] = input[fromJ+j][fromI+i]
		}
	}
	masArea := make([]string, 3)
	for i := range 3 {
		masArea[i] = string(masAreaBytes[i][:])
	}

	return masArea
}
