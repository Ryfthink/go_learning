package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// 斐波纳契闭包
// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	sum1 := 0
	sum2 := 1
	return func() int {
		out := sum1 + sum2
		sum1 = sum2
		sum2 = out
		return out

	}
}


/*
Go 函数可以是一个闭包。闭包是一个函数值，它引用了函数体之外的变量。 这个函数可以对这个引用的变量进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。

例如，函数 adder 返回一个闭包。每个返回的闭包都被绑定到其各自的 sum 变量上。
*/
func main() {
	// 1. 参考 https://tour.go-zh.org/moretypes/22
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			i,
			pos(i),
			neg(-2*i),
		)
	}

	fmt.Println()

	// 2. 参考 https://tour.go-zh.org/moretypes/23
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t",f())
	}
	fmt.Println()
}
