package main

import "fmt"

func main() {
	/*create a slice*/
	numbers := []int{0,1,2,3,4,5,6,7,8}

	/* print the numbers */
	for i,j := range numbers {
		fmt.Println("Slice item",i,"is",j)	
	}

	countryCapitalMap := map[string] string {"France":"Paris","Italy":"Rome","Japan":"Tokyo"}
	for k,v := range countryCapitalMap {
		fmt.Println("Country",k,"capital is",v)
	}

	
	
}