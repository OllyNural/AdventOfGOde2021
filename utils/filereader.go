package utils

import (
    "fmt"
	"io/ioutil"
    "strconv"
)

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
    // fmt.Println("Contents of file:", input)
    return input
}

func MapArrayToInt(array []string) []int64 {
    var array2 []int64
    for _, val := range array {
        intVal, _ := strconv.ParseInt(val, 10, 64)
        array2 = append(array2, intVal)
    }
    return array2
}