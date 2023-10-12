package main

import (
	"day14/common"
	"day14/parse"
	"day14/particle"
	"day14/reservoir"
	"fmt"
)

func main() {
	intervals := parse.Parse("input_short.txt")
	//intervals := parse.Parse("input.txt")
	reservoir := reservoir.InitFromIntervalls(intervals)
	startPosition := common.InitPoint(500, 0)
	overfilled := false
	for !overfilled {
		sand := particle.Init(startPosition)
		for true {
			nextPoints := sand.FallingPatern()
			if reservoir.IsAllBlocked(nextPoints) {
				reservoir.SetSand(sand.GetPosition())
				break
			}
			nextPoint := reservoir.FirstNoneBlocked(nextPoints)
			if !reservoir.OnField(nextPoint) {
				overfilled = true
				break
			}
			sand.SetPosition(nextPoint)
		}
	}
	fmt.Println(reservoir)
	fmt.Println(reservoir.GetSandCount())
}
