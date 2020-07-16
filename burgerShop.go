//Jasmine Barroa
//June 15, 2020
//Online ordering system for a buger shop

package main

import "fmt"

var name string
var total float64

// create the struct
type burger struct {
   bun bool
   price float64
   dressed bool
}

type drink struct {
	price float64
	drinkType string
}

type side struct{
	price float64
	name string
}

type combo struct{
	burgerType burger
	drinkType drink
	sideType side
	price float64
}

//functions that uses structs
func user_input_burger(){
	var input string
	var input2 string
	bType := burger{}
	Message1:
		fmt.Print("Do you want a bun with your burger (Yes/No)? ")
		fmt.Scan(&input)
		fmt.Print("Do you want your burger fully dressed (Yes/No)? ")
		fmt.Scan(&input2)
	if input == "Yes" && input2 == "Yes"{
		bType.bun = true
		bType.price = 5.0
		bType.dressed = true
	}else if input == "Yes" && input2 == "No"{
		bType.bun = true
		bType.price = 5.0
		bType.dressed = false
	}else if input == "No" && input2 == "Yes"{
		bType.bun = false
		bType.price = 4.0
		bType.dressed = true
	}else if input == "No" && input2 == "No"{
		bType.bun = false
		bType.price = 4.0
		bType.dressed = false
	}else{
		fmt.Println(input, "is invalid")
		goto Message1
	}
	total += bType.price
	fmt.Println("The cost of your burger is:", bType.price)
}

func user_input_drink(){
	var input string
	dType := drink{}
	Message2:
		fmt.Print("Choose between Water, Coke, Sprite, Milkshake: ")
		fmt.Scan(&input)
	if input == "Water"{
		dType.price = 1.0
		dType.drinkType = input
	}else if(input == "Coke" || input == "Sprite" || input == "Milkshake"){
		dType.price = 2.0
		dType.drinkType = input
	}else{
		fmt.Println(input, "is invalid")
		goto Message2
	}
	total += dType.price
	fmt.Println("The cost of your drink is:", dType.price)
}

func user_input_side(){
	var input string
	sType := side{}
	Message3:
		fmt.Print("Choose between Fries, Onion Rings, Salad, Coleslaw: ")
		fmt.Scan(&input)
	if input == "Fries"{
		sType.price = 1.0
		sType.name = input
	}else if input == "Onion Rings" || input == "Salad" || input == "Coleslaw"{
		sType.price = 2.0
		sType.name = input	
	}else{
		fmt.Println(input, "is invalid")
		goto Message3
	}
	total += sType.price
	fmt.Println("The cost of your side is:", sType.price)
}

func user_input_combo(){
	var input string
	var input2 string
	bType := burger{}
	dType := drink{}
	sType := side{}
	Message1:
		fmt.Print("Do you want a bun with your burger (Yes/No)? ")
		fmt.Scan(&input)
		fmt.Print("Do you want your burger fully dressed (Yes/No)? ")
		fmt.Scan(&input2)
	if input == "Yes" && input2 == "Yes"{
		bType.bun = true
		bType.price = 5.0
		bType.dressed = true
	}else if input == "Yes" && input2 == "No"{
		bType.bun = true
		bType.price = 5.0
		bType.dressed = false
	}else if input == "No" && input2 == "Yes"{
		bType.bun = false
		bType.price = 4.0
		bType.dressed = true
	}else if input == "No" && input2 == "No"{
		bType.bun = false
		bType.price = 4.0
		bType.dressed = false
	}else{
		fmt.Println(input, "is invalid")
		goto Message1
	}
	Message2:
		fmt.Print("Choose between Water, Coke, Sprite, Milkshake: ")
		fmt.Scan(&input)
	if input == "Water"{
		dType.price = 1.0
		dType.drinkType = input
	}else if(input == "Coke" || input == "Sprite" || input == "Milkshake"){
		dType.price = 2.0
		dType.drinkType = input
	}else{
		fmt.Println(input, "is invalid")
		goto Message2
	}
	Message3:
		fmt.Print("Choose between Fries, Onion Rings, Salad, Coleslaw: ")
		fmt.Scan(&input)
	if input == "Fries"{
		sType.price = 1.0
		sType.name = input
	}else if input == "Onion Rings" || input == "Salad" || input == "Coleslaw"{
		sType.price = 2.0
		sType.name = input	
	}else{
		fmt.Println(input, "is invalid")
		goto Message3
	}
	discount := (bType.price + dType.price + sType.price) * 0.2
	chooseCombo := combo{
		bType,
		dType,
		sType,
		(bType.price + dType.price + sType.price) - discount,
	}
	total += chooseCombo.price
	fmt.Println("The cost of your Combo is:", chooseCombo.price)
}

func take_order_from_user() bool{
	fmt.Print("Hello. Please enter your name: ")
	fmt.Scan(&name)
	var order string
	
	for{
		fmt.Println("Enter 'Done' if you are finished or 'Cancel' if you would like to cancel your order")
		fmt.Print("Please choose between Burger, Drink, Side, or Combo: ")
		fmt.Scan(&order)
		if order == "Cancel"{
			fmt.Println("Your order has been cancelled")
			return true;
		}else if order == "Done"{
			break;
		}else if order == "Combo"{
			user_input_combo()
		}else if order == "Burger"{
			user_input_burger()
		}else if order == "Drink"{
			user_input_drink()
		}else if order == "Side"{
			user_input_side()
		}else{
			fmt.Print(order, "is invalid. Please try again.")
		}
	}
	fmt.Println("The total for your order is:", total)
	return false
}

func main(){
	order := take_order_from_user()
	if order == false{
		fmt.Println("Thank you for your order", name)
	}
}