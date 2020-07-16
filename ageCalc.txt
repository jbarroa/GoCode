package main

import (
	"fmt"
	"time"
	"strconv"
)

func main(){
	date1 := enterFullDate()
	date2 := enterFullDate()

	diff := date1.Sub(date2)
	if date2.After(date1){
		diff = date2.Sub(date1)
	}

	months := int(diff.Hours())/730
	years := int(diff.Hours())/24/365
	
	remainMonths := int(diff.Hours())/730%12
	remainDays := (int(diff.Hours()) - (730 * months))/24
	fmt.Println(years, "years", remainMonths, "months", remainDays, "days")
	fmt.Println(months, "months", remainDays, "days")
	weeks := int(diff.Hours())/24/7
	daysLeft := int(diff.Hours())/24%7
	fmt.Println(weeks, "weeks", daysLeft, "days")
	numDays := int(diff.Hours()/24)
	fmt.Println(numDays, "days")
	fmt.Println(diff.Hours(), "hours")
	fmt.Println(diff.Minutes(), "minutes")
	fmt.Println(diff.Seconds(), "seconds")
}

func enterFullDate() time.Time{//use time to create a full date and time 
	var input string

	fmt.Println("Enter 'clear' to clear calculator")
	Message1:
		fmt.Print("Enter a year (in the format yyyy): ")
		fmt.Scan(&input)

	if input == "clear"{
		goto Message1
	}
	year, err := strconv.Atoi(input)
	if err != nil{
		goto Message1
	}
	Message2:
		fmt.Print("Enter a month (in the format mm): ")
		fmt.Scan(&input)

	if input == "clear"{
		goto Message1
	}
	month, err := strconv.Atoi(input)
	if err != nil{
		goto Message2
	}
	Message3:
		fmt.Print("Enter a day (in the format dd): ")
		fmt.Scan(&input)

	if input == "clear"{
		goto Message1
	}
	day, err := strconv.Atoi(input)
	if err != nil{
		goto Message3
	}

	Message4:
		fmt.Print("Enter in an hour (in the format hh): ")
		fmt.Scan(&input)

	if input == "clear"{
		goto Message1
	}
	hour, err := strconv.Atoi(input)
	if err != nil{
		goto Message4
	}

	Message5:
		fmt.Print("Enter in a minute (in the format mm): ")
		fmt.Scan(&input)

	if input == "clear"{
		goto Message1
	}
	min, err := strconv.Atoi(input)
	if err != nil{
		goto Message5
	}

	Message6:
		fmt.Print("Enter in a second (in the format ss): ")
		fmt.Scan(&input)

	if input == "clear"{
		goto Message1
	}
	sec, err := strconv.Atoi(input)
	if err != nil{
		goto Message6
	}

	date := time.Date(year, time.Month(month), day, hour, min, sec, int(0), time.Local)
	return date 
}