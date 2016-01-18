package main

import "fmt"

// declare a variable, by default map will be nil*/
// var map_variable map[key_data_type]value_data_type

// define the map as nil map can not be assigned any value*/
// map_variable = make(map[key_data_type]value_data_type)

func main() {

	// 1. map create
	countryCapitalMap := mapCreate()
	fmt.Println()	

	// 2. map retrieve
	mapRetrieve(countryCapitalMap)
	fmt.Println()

	// 3. map delete elements
	mapDelete(countryCapitalMap)
	fmt.Println()
}

func mapCreate() map[string] string {
	/* create a map 1*/
	var countryCapitalMap map[string]string
	
	/* create a map 2*/
	// countryCapitalMap := make(map[string]string)

	/* insert key-value pairs in the map*/
    countryCapitalMap = make(map[string]string)
    countryCapitalMap["France"] = "Paris"
    countryCapitalMap["Italy"] = "Rome"
    countryCapitalMap["Japan"] = "Tokyo"
    countryCapitalMap["India"] = "New Delhi"

    return countryCapitalMap;
}

func mapRetrieve(countryCapitalMap map[string]string){
	/* print map using keys*/
    for country := range countryCapitalMap {
    	fmt.Println("Capital of",country,"is",countryCapitalMap[country])
    }

    /* test contains the special key */
    capital,ok := countryCapitalMap["United States"]

    if ok {
    	fmt.Println("Has",capital)
    } else {
    	fmt.Println("Us not present")
    }
}

func mapDelete(countryCapitalMap map[string]string){
	fmt.Println("Delete key France")
	delete(countryCapitalMap,"France")
	mapRetrieve(countryCapitalMap)
}

