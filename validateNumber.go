//Activity 3
//Validate US Phone Number

package main 

import (
	"fmt"
	//"strconv"
) 

func main(){
	
	var phone string
	var valNumber string

	Message1:
		fmt.Print("Enter in your phone number: ")
		fmt.Scan(&phone)

	if(len(phone) > 10 || len(phone) < 10){
		fmt.Println("Enter a valid phone number")
		goto Message1
	} else{
		num1 := phone[0:3]
		num2 := phone[3:6]
		num3 := phone[6:10]
		
		valNumber = num1 + "-" + num2 + "-" + num3
	}
	
	fmt.Print("Your normalized number is ", valNumber)
}