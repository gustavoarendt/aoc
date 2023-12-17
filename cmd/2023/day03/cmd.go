package day03

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day03",
	Short: "day03",
	Long:  "day03",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	content, err := os.ReadFile(fmt.Sprintf(`cmd/%s/%s/input.txt`, parent, command))
	if err != nil {
		panic(err)
	}

	logrus.Infof("Score - Day 02 - Step One: %d", stepOne(string(content)))
	// logrus.Infof("Score - Day 02 - Step Two: %d", stepTwo(string(content)))
}

// adicionar lineNumber, preencher os números e arrays dos lineNumbers

func stepOne(input string) int {
	var score int
	numbersPositions := make(map[int][]int)
	symbolsPositions := make(map[int][]int)
	numberLineRelation := make(map[int]int)
	for lineNumber, line := range strings.Split(input, "\n") {
		var numberString string
		var firstDigitIndex int
		var lastDigitIndex int
		line = strings.Split(line, "\r")[0]
		for index, char := range line {
			if isDigit(char) {
				numberString += string(char)
				if firstDigitIndex == 0 {
					firstDigitIndex = index
				}
				lastDigitIndex = index

				if index == len(line)-1 {
					number, _ := strconv.Atoi(numberString)
					numberString = ""
					if number > 0 {
						numbersPositions[number] = append(numbersPositions[number], firstDigitIndex)
						numbersPositions[number] = append(numbersPositions[number], lastDigitIndex)
						numberLineRelation[number] = lineNumber
						firstDigitIndex = 0
					}
				}
			} else {
				number, _ := strconv.Atoi(numberString)
				numberString = ""
				if number > 0 {
					numbersPositions[number] = append(numbersPositions[number], firstDigitIndex)
					numbersPositions[number] = append(numbersPositions[number], lastDigitIndex)
					numberLineRelation[number] = lineNumber
					firstDigitIndex = 0
				}
			}
			if isValidSymbol(char) && !isDigit(char) {
				symbolsPositions[lineNumber] = append(symbolsPositions[lineNumber], index)
			}
		}
	}
	score = calculateScore(numbersPositions, symbolsPositions, numberLineRelation)
	return score
}

func isValidSymbol(char rune) bool {
	return char != 46
}

func isDigit(char rune) bool {
	return unicode.IsDigit(rune(char))
}

func calculateScore(numbersPositions, symbolsPositions map[int][]int, numberLineRelation map[int]int) int {
	var score int
	var validNumber []bool
	for number, positions := range numbersPositions {
		for _, position := range positions {
			validNumber = append(validNumber, contains(symbolsPositions, number, position, numberLineRelation[number]))
		}
		found := false
		for _, check := range validNumber {
			if check {
				found = true
				break
			}
		}
		if found {
			score += number
			validNumber = []bool{}
		}
	}
	return score
}

func contains(array map[int][]int, number, position, lineNumber int) bool {
	symbolsAbove := array[lineNumber-1]
	symbolsCurrent := array[lineNumber]
	symbolsBelow := array[lineNumber+1]
	if containsInt(symbolsAbove, position) || containsInt(symbolsCurrent, position) || containsInt(symbolsBelow, position) {
		return true
	}
	return false
}

func containsInt(array []int, position int) bool {
	for _, symbol := range array {
		if symbol == position || symbol == position-1 || symbol == position+1 {
			return true
		}
	}
	return false
}
