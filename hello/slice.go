package main

import (
	"fmt"
	"strings"
)
// 主要介绍的方法:len() cap() append() copy() make()
func main() {
	// 1. normal
	fmt.Println("\nnormal")
  	var numbers = make([]int,3,5)
  	printSlice(numbers)

  	// 2. slice nil
  	fmt.Println("\nslice nil")
	sliceNil()

	// 3. sub slice
	fmt.Println("\nsub slice")
  	sliceSub()

  	// 4. append and copy slice
  	fmt.Println("\nappend and copy slice")
  	sliceAppendCopy()

	fmt.Println("\n多维")
  	// 5.cttt
  	cttt()
}

// sub slice
func sliceSub(){
	numbers := []int{0,1,2,3,4,5,6,7,8}
	fmt.Println(numbers);
	fmt.Printf("sub 1:4 %v\n",numbers[1:4]);
	fmt.Println("sub :4 ",numbers[:4]);
	fmt.Println("sub 5: ",numbers[5:]);

	numbers1 := make([]int,3,5)
	printSlice(numbers1)

	numbers2 := numbers[:2]
	printSlice(numbers2)

	numbers3 := numbers[2:5]
	printSlice(numbers3)
}

// nil slice
func sliceNil(){
	var numbers []int
	printSlice(numbers)

	if(numbers == nil){
    	fmt.Printf("slice is nil\n")
   	}
}

// append and copy slice
func sliceAppendCopy(){
	var numbers []int
	printSlice(numbers)

	/* append allows nil slice */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* add one element to slice*/
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* add more than one element at a time*/
	numbers = append(numbers, 2,3,4)
	printSlice(numbers)

	numbers1 := make([]int,len(numbers),(cap(numbers))*2)

	copy(numbers1,numbers)
	printSlice(numbers1)

}

// len() and cap() functions
func printSlice(x []int){
   fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

// Create a tic-tac-toe board.
func cttt(){
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	game[0][0] = "X"
	game[2][2] = "O"
	game[2][0] = "X"
	game[1][0] = "O"
	game[0][2] = "X"

	printBoard(game)
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}
