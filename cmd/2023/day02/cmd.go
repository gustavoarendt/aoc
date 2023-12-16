package day02

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:  "day02",
	Short: "day02",
	Long: "day02",
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
	logrus.Infof("Score - Day 02 - Step Two: %d", stepTwo(string(content)))
}

func stepOne(input string) int {
	var score int
	for _, line := range strings.Split(input, "\n") {
		gameId, isPossible := isGamePossible(line)
		if isPossible {
			score += gameId
		}
	}
	return score
}

func stepTwo(input string) int {
	var score int
	for _, line := range strings.Split(input, "\n") {
		powerSet := getPowerSets(line)
		score += powerSet
	}
	return score
}

func isGamePossible(line string) (int, bool) {
	maxRedCube := 12
	maxGreenCube := 13
	maxBlueCube := 14
	gameStringAndRound := strings.Split(line, ": ")
	gameString := gameStringAndRound[0]
	gameIdAsString := strings.Split(gameString, " ")[1]
	game := gameStringAndRound[1]
	for _, round := range strings.Split(game, "; ") {
		round = strings.Split(round, "\r")[0]
		red, hasRed, green, hasGreen, blue, hasBlue := calcRound(round)
		if hasRed && red > maxRedCube || hasGreen && green > maxGreenCube || hasBlue && blue > maxBlueCube {
			return 0, false
		}
	}
	gameId, _ := strconv.Atoi(gameIdAsString)
	return gameId, true
}

func calcRound(round string) (int, bool, int, bool, int, bool) {
	var red int
	var green int
	var blue int
	var hasRed bool
	var hasGreen bool
	var hasBlue bool
	cubes := strings.Split(round, ", ")
	for _, cube := range cubes {
		if strings.Contains(cube, "red") {
			redString, hasRedBool := strings.CutSuffix(cube, " red")
			hasRed = hasRedBool
			red, _ = strconv.Atoi(redString);
		}
		if strings.Contains(cube, "green") {
			greenString, hasGreenBool := strings.CutSuffix(cube, " green")
			hasGreen = hasGreenBool
			green, _ = strconv.Atoi(greenString);
		}
		if strings.Contains(cube, "blue") {
			blueString, hasBlueBool := strings.CutSuffix(cube, " blue")
			hasBlue = hasBlueBool
			blue, _ = strconv.Atoi(blueString);
		}
	}
	return red, hasRed, green, hasGreen, blue, hasBlue
}

func getPowerSets(line string) (int) {
	var maxRed int
	var maxGreen int
	var maxBlue int
	game := strings.Split(line, ": ")[1]
	for _, round := range strings.Split(game, "; ") {
		round = strings.Split(round, "\r")[0]
		roundRed, roundGreen, roundBlue := calcMaxCubePerRound(round)
		if roundRed > maxRed {
			maxRed = roundRed
		}
		if roundGreen > maxGreen {
			maxGreen = roundGreen
		}
		if roundBlue > maxBlue {
			maxBlue = roundBlue
		}
	}
	return maxRed * maxGreen * maxBlue
}

func calcMaxCubePerRound(round string) (int, int, int) {
	var red int
	var green int
	var blue int
	cubes := strings.Split(round, ", ")
	for _, cube := range cubes {
		if strings.Contains(cube, "red") {
			redString, hasRedBool := strings.CutSuffix(cube, " red")
			if hasRedBool {
				red, _ = strconv.Atoi(redString);
			}
		}
		if strings.Contains(cube, "green") {
			greenString, hasGreenBool := strings.CutSuffix(cube, " green")
			if hasGreenBool {
				green, _ = strconv.Atoi(greenString);
			}
		}
		if strings.Contains(cube, "blue") {
			blueString, hasBlueBool := strings.CutSuffix(cube, " blue")
			if hasBlueBool {
				blue, _ = strconv.Atoi(blueString);
			}
		}
	}
	return red,  green,  blue
}