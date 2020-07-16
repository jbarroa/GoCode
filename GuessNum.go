//Exercise 3

package main

import (
	"fmt"
	"strconv"
	"math/rand"
)

func main() {

	var name string
	var level string
	var randomNum int
	var guessedS string
	guesses := 0

	fmt.Println("What is your name? ")
	fmt.Scan(&name)

	fmt.Println("What level would you like to play? Easy/Normal/Hard ")
	fmt.Scan(&level)

	if(level == "Easy"){
		randomNum = rand.Intn(5)
	} else if(level == "Normal"){
		randomNum = rand.Intn(20)
	} else{
		randomNum = rand.Intn(50)
	}

	for{
		fmt.Println("Guess a number: ")
		fmt.Scan(&guessedS)
		guessNum, _ := strconv.Atoi(guessedS)
		guesses++;
		if(guessNum < 1 || (level == "Easy" && guessNum > 5) || (level == "Normal" && guessNum > 20) || (level == "Hard" && guessNum > 50)){
			guesses--;
			fmt.Println(name, "you have to enter a valid number")
			continue;
		} else if(guesses == 1 && randomNum == guessNum){
			fmt.Println("Congratulations", name, "! You guessed it correctly on your first try!")
			break;
		} else if(randomNum == guessNum){
			fmt.Println(name, ", you guessed the number! It took you" ,guesses, "to guess the number!")
			break;
		}
	}
}