//Activity 1 
//Quiz Generator

package main 

import (
	"fmt"
	//"strconv"
) 

func main(){
	
	var ans1 string
	var ans2 string
	var ans3 string
	var ans4 string
	var ans5 string
	var correct int
	var wrong int
	var total int
	var inRow int

	fmt.Print("What is my favorite color? ")
	fmt.Scan(&ans1)

	if(ans1 == "Pink"){
		correct++
		total++
		fmt.Println("You are correct. You have answered ", correct, "answer(s).")
		fmt.Println("You have an accurately answered ", float(correct/total), "% of the quiz")
	} else{
		wrong++
		total++
		inRow++ 
		if(ans3 == "Jasmine"){
			correct++
			total++
			inRow = 0
			fmt.Println("You are correct. You have answered", correct, "answer(s).")
			
		} else if(inRow < 2){
			wrong++
			total++
			inRow++
			fmt.Println("You are incorrect")
			fmt.Println("You have accurately answered", float(correct/total), "% of the quiz")
		}else{
			inRow++
			fmt.Println("The quiz ended because you answered 3 questions in a row incorrectly")
			fmt.Println("You have an accurately answered", float(correct/total), "% of the quiz")
		}
	}

	fmt.Print("How old am I? ")
	fmt.Scan(&ans2)

	if(ans2 == "21"){
		correct++
		total++
		inRow = 0
		fmt.Println("You are correct. You have answered ", correct, "answer(s).")
		fmt.Println("You have an accurately answered ", float(correct/total), "% of the quiz")
	} else{
		wrong++
		total++
		inRow++ 
		fmt.Println("You are incorrect")
		fmt.Println("You have accurately answered ", float(correct/total), "% of the quiz")
	}

	fmt.Print("What is my name? ")
	fmt.Scan(&ans3)

	if(ans3 == "Jasmine"){
		correct++
		total++
		inRow = 0
		fmt.Println("You are correct. You have answered ", correct, "answer(s).")
		
	} else if(inRow < 2){
		wrong++
		total++
		inRow++
		fmt.Println("You are incorrect")
		fmt.Println("You have accurately answered ", float(correct/total), "% of the quiz")
	}else{
		inRow++
		fmt.Println("The quiz ended because you answered 3 questions in a row incorrectly")
		fmt.Println("You have an accurately answered ", float(correct/total), "% of the quiz")
	}

	fmt.Print("When is my birthday? ")
	fmt.Scan(&ans4)

	if(ans4 == "9/14"){
		correct++
		total++
		inRow = 0
		fmt.Println("You are correct. You have answered", correct, "answer(s).")
		
	} else if(inRow < 3){
		wrong++
		total++
		inRow++
		fmt.Println("You are incorrect")
		fmt.Println("You have accurately answered", float(correct/total), "% of the quiz")
	}else{
		inRow++
		fmt.Println("The quiz ended because you answered 3 questions in a row incorrectly")
		
	}

	fmt.Print("Where was I born? ")
	fmt.Scan(&ans5)

	if(ans5 == "NYC"){
		correct++
		total++
		inRow = 0
		fmt.Println("You are correct. You have answered ", correct, "answer(s).")
		fmt.Println("You have an accurately answered ", float(correct/total), "% of the quiz")
	} else if(inRow < 3){
		wrong++
		total++
		inRow++
		fmt.Println("You are incorrect")
		fmt.Println("You have accurately answered ", float(correct/total), "% of the quiz")
	}else{
		inRow++
		fmt.Println("The quiz ended because you answered 3 questions in a row incorrectly")
		fmt.Println("You have an accurately answered ", float(correct/total), "% of the quiz")
	}
}