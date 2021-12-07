package main

import (
	"fmt"
	"github.com/ollynural/AdventOfGOde2021/utils"
)

func checkFish(array []int64, i int) bool {
	if array[i] == 0 {
		return true
	}
	return false
}

func main() {
	input := utils.Readfile("day6/part1.txt")
	input_array := utils.MapArrayToInt64(utils.SplitArrayDelim(input, ","))
	// fmt.Println(input_array)
	c := 0
	limit := 79
	new_fish_array := []int64{}
	for {
		for i := 0; i < len(input_array); i++ {
			// fmt.Println(i)
			// new_fish := checkFish(input_array, i)
			// fmt.Println(new_fish, i)
			if checkFish(input_array, i) {
				new_fish_array = append(new_fish_array, int64(8))
				input_array[i] = int64(6)
			} else {
				input_array[i]--
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
	fmt.Println(len(input_array))
}