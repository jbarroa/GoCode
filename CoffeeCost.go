//Jasmine Barroa
//June 4, 2020
//Calculate the cost of coffee

package main

import (
	"fmt"
)
func main() {

	var size string
	var coffee string
	var flavoring string
	var flavorType string
	var price float32 

	Message1:
		fmt.Print("Do you want small, medium, or large? ")//asks user the size of coffee
		fmt.Scan(&size)

	//Determines the price of the size
	if(size == "small"){//user chooses a small size
		price += 2.5
	}else if(size == "medium"){//user chooses a medium
		price += 3.0
	}else if(size == "large"){//user chooses a large
		price += 4.0
	}else{
		fmt.Println("Please enter a valid size.")
		goto Message1//if not a valid size, ask the question again
	}

	Message2:
		fmt.Print("Do you want brewed, espresso, or coldBrew? ")//asks user type of coffee
		fmt.Scan(&coffee)

	//Determines the price pf coffee type
	if(coffee == "brewed"){
		price += 0
	}else if(coffee == "espresso"){
		price += 0.5
	}else if(coffee == "coldBrew"){
		price += 1.0
	}else{
		fmt.Println("Please enter the valid type of coffee.")
		goto Message2//if not a valid coffee type, ask the question again
	}

	Message3:
		fmt.Print("Do you want a flavored syrup? (yes/no) ")//asks the user if they want a flavored syrup
		fmt.Scan(&flavoring)

	if(flavoring != "yes" && flavoring != "no"){
		fmt.Println("Please enter if you want flavoring")
		goto Message3//if not a valid answer, ask the question again
	}else if(flavoring == "yes"){
		Message4:
			fmt.Print("Do you want hazelnut, vanilla, or caramel? ")
			fmt.Scan(&flavorType)
		if(flavorType != "hazelnut" && flavorType != "vanilla" && flavorType != "caramel"){
			fmt.Print("Please enter a valid flavor type")
			goto Message4//if not a valid flavoring syrup, ask the question again
		}else{
			price += 0.5//adds the price of the flavoring syrup
		}
	}else{
		flavorType = "no"//set flavor type to no to properly print
	}

	fmt.Println("You asked for a ", size," cup of ", coffee, " with ", flavorType, " syrup.")
	fmt.Print("Your cup of coffee costs ")
	fmt.Printf("%0.2f \n", price)//print out price with two decimal points ex: 0.00
	tip := price * 0.15//calculate the tip
	price += tip//add together total price with tip
	fmt.Print("The price with a tip is ")
	fmt.Printf("%0.2f \n", price)
}