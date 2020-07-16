//Jasmine Barroa
//June 30, 2020
//Dollar program to practice test driven development

package main

import "fmt"

// function formats the output
func FormatAmount(a float64) string {
	// use %.2f for precision 2 which is enough to represent dollar amounts for now
	return "USD " + fmt.Sprintf("%.2f", a)
 }

 func SubtractFormatAmount(a, b float64) string {
	return "USD " + fmt.Sprintf("%.2f", a - b)
 }

func main(){
   fmt.Println("main function")
   fmt.Println(FormatAmount(2.00))
   fmt.Println(SubtractFormatAmount(2.00, 1.14))
}