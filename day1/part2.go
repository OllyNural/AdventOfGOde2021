package main

import (
    "fmt"
    "strings"
    "github.com/ollynural/AdventOfGOde2021/utils"
)

func main() {
    input := utils.Readfile("day1/part1.txt")
    values := strings.Split(input, "\n")
    int_input := utils.MapArrayToInt64(values)
    fmt.Println(int_input)
    count := 0
    for i := 1; i < len(int_input) - 3; i++ {
        curr := int_input[i] + int_input[i+1] + int_input[i+2]
        next := int_input[i+1] + int_input[i+2] + int_input[i+3]
        if next > curr {
            count++
        }
    }
    fmt.Println(count)
}
