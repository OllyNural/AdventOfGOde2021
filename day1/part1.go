package main

import (
    "fmt"
    "strings"
    "strconv"
    "github.com/ollynural/AdventOfGOde2021/utils"
)

func main() {
    input := utils.Readfile("day1/part1.txt")
    values := strings.Split(input, "\n")
    int_input := utils.MapArrayToInt(values)
    count := 0
    for i := 1; i < len(int_input); i++ {
        curr, _ := strconv.ParseInt(int_input[i], 10, 64)
        prev, _ := strconv.ParseInt(int_input[i-1], 10, 64)
        if curr > prev {
            count++
        }
    }
    fmt.Println(count)
}
