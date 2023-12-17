package day03

import (
	"fmt"
	"os"
	"strconv"
	"strings"

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

	logrus.Infof("Score - Day 03 - Step One: %d", stepOne(string(content)))
	logrus.Infof("Score - Day 03 - Step Two: %d", stepTwo(string(content)))
}

type coord struct {
	x int
	y int
}
type gridNumBuilder struct {
	sb  strings.Builder
	pos coord
}

func (gnb *gridNumBuilder) setPos(c coord) {
	gnb.pos = c
}

func (gnb *gridNumBuilder) addDigit(c rune) {
	gnb.sb.WriteRune(c)
}

func (gnb gridNumBuilder) empty() bool {
	return gnb.sb.Len() == 0
}

func (gnb *gridNumBuilder) flush() (coord, gridNum) {
	pos := gnb.pos
	gn := parseNum(gnb.sb.String())
	gnb.reset()
	return pos, gn
}

func (gnb *gridNumBuilder) reset() {
	gnb.pos = coord{}
	gnb.sb.Reset()
}

type gridNum int

func (g gridNum) bounds(pos coord) (coord, coord) {
	l := len(strconv.Itoa(int(g)))
	return coord{x: pos.x - 1, y: pos.y - 1}, coord{x: pos.x + l, y: pos.y + 1}
}

type numGrid map[coord]gridNum
type symbolGrid map[coord]rune

func contains(p coord, min coord, max coord) bool {
	return p.x >= min.x && p.y >= min.y && p.x <= max.x && p.y <= max.y
}

func parseNum(s string) gridNum {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return gridNum(n)
}

func parseGrid(lines []string) (numGrid, symbolGrid) {
	nums := numGrid{}
	symbols := symbolGrid{}
	for y, line := range lines {
		var gnb gridNumBuilder
		for x, ch := range line {
			c := coord{x: x, y: y}
			switch ch {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
				if gnb.empty() {
					gnb.setPos(c)
				}
				gnb.addDigit(ch)
				if x < len(line)-1 {
					continue
				}
			case '.':
			default:
				symbols[c] = ch
			}

			if !gnb.empty() {
				pos, gn := gnb.flush()
				nums[pos] = gn
			}
		}
	}
	return nums, symbols
}

func findParts(nums numGrid, symbols symbolGrid) (parts []int) {
	for npos, num := range nums {
		for spos := range symbols {
			min, max := num.bounds(npos)
			if contains(spos, min, max) {
				parts = append(parts, int(num))
			}
		}
	}
	return parts
}

func stepOne(input string) int {
	var score int
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSuffix(line, "\r")
	}
	nums, symbols := parseGrid(lines)
	for _, n := range findParts(nums, symbols) {
		score += n
	}
	return score
}

func stepTwo(input string) int {
	var score int
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSuffix(line, "\r")
	}
	nums, symbols := parseGrid(lines)
	for spos := range symbols {
		adjacent := []int{}
		for npos, gn := range nums {
			min, max := gn.bounds(npos)
			if contains(spos, min, max) {
				adjacent = append(adjacent, int(gn))
			}
		}
		if len(adjacent) == 2 {
			score += adjacent[0] * adjacent[1]
		}
	}
	return score
}
