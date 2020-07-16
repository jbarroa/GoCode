//Rock Paper Scissors Group Work
package main

import (
	"fmt"
	"strconv"
	"math/rand"
)

func main() {
	var input string
	var user string //choice of user Rock Paper Scissors
	//var numRounds int //number of rounds must be less than 10
	var userWins int
	var computerWins int
	var tie int 
	var computer string

	for{
		fmt.Print("How many rounds would you like to play? (1-10) ")
		fmt.Scan(&input)
		numRounds, _ := strconv.Atoi(input)

		if(numRounds < 1 || numRounds > 10){//check if the number of rounds entered is valid
			fmt.Print("You must enter the number of rounds between 1-10")
			break;//exits 
		}
		for i := 0; i < numRounds; i++{//plays for number of rounds entered in by the useer
			fmt.Print("Enter a choice between Rock, Paper, or Scissors: ")
			fmt.Scan(&user)//user's choice

			if(user != "Rock" && user != "Paper" && user != "Scissors"){//checks for vaild input
				fmt.Print("A valid choice must be entered of either Rock, Paper, or Scissors.")
				i--;//decrements so that the round is not included 
				continue;//back to the beginning of for loop
			}
			compChoice := rand.Intn(3)//3 random integers as rock paper or scissors
			if(compChoice == 1){//declares computer choice as rock
				computer = "Rock"
			}else if(compChoice == 2){//declares computer choice as paper
				computer = "Paper"
			}else{//declares computer choice as scissors
				computer = "Scissors"
			}

			if(computer == user){//both picked same thing
				fmt.Println("It's a tie!")
				tie++;
			}else if(computer == "Rock" && user == "Scissors"){//rock beats scissors
				fmt.Println("The computer won!")
				computerWins++;
			}else if(computer == "Scissors" && user == "Paper"){//scissors beats paper
				fmt.Println("The computer won!")
				computerWins++;
			}else if(computer == "Paper" && user == "Rock"){//paper beats rock
				fmt.Println("The computer won!")
				computerWins++;
			}else if(user == "Rock" && computer == "Scissors"){
				fmt.Println("The user won!")
				userWins++;
			}else if(user == "Scissors" && computer == "Paper"){
				fmt.Println("The user won!")
				userWins++;
			}else{
				fmt.Println("The user won!")
				userWins++;
			}
		}
		if(userWins == computerWins){//Total of wins is equal for both user and computer
			fmt.Println("Number of ties: ", tie)
			fmt.Println("Number of user wins: ", userWins)
			fmt.Println("Number of computer wins: ", computerWins)
			fmt.Println("It's a tie")
		}else if(userWins > computerWins){//user wins is greater than computer
			fmt.Println("Number of ties: ", tie)
			fmt.Println("Number of user wins: ", userWins)
			fmt.Println("Number of computer wins: ", computerWins)
			fmt.Println("You won!")
		}else{//computer wins is greater than computer
			fmt.Println("Number of ties: ", tie)
			fmt.Println("Number of user wins: ", userWins)
			fmt.Println("Number of computer wins: ", computerWins)
			fmt.Println("The computer won! :(")
		}
		fmt.Print("Would you like to play another round? (yes/no) ")
		fmt.Scan(&input)
		if(input != "yes" && input != "no"){//checks for valid input
			fmt.Print("You must enter a valid input. This game is going to exit")
			break;//exits program
		}else if(input == "no"){//check for no
			fmt.Print("Thanks for playing!")
			break;//exits program
		}//program continues if the user selects yes
	}
	
}