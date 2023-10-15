package main

import (
	"day14/common"
	"day14/parse"
	"day14/particle"
	"day14/reservoir"
	"fmt"
)

func main() {
	//intervals := parse.Parse("input_short.txt")
	intervals := parse.Parse("input.txt")
	startPosition := common.InitPoint(500, 0)
	reservoir := reservoir.InitPart2(intervals, startPosition.GetRow())
	overfilled := false
	for !overfilled {
		sand := particle.Init(startPosition)
		for true {
			nextPoints := sand.FallingPatern()
			if reservoir.IsAllBlocked(nextPoints) {
				reservoir.SetSand(sand.GetPosition())
				if reservoir.IsBlocked(startPosition) {
					overfilled = true
				}
				break
			}
			nextPoint := reservoir.FirstNoneBlocked(nextPoints)
			sand.SetPosition(nextPoint)
		}
	}
	fmt.Println(reservoir)
	fmt.Println(reservoir.GetSandCount())
}
