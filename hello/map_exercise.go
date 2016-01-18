package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)

	arr := strings.Fields(s) // 相当于 java 中的 String.split(" ")

	for _, val := range arr {
		ret[val]++
	}
	return ret
}

// 参考 https://tour.go-zh.org/moretypes/20
// map 存储每个单词在一段长文中出现的次数
func main() {
	wc.Test(WordCount)
}