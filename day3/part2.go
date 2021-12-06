package main

import (
	"fmt"
    "strconv"
    "github.com/ollynural/AdventOfGOde2021/utils"
)

type filter_res struct {
	filtered []string
	value string
}

func rating_finder(array []string, curr_index int, most_common bool) filter_res {
	filtered_one := utils.FilterArray(array, func(val string) bool {
		return val[curr_index:curr_index+1] == "1"
	})
	filtered_zero := utils.FilterArray(array, func(val string) bool {
		return val[curr_index:curr_index+1] == "0"
	})
	largest_array := len(filtered_one) >= len(filtered_zero)
	filter_one_res := filter_res{ filtered: filtered_one, value: "1" }
	filter_zero_res := filter_res{ filtered: filtered_zero, value: "0" }
	if (most_common) {
		if (largest_array) {
			return filter_one_res
		} else {
			return filter_zero_res
		}
	} else {
		if (largest_array) {
			return filter_zero_res
		} else {
			return filter_one_res
		}
	}
}

func main() {
	input := utils.Readfile("day3/part1.txt")
    input_array := utils.SplitArray(input)

	oxygen_input := input_array
	curr_str := ""
	i := 0
	for {
		oxygen_result := rating_finder(oxygen_input, i, true)
		curr_str += oxygen_result.value
		oxygen_input = oxygen_result.filtered
		if len(oxygen_input) == 1 {
			break
		}
		i++
	}
	fmt.Println(curr_str)
	fmt.Println(oxygen_input)
	fmt.Println(strconv.ParseInt(curr_str, 2, 64))

	co2_input := input_array
	co2_str := ""
	j := 0
	for {
		co2_result := rating_finder(co2_input, j, false)
		co2_str += co2_result.value
		co2_input = co2_result.filtered
		if len(co2_input) == 1 {
			break
		}
		j++
	}
	fmt.Println(co2_str)
	fmt.Println(co2_input)
	// Got to 01101010 before finding unique answer
	fmt.Println(strconv.ParseInt("011010100101", 2, 64))
}