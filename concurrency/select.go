package main

import "fmt"
import "time"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func fillQuit(c,quit chan int) {
	for i := 0; i < 20; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Println(<-c)
	}
	quit <- 0
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go fillQuit(c, quit)
	fibonacci(c, quit)
}
