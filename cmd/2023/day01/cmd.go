package day01

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
	Use:  "day01",
	Short: "day01",
	Long: "day01",
	Run: func(cmd *cobra.Command, args []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	content, err := os.ReadFile(fmt.Sprintf(`cmd/%s/%s/input.txt`, parent, command))
	if err != nil {
		panic(err)
	}

	logrus.Infof("Score - Day 01 - Step One: %d", stepOne(string(content)))
	logrus.Infof("Score - Day 01 - Step Two: %d", stepTwo(string(content)))
}

func stepOne(input string) int {
	var score int
	validNumbers := "0123456789"
	for _, line := range strings.Split(input, "\n") {
		leftPosition := strings.IndexAny(line, validNumbers)
		left := string(line[leftPosition])

		rightPosition := strings.LastIndexAny(line, validNumbers)
		right := string(line[rightPosition])

		compositeNumber := string(left + right)
		composite, err := strconv.Atoi(compositeNumber); if err == nil {
			score += composite
		}
	}
	return score
}

func stepTwo(input string) int {
	var score int
	for _, line := range strings.Split(input, "\n") {
		numbersOnly := []string{}

		for idx, character := range line {
			if unicode.IsDigit(character) {
				numbersOnly = append(numbersOnly, string(character))
			} else {
				if digitStr, found := hasDigitPrefix(line[idx:]); found {
					numbersOnly = append(numbersOnly, digitStr)
				}
			}
		}

		concat := numbersOnly[0] + numbersOnly[len(numbersOnly)-1]

		result, err := strconv.Atoi(concat)
		if err != nil {
			panic(err)
		}
		score += result
	}
	return score
}

func hasDigitPrefix(line string) (string, bool) {
	digits := []string{
		"zero", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine",
	}

	for i, digitAsString := range digits {
		if strings.HasPrefix(line, digitAsString) {
			return strconv.Itoa(i), true
		}
	}
	return "", false
}