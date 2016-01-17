package main

import "fmt"

func main() {
	// 1.
	var i int = 20
	var point *int
	fmt.Printf("nil point value: %x \n",point)

	point = &i
	fmt.Printf("Address of i is: %x \n",point)
	fmt.Printf("Access the value of *point: %d \n\n",*point)

	// 2.
	arrayOfPointers()

	// 3.
	pointerToPointer()


	// 4. passing pointer to function
	var a,b int = 4,5
	fmt.Printf("Before,a is %d, b is %d \n",a,b)
	passPointerToFun(&a,&b)
	fmt.Printf("Swaped,a is %d, b is %d \n",a,b)
}

//  Array of pointers
func arrayOfPointers(){
	const MAX int = 3
    a := []int{10,100,200}

    var i int
    var ptr [MAX]*int;

    for  i = 0; i < MAX; i++ {
       ptr[i] = &a[i] /* assign the address of integer. */
    }

    for  i = 0; i < MAX; i++ {
       fmt.Printf("Value of a[%d] = %d\n", i,*ptr[i] )
    }
}

// Pointer to pointer
func pointerToPointer(){
	var a int
   	var ptr *int
   	var pptr **int
	
   	a = 3000
	
   	/* take the address of var */
   	ptr = &a
	
   	/* take the address of ptr using address of operator & */
   	pptr = &ptr
	
   	/* take the value using pptr */
   	fmt.Printf("Value of a = %d\n", a )
   	fmt.Printf("Value available at *ptr = %d\n", *ptr )
   	fmt.Printf("Value available at **pptr = %d\n", **pptr)
}	

// passing pointer to function
func passPointerToFun(x *int,y *int){
	var temp int
	temp = *x
	*x = *y
	*y = temp
}
