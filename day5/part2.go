package main

import (
	"fmt"
	"github.com/ollynural/AdventOfGOde2021/utils"
	"github.com/adam-lavrik/go-imath/i64"
)

func isHorizontal(start []int64, end []int64) bool {
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
	input := utils.Readfile("day5/part1.txt")
	input_array := utils.SplitArray(input)

	board := [1000][1000]uint8{}
	for i := 0; i < len(input_array); i++ {
		line := utils.SplitArrayDelim(input_array[i], " -> ")
		start, end := 
			utils.MapArrayToInt(utils.SplitArrayDelim(line[0], ",")), 
			utils.MapArrayToInt(utils.SplitArrayDelim(line[1], ","))

		if isHorizontal(start, end) {
			// Horizontal
			x_start := utils.Min(start[0], end[0])
			x_end := utils.Max(start[0], end[0])
			y_start := utils.Min(start[1], end[1])
			y_end := utils.Max(start[1], end[1])
			for m := y_start; m <= y_end; m++ {
				for n := x_start; n <= x_end; n++ {
					board[m][n] += 1
				}
			}
		} else {
			// Diagonal
			x_diff := start[0] - end[0]
			y_diff := start[1] - end[1]

			diff := i64.Abs(x_diff)   

			x_move := x_diff / diff
			y_move := y_diff / diff

			for p := int64(0); p <= diff; p++ {
				fmt.Println(start, end)
				board[start[1] - (y_move*p)][start[0] - (x_move*p)] += 1
			}
		}
	}
	result := countOverlaps(board, 2)
	fmt.Println(result)
}