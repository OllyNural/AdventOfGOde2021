package main

import (
	"fmt"
    "strconv"
    "github.com/ollynural/AdventOfGOde2021/utils"
)

type power struct {
	gamma int
	epsilon int
}

func main() {
	input := utils.Readfile("day3/part1.txt")
    input_array := utils.SplitArray(input)
	fmt.Println(input_array)
	gamma_pow := ""
	// power := power{gamma: 0, epsilon: 0}
	for j := 0; j < len(input_array[0]); j++ {
		count := 0
		for i := 0; i < len(input_array); i++ {
			fmt.Println(input_array[i][j:j+1])
			if input_array[i][j:j+1] == "1" {
				count++
			}
		}
		// fmt.Println(count)
		if count > len(input_array) / 2 {
			gamma_pow += "1"
		} else {
			gamma_pow += "0"
		}
	}
	fmt.Println(gamma_pow)
	gamma, _ := strconv.ParseInt(gamma_pow, 2, 64)
	epsilon := 4095 - gamma
	fmt.Println(gamma)
	fmt.Println(epsilon)
	fmt.Println(gamma * epsilon)
}