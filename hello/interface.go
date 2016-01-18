package main

import (
   "fmt"
   "math"
)

/* define an interface */
type Shape interface {
   area() float64
}

/* define a circle */
type Circle struct {
   x,y,radius float64
}

/* define a rectangle */
type Rectangle struct {
   width, height float64
}

/* define a method for circle (implementation of Shape.area())*/
func(c Circle) area() float64 {
   return math.Pi * c.radius * c.radius
}

/* define a method for rectangle (implementation of Shape.area())*/
func(r Rectangle) area() float64 {
   return r.width * r.height
}

/* define a method for shape */
func getArea(shape Shape) float64 {
   return shape.area()
}

func main() {
   circle := Circle{x:0,y:0,radius:2}
   rectangle := Rectangle {10, 5}
   
   fmt.Printf("Circle area: %f\n",getArea(circle))
   fmt.Printf("Rectangle area: %f\n",getArea(rectangle))
}