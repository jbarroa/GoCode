//Activity 1A
// output the text in quotation marks
/*package main

import "fmt"

func main() {
   fmt.Println("Hello, world!")
}*/

//Activity 1B
// display the text in quotation marks to an output block
// without moving any of the existing code to a different line
/*package main

import "fmt"

func main() {
  fmt.Println("Go is fun!"); fmt.Println("Go is also easy.")
}*/

//Activity 1C
// create separate variables for first name and last name
// print each name on a separate line
/*package main

import "fmt"

func main() {
  var firstName string = "Rebecca" // first name
  var lastName string = "Roberts" // last name

  fmt.Println(firstName)
  fmt.Println(lastName)
}*/

//Activity 2
// Change this program so that it outputs only the text Go is fun!
// Do not delete any of the existing code
/*package main

import "fmt"

func main() {
  fmt.Println("Go is fun!")
  //fmt.Println("Go is also easy!")
}*/

//Activity 3
// Add one line of code that displays the text in quotation marks to 
// an output block without repeating the text in quotation marks.
/*package main

import "fmt"

func main() {
  var output string = "I love Go!"
  fmt.Println(output)
}*/

//Activity 4
/*Create a script that prompts the user for the name of the state 
where they were born and the name of the state where they live now. 
Save each value in its own variable and display the input values to 
the user.*/

package main

import "fmt"

func main() {
	var born string
	var live string

	fmt.Print("What state were you born in? ")
	fmt.Scan(&born)

	fmt.Print("What state do you live in? ")
	fmt.Scan(&live)

	fmt.Println("You were born in", born)
	fmt.Println("You live in", live)
}