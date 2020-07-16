/*package main

import "fmt"

func main() {

}*/
//Activity 1
/*package main

import "fmt"

func main() {
   
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i := 0; i < len(alphabet); i++{
		fmt.Println(string(alphabet[i]))
	}
}*/

//Activity 2
/*package main

import "fmt"

func main() {
	fmt.Println("Even numbers:")
	for i := 0; i <= 100; i += 2{
		fmt.Print(i, + " ")
	}

	fmt.Println("Odd numbers:")
	for j := 0; j <= 100; j++{
		if(j%2 != 0){
			fmt.Println(j, " ")
		}
	} 
}*/

//Activity 3
/*package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++{

	}
}*/

//Activity 4
/*package main

import "fmt"

func main() {
	var total int
	for i := 0; i <= 100; i++{
		total += i
	}
	fmt.Print("The total of all numbers from 0-100 is: ", total)
}*/

//Activity 5
/*package main

import "fmt"

func main() {
	for i := 100; i <= 1000; i++{
		if(i%50 == 0){
			fmt.Print(i, " ")
		}
	}
}*/

//Activity 6
/*package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++{

	}
}*/

//Activity 7
/*package main

import "fmt"

func main() {
	var x int
	var y int
	total := 1
	fmt.Print("Enter in a number: ")
	fmt.Scan(&x)

	fmt.Print("Enter in another number: ")
	fmt.Scan(&y)

	for i := 1; i <= y; i++{
		total *= x
	}

	fmt.Print("The total is: ", total)
}*/

//Activity 8
/*package main

import "fmt"

func main() {
	var n int
	var sum int
	var total int
	fmt.Print("Enter in a number: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++{
		for j:= 1; j <= i; j++{
			sum += j
		}
		total += sum
		sum = 0
	}

	fmt.Print("The total is: ", total)
}*/

//Activity 9
/*package main

import "fmt"

func main() {

}*/
//Activity 10
/*package main

import "fmt"

func main() {
	var rows int
	fmt.Print("Input number of rows: ")
	fmt.Scan(&rows)

	for i := 1; i<=rows; i++{
		for j:= 1; j <= i; j++{
			fmt.Print(i)
		}
		fmt.Println()
	}
}*/

//Activity 11

//Activity 12

//Activity 13

//Activity 14
/*package main

import "fmt"

func main() {
	var input string
	//var output string
	
	fmt.Print("Enter a number: ")
	fmt.Scan(&input)
	in := []rune(input)

	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1{
		in[i], in[j] = in[j], in[i]
	}

	fmt.Print(string(in))
}*/

//Activity 15
/*package main

import "fmt"

func main() {
	var message string
	fmt.Print("Enter a word: ")
	fmt.Scan(&message)

	for i, c := range message{
		fmt.Println(string(c)) //value
		i++
	}
}*/

//Activity 16

//Activity 17
/*package main

import "fmt"

func main() {
	var input string
	//var output string
	
	fmt.Print("Enter a word: ")
	fmt.Scan(&input)
	in := []rune(input)

	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1{
		in[i], in[j] = in[j], in[i]
	}

	fmt.Print(string(in))
}*/

