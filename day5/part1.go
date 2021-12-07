package main

import (
	"fmt"
	"github.com/ollynural/AdventOfGOde2021/utils"
)

func isValid(start []int64, end []int64) bool {
	if start[0] == end[0] || start[1] == end[1] {
		return true
	}
	return false
}

func countOverlaps(array [1000][1000]uint8, limit uint8) int {
	count := 0
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if array[i][j] >= limit {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println("Hello world")
	input := utils.Readfile("day5/part1.txt")
	input_array := utils.SplitArray(input)

	board := [1000][1000]uint8{}
	fmt.Println(board)
	fmt.Println(input_array)
	for i := 0; i < len(input_array); i++ {
		line := utils.SplitArrayDelim(input_array[i], " -> ")
		start, end := 
			utils.MapArrayToInt64(utils.SplitArrayDelim(line[0], ",")), 
			utils.MapArrayToInt64(utils.SplitArrayDelim(line[1], ","))

		if !isValid(start, end) {
			continue
		}
		fmt.Println(start[0], start[1])
		fmt.Println(end[0], end[1])

		x_start := utils.Min(start[0], end[0])
		x_end := utils.Max(start[0], end[0])
		y_start := utils.Min(start[1], end[1])
		y_end := utils.Max(start[1], end[1])
		for m := y_start; m <= y_end; m++ {
			for n := x_start; n <= x_end; n++ {
				// fmt.Println("Inside")
				board[n][m] += 1
			}
		}
	}
	fmt.Println("Finished")
	// fmt.Println(board)
	result := countOverlaps(board, 2)
	fmt.Println(board[287])
	fmt.Println(result)
}