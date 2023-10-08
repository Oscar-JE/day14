package parse

import (
	"bufio"
	"day14/common"
	"os"
	"strconv"
	"strings"
)

func Parse(filepath string) []common.ClosedInterval {
	intervals := []common.ClosedInterval{}
	file, err := os.Open(filepath)
	if err != nil {
		panic("error rading file")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input_line := scanner.Text()
		lineIntervals := parseIntervals(input_line)
		intervals = append(intervals, lineIntervals...)
	}
	return intervals
}

func parseIntervals(input_line string) []common.ClosedInterval {
	pointReps := strings.Split(input_line, " -> ")
	intervals := []common.ClosedInterval{}
	for i := 0; i < len(pointReps)-1; i++ {
		firstPoint := convertToPoint(pointReps[i])
		secondPoint := convertToPoint(pointReps[i+1])
		intervals = append(intervals, common.Init(firstPoint, secondPoint))
	}
	return intervals
}

func convertToPoint(pointRep string) common.Point {
	pointRep = strings.Trim(pointRep, " ")
	numbersReps := strings.Split(pointRep, ",")
	row := convertToNumber(numbersReps[0])
	col := convertToNumber(numbersReps[1])
	return common.InitPoint(row, col)
}

func convertToNumber(nrRep string) int {
	nr, errorRow := strconv.Atoi(nrRep)
	if errorRow != nil {
		panic("number pars failed")
	}
	return nr
}
