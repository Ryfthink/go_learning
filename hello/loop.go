package main

import "fmt"

func main() {
	var a int
	var b int = 15

	numbers := [6]int{1,2,3,5}

	// 1.
	fmt.Println("\n1.")
	for a := 0; a < 10; a++ {
		fmt.Printf("values of a: %d\n", a)
	}

	// 2.
	fmt.Println("\n2.")
	for a < b {
		a++
		if a < 5 {
			continue
		} else {
			fmt.Printf("values of a: %d\n", a)
		}
		
		if a == 10 {
			break;
		}
	}

	// 3.
	fmt.Println("\n3.")
	for index,value := range numbers {
		fmt.Printf("value = %d at index %d\n", value,index)
	}

	// 4.
	fmt.Println("\n4.")
	a = 10;
	LOOP: for a < 20 {
		if a == 15 {
			a = a + 1
			goto LOOP
		}
		fmt.Printf("value of a: %d \n", a)
		a++
	}




}