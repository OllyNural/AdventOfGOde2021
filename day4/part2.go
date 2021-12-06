package main

import (
	"fmt"
	"strconv"
	"github.com/ollynural/AdventOfGOde2021/utils"
)

type cell struct {
	value int64
	found int64
}

func mapToCell(array []string) []cell {
	var array_c []cell
	for _, val := range array {
		value, _ := strconv.ParseInt(val, 10, 64)
		cell := cell{value: value, found: 0}
		array_c = append(array_c, cell)
	}
	return array_c
}

func sumCells(array []cell) int64 {
    result := int64(0)
    for _, v := range array {
        result += v.found
    }
    return result
}

func checkBoard(board [][]cell) bool {
	isWin := false
	for i := 0; i < len(board); i++ {
		if sumCells(board[i]) == 5 {
			isWin = true
		}
	}

	for j := 0; j < len(board[0]); j++ {
		count := board[0][j].found + board[1][j].found + board[2][j].found + board[3][j].found + board[4][j].found
		if count == 5 {
			isWin = true
		}
	}
	// fmt.Println(board)
	return isWin
}

func checkAndUpdateBoard(board *[][]cell, num int64) bool {
	hasUpdated := false
	for i := 0; i < len(*board); i++ {
		for j := 0; j < len((*board)[i]); j++ {
			if (*board)[i][j].value == num {
				hasUpdated = true
				(*board)[i][j].found = 1
			}
		}
	}
	return hasUpdated
}

func removeIndex(array [][][]cell, index int) [][][]cell {
	return append(array[:index], array[index+1:]...)
}

func main() {
	// fmt.Println(input_array)
	input := utils.Readfile("day4/part1.txt")
	input_array := utils.SplitArray(input)
	numbers := utils.MapArrayToInt(utils.SplitArrayDelim(input_array[0], ","))
	var bingo_cards [][][]cell

	var temp_board [][]cell
	for i := 1; i < len(input_array); i++ {
		if input_array[i] == "" {
			bingo_cards = append(bingo_cards, temp_board)
			temp_board = [][]cell{}
			continue
		}
		curr_row := mapToCell(utils.RemoveSpaces(input_array[i]))
		temp_board = append(temp_board, curr_row)
	}

	for _, num := range numbers {
		for i := 0; i < len(bingo_cards); i++ {
			curr := bingo_cards[i]
			// Loop through and update any boards
			isUpdated := checkAndUpdateBoard(&curr, num)
			if isUpdated {
				// Check board for a bingo
				isBingo := checkBoard(curr)
				if isBingo {
					fmt.Println("Found Bingo!")
					fmt.Println("num:", num)
					fmt.Println(curr)
					unmarked := int64(0)
					for i := 0; i < len(curr); i++ {
						for j := 0; j < len(curr[i]); j++ {
							if curr[i][j].found == 0 {
								unmarked += curr[i][j].value
							}
						}
					}
					fmt.Println(unmarked)
					bingo_cards = removeIndex(bingo_cards, i)
					i--
				}
			}
		}
	}
}