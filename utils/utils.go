package utils

import (
    "fmt"
    "strings"
	"io/ioutil"
    "strconv"
)

func FilterArray(arr []string, f func(string) bool) []string {
    arr_c := make([]string, 0)
    for _, v := range arr {
        if f(v) {
            arr_c = append(arr_c, v)
        }
    }
    return arr_c
}

func check(e error) {
    if e != nil {
		fmt.Println("File reading error", e)
        panic(e)
    }
}

func Readfile(filePath string) string {
	data, err := ioutil.ReadFile(filePath)
    check(err)
	input := string(data)
    return input
}

func SplitArray(array string) []string {
    return strings.Split(array, "\n")
}

func SplitArrayDelim(array string, delim string) []string {
    return strings.Split(array, delim)
}

func RemoveSpaces(word string) []string {
    return strings.Fields(word)
}

func MapArrayToInt(array []string) []int64 {
    var array2 []int64
    for _, val := range array {
        intVal, _ := strconv.ParseInt(val, 10, 64)
        array2 = append(array2, intVal)
    }
    return array2
}

func Sum(array []int64) int64 {
    result := int64(0)
    for _, v := range array {
        result += v
    }
    return result
}