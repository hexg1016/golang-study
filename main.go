package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "1|1|1|1|1|"
	str2 := "1|1|1|1|1"
	arr1 := strings.Split(str1, "|")
	arr2 := strings.Split(str2, "|")
	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
}
