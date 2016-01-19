package main

import "fmt"

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}

	var n = len(a)/2

	c := make(chan int)
	go sum(a[:n], c)
	go sum(a[n:], c)

	// 从 c 中获取
	// x, y := <-c, <-c 
	x := <-c
	y := <-c

	fmt.Println(x, y, x+y)
}
