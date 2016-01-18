package main

import "fmt"

func main() {
	value := factorial(5)
	fmt.Println("5的阶乘=",value)

	fmt.Printf("10项目斐波那契数列:")
    for i := 0; i < 10; i++ {
    	value = fibonacci(i);
       	fmt.Printf("%d\t", value)
    }
}

// factorial 阶乘
func factorial(i int) int{
	if i <= 1{
		return 1
	}
	return i*factorial(i-1)
}

// 斐波那契数列（Fibonacci sequence)
func fibonacci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}
