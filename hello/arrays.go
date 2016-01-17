package main

import "fmt"

func main() {
	// 1. normal
	var elements [10] int
	var i,j int
	for i = 0; i < len(elements); i++ {
		elements[i] = 100 + i
	}

	for j = 0; j < len(elements) ; j++ {
		fmt.Printf("elements[%d] is %d \n",j,elements[j])
	}

	// 2. multi-dimensional arrays
	/* an array with 5 rows and 2 columns*/
	var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}

    /* output each array element's value */
    for  i = 0; i < 5; i++ {
       for j = 0; j < 2; j++ {
          fmt.Printf("a[%d][%d] = %d ", i,j, a[i][j] )
       }
       fmt.Println()
    }

    // 3. passing arrays to function
    var arr = []int {2,3,4,6,7}
    avg := getAverage(arr,5)
	fmt.Printf("average is: %f \n",avg)
}

func getAverage(arr []int, size int) float32 {
   var i,sum int
   var avg float32  

   for i = 0; i < size;i++ {
      sum += arr[i]
   }

   avg = float32(sum / size)

   return avg;
}




