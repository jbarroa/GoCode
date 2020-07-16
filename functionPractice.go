//Jasmine Barroa
//June 11, 2020
//Practice functions

package main

import (
   "fmt"
)

func checkUsername(entered string) bool {
   return entered == "jbarroa"
}

func checkPassword(entered string) bool{
	return entered == "henlo"
}

func main() {
   
	var username string
	var password string

	fmt.Print("Enter your username: ")
	fmt.Scan(&username)

	fmt.Print("Enter your password: ")
	fmt.Scan(&password)

	if(checkUsername(username) && checkPassword(password)){
		fmt.Println("Correct!")
	}else{
		fmt.Println("Time to go")
	}
}