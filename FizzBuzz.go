//Jasmine Barroa
//June 4, 2020
//Print out fizz buzz as many times as the inputted number 

package main

import (
	"fmt"
	"strconv"
)

func main() {

	var input string//reads in the user's input
	var count int//counts how many times, fizz, buzz, and fizz buzz have been printed
	var i int //counter variable for the loop to print and to print out numbers
	fmt.Print("How many fizzing and buzzing units do you need in your life? ")
	fmt.Scan(&input)
	units, _ := strconv.Atoi(input)//convert string to int

	for{
		if(i == 0){//if it is on the first iteration print our number bc 
			fmt.Println(i)
		}else if(i%3 == 0 && i%5 == 0){//if the number is divisible by 3 and 5
			fmt.Println("fizz buzz")
			count++
		}else if(i%3 == 0){//if the number is only divisible by 3
			fmt.Println("fizz")
			count++
		}else if(i%5 == 0){//if the number is only divisible by
			fmt.Println("buzz")
			count++
		}else{//if the number is not divisble by 3 or 5
			fmt.Println(i)
		}
		//ends the loop once fizz buzz has been printed out the number of times that the user entered
		if(count == units){
			fmt.Println("TRADITION!!")
			break;//ends loop
		}
		i++;//increments the loop
	}
}