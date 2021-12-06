package main

import (
    "fmt"
    "strings"
    "strconv"
    "github.com/ollynural/AdventOfGOde2021/utils"
)

type Location struct {
    aim int
    distance int
    depth int
}

func add_dir(current Location, direction string, value int) Location {
    if direction == "forward" {
        current.distance += value
        current.depth += (value * current.aim)
    } else if direction == "up" {
        current.aim -= value
    } else if direction == "down" {
        current.aim += value
    }
    return current
}

func main() {
    loc := Location{distance: 0, depth: 0}
    input := utils.Readfile("day2/part1.txt")
    input_array := utils.SplitArray(input)
    for i := 0; i < len(input_array); i++ {
        actions := strings.Split(input_array[i], " ")
        dir, val := actions[0], actions[1]
        distance, _ := strconv.Atoi(val)
        loc = add_dir(loc, dir, distance)
    }
    fmt.Println("Final loc")
    fmt.Println(loc)
    fmt.Println(loc.distance * loc.depth)
}