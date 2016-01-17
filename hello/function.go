package main

import (
	"fmt"
	"math"
)

/* define a circle */
type Circle struct {
	x,y,radius float64
}

func (c Circle) area() float64 {
	return math.Pi*c.radius*c.radius
}

func main() {

	// 1. build-in function
	var hello string = "hello ,go!"
	fmt.Printf("value: %s\n", hello)
	length := len(hello)
	fmt.Printf("length: %d\n", length)

	numbers := [20]int{1,2,3,5,7}
	length = len(numbers)
	fmt.Printf("length: %d\n", length)

	// 2.custom function
	maxNum,msg := max(1,2)
	fmt.Printf("maxNum is: %d %s\n", maxNum, msg)	

	// 3.fuction called by value 
	var a int = 100
	var b int = 200
	swap_value(a, b)
	fmt.Printf("swap_value a is: %d \n", a)	
	fmt.Printf("swap_value b is: %d \n", b)	

	// 4.function called by reference
	a = 100
	b = 200
	swap_reference(&a, &b)
	fmt.Printf("swap_reference a is: %d \n", a)	
	fmt.Printf("swap_reference b is: %d \n", b)		

	// 5.function as value
	getSquare := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println("getSquare(9) is", getSquare(9))

	// 6.function closures
	nextNumber := getSequence()
	fmt.Println("nextNumber is", nextNumber())	
	fmt.Println("nextNumber is", nextNumber())	
	fmt.Println("nextNumber is", nextNumber())	
	fmt.Println("nextNumber is", nextNumber())	


	// 7. function method , see method.go
	circle := Circle{0,0,3}
	fmt.Printf("Circle area: %f \n",circle.area())

}

func max(num1, num2 int) (int, string) {
	if num1 < num2 {
		return num2 , "num2 max"
	} else {
		return num1 , "num1 max"
	}
}

func swap_value(x,y int)(int){
	var temp int
	temp = x
	y = x
	x = temp
	return temp
}

func swap_reference(x *int, y *int) {
   var temp int
   temp = *x    /* save the value at address x */
   *x = *y    /* put y into x */
   *y = temp    /* put temp into y */
}

// 返回一个闭包函数
func getSequence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}





