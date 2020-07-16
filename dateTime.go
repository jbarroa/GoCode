//Jasmine Barroa
//June 24, 2020
//Practice Date-Time Operation

package main

import (
	"fmt"
//	"math/rand"
	"time"
	"strconv"
)

func main(){

	var input string

	for input != "8"{
		fmt.Println("Menu:")
		fmt.Println("1. Time now")
		fmt.Println("2. Leap Year")
		fmt.Println("3. Subtract 5 days")
		fmt.Println("4. Print dates")
		fmt.Println("5. Next 5 days")
		fmt.Println("6. Add 5 seconds")
		fmt.Println("7. Difference")
		fmt.Println("8. Exit")
		fmt.Print("Choose 1-8: ")
		fmt.Scan(&input)
	//	choice, _ = strconv.Atoi(input)
	
		switch input {
		case "1":
			//Practice 1
			todayTime()
	
		case "2":
			//Practice 2
			leapYear()
		
		case "3":
			//Practice 3
			subtract5()
	
		case "4":
			//Practice 4
			printDates()
	
		case "5":
			//Practice 5
			next5()
	
		case "6":
			//Practice 6
			fivesec()
	
		case "7":
			//Practice 7
			difference()	
		
		default:
			break
		}

	}

}

func todayTime(){
	now := time.Now()
	fmt.Println("Today's date and time:", now)
	fmt.Println("Current year:", now.Year())
	fmt.Println("Current month:", now.Month())
	tn := time.Now().UTC()
    fmt.Println(tn)
    _, week := tn.ISOWeek()
	fmt.Println("Week number:", week)
	fmt.Println("Weekday:")
	fmt.Println("Day of the year:")
	fmt.Println("Day of the month:")
	fmt.Println("Day of the week:", now.Weekday())
}

func leapYear(){
	var input string
	fmt.Print("Enter a year: ")
	fmt.Scan(&input)
	choice, _ := strconv.Atoi(input)
	if choice % 4 == 0{
		if choice % 100 == 0 && choice % 400 != 0{
			fmt.Println("Not a leap year :(")
		}else{
			fmt.Println("LEAP YEAR!")
		}
	}else{
		fmt.Println("Not a leap year :(")
	}
}

func subtract5(){

	date := time.Now()
	var input string
	fmt.Print("Enter a year (yyyy): ")
	fmt.Scan(&input)
	choice, _ := strconv.Atoi(input)

	fmt.Print("Enter a month (mm): ")
	fmt.Scan(&input)
	choice2, _ := strconv.Atoi(input)

	fmt.Print("Enter a day (dd): ")
	fmt.Scan(&input)
	choice3, err := strconv.Atoi(input)

	if err != nil{
		fmt.Println(date)
	}else{
		date = time.Date(choice, time.Month(choice2), choice3, int(0), int(0),int(0), int(0),time.Local)
	}
	after := date.AddDate(0, 0, -5)
	fmt.Println("Five days before:", after)
}

func printDates(){

	date := time.Now()
	yesterday := date.AddDate(0, 0, -1)
	tomorrow := date.AddDate(0, 0, 1)
	fmt.Println("Today:", date)
	fmt.Println("Yesterday:", yesterday)
	fmt.Println("Tomorrow:", tomorrow)
}

func next5(){
	date := time.Now()
	fmt.Println("Today:", date)
	tomorrow := date.AddDate(0, 0, 1)
	fmt.Println("Day 1:", tomorrow)
	tomorrow = date.AddDate(0, 0, 2)
	fmt.Println("Day 2:", tomorrow)
	tomorrow = date.AddDate(0, 0, 3)
	fmt.Println("Day 3:", tomorrow)
	tomorrow = date.AddDate(0, 0, 4)
	fmt.Println("Day 4:", tomorrow)
	tomorrow = date.AddDate(0, 0, 5)
	fmt.Println("Day 5:", tomorrow)
}

func fivesec(){

}

func difference(){

}