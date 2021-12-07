// Taken inspiration from: https://www.reddit.com/r/adventofcode/comments/r9z49j/2021_day_6_solutions/

package main

import (
	"fmt"
	"math/big"
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
	input := utils.Readfile("day6/part1.txt")
	input_array := utils.MapArrayToInt(utils.SplitArrayDelim(input, ","))
	limit := 79
	final := make(map[int]*big.Int)
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range nums {
		final[v] = big.NewInt(int64(0))
	}

	for i := 0; i < len(input_array); i++ {
		final[input_array[i]] = final[input_array[i]].Add(final[input_array[i]], big.NewInt(1))
	}
	
	fmt.Println(final)

	for j := 0; j < limit; j++ {
		curr := final[0]
		fmt.Println(curr)
		for m := 1; m <= 9; m++ {
			final[m-1] = final[m]
			final[8] = curr
			final[6] = final[6].Add(final[6], curr)
		}
	}

	fmt.Println(final)
	sum := big.NewInt(int64(0))
	for _, v := range final {
		sum = sum.Add(sum, v)
	}
	fmt.Println(sum)
	// for i := 0; i < len(input_array); i++ {
	// 	fmt.Println(input_array[i])
	// 	totalLen += calculateFish(input_array[i])
	// 	fmt.Println("Found end of first one", totalLen)
	// }
	// fmt.Println("totalLen")
	// fmt.Println(input_array)
	// fmt.Println(final)
}

func calculateFish(val int64) int {
	input_array := make([]int64, val)

	// start := time.Now()

	// fmt.Println(input_array)

	c := 0
	limit := 255
	new_fish_array := []int64{}
	for {
		for i := 0; i < len(input_array); i++ {
			curr := &input_array[i]
			// fmt.Println("curr", curr)
			// new_fish := checkFish(input_array, i)
			// fmt.Println(new_fish, i)
			if checkFish(*curr) {
				new_fish_array = append(new_fish_array, int64(8))
				updateVal(curr, int64(6))
				// &curr = int64(6)
			} else {
				// fmt.Println(curr)
				updateVal(curr, *curr - 1)
				// fmt.Println(curr)
				// &curr = curr - int64(1)
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
	// elapsed := time.Since(start)
	// fmt.Printf("%d took %s", limit, elapsed)
	return len(input_array)
}