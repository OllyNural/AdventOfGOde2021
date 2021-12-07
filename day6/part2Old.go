package main

import (
	"fmt"
	"time"
	"github.com/ollynural/AdventOfGOde2021/utils"
)

func checkFish(val int64) bool {
	if val == 0 {
		return true
	}
	return false
}

func updateVal(curr *int64, val int64) {
	*curr = val
}

func main() {

	start := time.Now()

	input := utils.Readfile("day6/part1.txt")
	input_array := utils.MapArrayToInt64(utils.SplitArrayDelim(input, ","))
	// fmt.Println(input_array)
	c := 0
	limit := 79
	new_fish_array := []int64{}
	for {
		for i := 0; i < len(input_array); i++ {
			curr := input_array[i]
			// fmt.Println(i)
			// new_fish := checkFish(input_array, i)
			// fmt.Println(new_fish, i)
			if checkFish(curr) {
				new_fish_array = append(new_fish_array, int64(8))
				// updateVal(&curr, int64(6))
				curr = int64(6)
			} else {
				// updateVal(&curr, curr - 1)
				curr--
			}
		}

		if (c > limit) {
			break
		}
		c++
		// fmt.Println(input_array)
		// fmt.Println(new_fish_array)
		input_array = append(input_array, new_fish_array...)
		new_fish_array = nil
	}
	// fmt.Println(len(input_array))
	elapsed := time.Since(start)
	fmt.Printf("%d took %s", limit, elapsed)
}