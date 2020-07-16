//Jasmine Barroa
//June 16, 2020
//Keyword search using maps

package main

import (
	"fmt"
//	"strings"
//	"math"
)
func main(){

	var input string
	var keys string
	var values string
	var run string

    capitals := map[string]string{
		"Paris" : "France",
		"Stockholm" : "Sweden",
		"Manila" : "Philippines",
		"Seoul" : "South Korea",
		"Washington DC" : "USA",
		"London" : "UK",
		"Tokyo" : "Japan",
		"Rio" : "Brazil",
		"Copenhagen" : "Denmark",
		"Riga" : "Latvia",
	/*	"France" : "Europe",
        "Sweden" : "Europe",
        "Philippines" : "Asia",
        "South Korea" : "Asia",
        "USA" : "North America",*/
	}

	for{
		fmt.Print("Enter a country or a capital: ")
		fmt.Scan(&input)

		//iterate through a map
		for key, value := range capitals{
			if input == key{
				fmt.Println("The country is:", value)
				keys = input
				values = values	
				break
			}else if input == value{
				fmt.Println("The capital is:", key)
				keys = key
				values = input
				break
			}
		}

		if keys == "" && values == ""{
			fmt.Println(input, "is not a capital or country in the map")
		}

		fmt.Print("Enter done or again: ")
		fmt.Scan(&run)

		if run == "done"{
			break
		}
	}
}