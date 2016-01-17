package main

import "fmt"

func main() {
	// 1. simple output
	fmt.Printf("Hello,go\n")

	// 2. 常量
	const LENTH int = 30
	const HEIGHT int = 40
	fmt.Printf("Size:%d\n",30*40)
	
	// 3. 隐式声明变量类型
	var x float64 = 20.0
	y := 42 
	z := x
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
	fmt.Printf("z is of type %T\n", z)

	// 4. 多类型声明
	var a1, b2, c3 = 3, 4, "foo"
	fmt.Printf("a1 ,b2 ,c3 is of type %T, %T, %T\n", a1, b2, c3)

	// 5. 算数运算
	var a int = 21
   	var b int = 10
   	var c int

  	c = a + b
  	fmt.Printf("Line 1 - Value of c is %d\n", c )
  	c = a - b
  	fmt.Printf("Line 2 - Value of c is %d\n", c )
  	c = a * b
  	fmt.Printf("Line 3 - Value of c is %d\n", c )
  	c = a / b
  	fmt.Printf("Line 4 - Value of c is %d\n", c )
  	c = a % b
  	fmt.Printf("Line 5 - Value of c is %d\n", c )
  	a++
  	fmt.Printf("Line 6 - Value of a is %d\n", a )
  	a--
  	fmt.Printf("Line 7 - Value of a is %d\n", a )
	
  	// 6. * &
  	var d int = 30
  	var ptr *int
  	ptr = &a

  	fmt.Printf("d is %d,*ptr is %d.\n", d, *ptr);

  	// 7. select statement

}
