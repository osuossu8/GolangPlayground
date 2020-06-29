package main

import (
	"golang.org/x/tour/wc"
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	slice := strings.Split(s, " ")
	//fmt.Println(slice)
	//fmt.Println(slice[0])
	//fmt.Println(len(slice))
	var map_ex = map[string]int{}

	for i:=0; i<len(slice); i++{
		fmt.Println(slice[i])
		map_ex[slice[i]] += 1
	}
	return map_ex //map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
	//WordCount("I am a pen")
	//fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
}
