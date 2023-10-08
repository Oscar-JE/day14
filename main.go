package main

import (
	"day14/parse"
	"fmt"
)

func main() {
	intervals := parse.Parse("input_short.txt")
	fmt.Println(intervals)
}
