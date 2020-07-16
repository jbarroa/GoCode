//Activity 1
// The program should output 0
/*package main

import "fmt"

func main() {
  // do not change the order in which the numbers and operators appear
  fmt.Println(((5 + 3) % 2) * 9)
}*/

//Activity 2
//Create a computer program that prompts the user for a float number 
//and returns the integer portion of the floating number.

/*package main

import (
	"fmt"
	"strconv"
)

func main(){
	var doubleS string
	var double float64

	fmt.Print("Enter a float: ")
	fmt.Scan(&doubleS)

	double, _ = strconv.ParseFloat(doubleS, 0)
//	fmt.Println("The float you entered is", double)

	fmt.Println("The integer portion is", int(double))
}*/

//Activity 3
//Write a computer program that calculates and displays the current 
//value of a deposit for a given initial deposit, interest rate, how 
//many times interest is calculated per year, and the number of years 
//since the initial deposit.