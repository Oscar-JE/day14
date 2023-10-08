package main

import (
	"day14/parse"
	"day14/reservoir"
	"fmt"
)

func main() {
	intervals := parse.Parse("input_short.txt")
	reservoir := reservoir.InitFromIntervalls(intervals)
	fmt.Println(reservoir)
}
