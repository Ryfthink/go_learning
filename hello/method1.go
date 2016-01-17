/*
method 语法 
func (r ReceiverType) funcName(parameters) (results)
*/

package main

import (
   "fmt"
   "math"
)

/* define a rectangle */
type Rectangle struct {
	width,height float64
}

/* define a circle */
type Circle struct {
   radius float64
}

/* define a method for rectangle */
func (r Rectangle) area() float64 {
	return r.width*r.height
}

/* define a method for circle */
func (c Circle) area() float64 {
   return math.Pi * c.radius * c.radius
}

func main(){
	// circle := Circle(radius:5)
	r1 := Rectangle{3, 4}
	r2 := Rectangle{4, 4}
	c1 := Circle{3}
	c2 := Circle{5}
	fmt.Printf("r1 area: %f \n", r1.area())
	fmt.Printf("r2 area: %f \n", r2.area())
	fmt.Printf("c1 area: %f \n", c1.area())
	fmt.Printf("c2 area: %f \n", c2.area())
}