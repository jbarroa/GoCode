//Jasmine Barroa
//June 25, 2020
//Date-Time Calculator 

package main

import (
	"fmt"
	"time"
	"strconv"
)

//create new time struct for time duration calculator
type CalcTime struct{
	days int 
	hours int 
	minutes int
	seconds int
}

//assign days, hours, minutes, seconds to struct
func (t *CalcTime) createTime(d int, h int, m int, s int){
	t.days = d 
	t.hours = h 
	t.minutes = m 
	t.seconds = s
}

func main(){

	var input string
	//Loop through menu until the user is done 
	//For ease of using program
	for input != "4"{
		fmt.Println("Menu:")
		fmt.Println("1. Time Duration")
		fmt.Println("2. Add/Subtract Time from Date")
		fmt.Println("3. Age Calculator")
        fmt.Println("4. Exit")
        fmt.Print("Choose a number 1-4: ")
        fmt.Scan(&input)

        switch input {
		case "1":
			timeDuration()
	
		case "2":
			addSubtractTime()
		
		case "3":
			ageCalc()

		case "4":
			fmt.Println("Thank you for using the calculator")
	
		default:
			fmt.Println("Please choose a valid value between 1-3 or 4 to exit")
        }
	}
}
//used for calculator 1: TimeDuration
func enterDate() CalcTime{//return CalcTime struct
	var input string
	var entered CalcTime

	//Clear calculator and start from the beginning 
	fmt.Println("Enter 'clear' to clear calculator")
	Message1:
		fmt.Print("Enter a day (in the format dd): ")
		fmt.Scan(&input)

	if input == "clear"{
		goto Message1
	}
	day, err := strconv.Atoi(input)
	if err != nil{//if numbers were not inputted, go back to message
		goto Message1
	}

	Message2:
		fmt.Print("Enter in an hour (in the format hh): ")
		fmt.Scan(&input)

	if input == "clear"{
		goto Message1
	}
	hour, err := strconv.Atoi(input)
	if err != nil{
		goto Message2
	}

	Message3:
		fmt.Print("Enter in a minute (in the format mm): ")
		fmt.Scan(&input)

	if input == "clear"{
		goto Message1
	}
	min, err := strconv.Atoi(input)
	if err != nil{
		goto Message3
	}

	Message4:
		fmt.Print("Enter in a second (in the format ss): ")
		fmt.Scan(&input)

	if input == "clear"{
		goto Message1
	}
	sec, err := strconv.Atoi(input)
	if err != nil{
		goto Message4
	}

	entered.createTime(day, hour, min, sec)//create a new struct
	return entered
}
//used for calculator 2 and 3
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
	if err != nil || month > 12 || month < 1{
		fmt.Println("Please enter a valid value between 1-12")
		goto Message2
	}
	Message3:
		fmt.Print("Enter a day (in the format dd): ")
		fmt.Scan(&input)

	if input == "clear"{
		goto Message1
	}
	day, err := strconv.Atoi(input)
	if err != nil || day > 31 || day < 1{
		fmt.Println("Please enter a valid value between 1-31")
		goto Message3
	}

	Message4:
		fmt.Print("Enter in an hour (in the format hh): ")
		fmt.Scan(&input)

	if input == "clear"{
		goto Message1
	}
	hour, err := strconv.Atoi(input)
	if err != nil || hour >= 24 || hour < 1 {
		goto Message4
	}

	Message5:
		fmt.Print("Enter in a minute (in the format mm): ")
		fmt.Scan(&input)

	if input == "clear"{
		goto Message1
	}
	min, err := strconv.Atoi(input)
	if err != nil || min >= 60 || min < 1{
		fmt.Println("Please enter a valid value between 1-59")
		goto Message5
	}

	Message6:
		fmt.Print("Enter in a second (in the format ss): ")
		fmt.Scan(&input)

	if input == "clear"{
		goto Message1
	}
	sec, err := strconv.Atoi(input)
	if err != nil || sec >= 60 || sec < 1{
		fmt.Println("Please enter a valid value between 1-59")
		goto Message6
	}

	date := time.Date(year, time.Month(month), day, hour, min, sec, int(0), time.Local)
	return date 
}
func timeDuration(){

	var input string
	var total CalcTime

	Message:
		fmt.Println("Would you like to add or subtract your dates?")
		fmt.Print("Enter 'add' or 'subtract': ")
		fmt.Scan(&input)
	
	if input != "add" && input != "subtract"{
		fmt.Println("Please enter either 'add' or 'subtract'")
		goto Message
	}

	choice1 := enterDate()
	choice2 := enterDate()


	if input == "add"{
		secondsTotal := choice1.seconds + choice2.seconds//add seconds 
		minutesTotal := choice1.minutes//assign minutes total to the first choice minutes
		hoursTotal := choice1.hours//assign hours total to the first choice hours
		daysTotal := choice1.days//assign days total to the the first choice days

		if secondsTotal >= 60{//if seconds is greater than a minute
			minutesTotal += secondsTotal/60 //divide seconds by 60 to determine minutes
											//add to minutesTotal
			secondsTotal = secondsTotal%60 //remainder of the seconds is the total 
		}
		minutesTotal += choice2.minutes
		if minutesTotal >= 60{
			hoursTotal += minutesTotal/60
			minutesTotal = minutesTotal%60
		}
		hoursTotal += choice2.hours
		if hoursTotal >= 24{
			daysTotal += hoursTotal/24
			hoursTotal = hoursTotal%24
		}
		daysTotal += choice2.days

		total.seconds = secondsTotal
		total.minutes = minutesTotal
		total.hours = hoursTotal
		total.days = daysTotal
	}else{

		if choice1.days > choice2.days{
			total.days = choice1.days - choice2.days
			total.hours = choice1.hours - choice2.hours
			total.minutes = choice1.minutes - choice2.minutes
			total.seconds = choice1.seconds - choice2.seconds
		}else{
			total.days = (choice1.days - choice2.days) * (-1)
			total.hours = (choice1.hours - choice2.hours) * (-1)
			total.minutes = (choice1.minutes - choice2.minutes) * (-1)
			total.seconds = (choice1.seconds - choice2.seconds) * (-1)
		}
	}
	fmt.Println(total.days, "days", total.hours, "hours", total.minutes, "minutes", total.seconds, "seconds")

	secondsInDay := 86400.0//how man seconds are in a day
	totalTime := float64(total.days) + float64(total.hours*3600 + total.minutes*60 + total.seconds)/secondsInDay
	//convert hours to seconds and minutes to seconds. Add everything together, divide by secondsInDay and add to total.days
	fmt.Println(totalTime, "days")
	totalTime *= 24.0
	//multiple totalTime by 24 to convert days to hours
	fmt.Println(totalTime, "hours")
	totalTime*=60.0
	fmt.Println(totalTime, "minutes")
	totalTime*=60.0
	fmt.Println(totalTime, "seconds")
}

func addSubtractTime(){
	var input string
	Message:
		fmt.Println("Would you like to add or subtract your dates?")
		fmt.Print("Enter 'add' or 'subtract': ")
		fmt.Scan(&input)
	
	if input != "add" && input != "subtract"{
		fmt.Println("Please enter either 'add' or 'subtract'")
		goto Message
	}
	date := enterFullDate()
	
	var choice string
	Message1:
		fmt.Print("Enter in number of days to ", input, " (dd): ")
		fmt.Scan(&choice)

	if choice == "clear"{
		goto Message1
	}
	days, err := strconv.Atoi(choice)
	if err != nil{
		goto Message1
	}

	Message2:
		fmt.Print("Enter in hours ", input, " (in the format hh): ")
		fmt.Scan(&choice)

	if choice == "clear"{
		goto Message1
	}
	hour, err := strconv.Atoi(choice)
	if err != nil{
		goto Message2
	}

	Message3:
		fmt.Print("Enter in minutes ", input, " (in the format mm): ")
		fmt.Scan(&choice)

	if choice == "clear"{
		goto Message1
	}
	min, err := strconv.Atoi(choice)
	if err != nil{
		goto Message3
	}

	Message4:
		fmt.Print("Enter in seconds ", input, " (in the format ss): ")
		fmt.Scan(&choice)

	if input == "clear"{
		goto Message1
	}
	sec, err := strconv.Atoi(choice)
	if err != nil{
		goto Message4
	}
	addDays := 0
	addHours := 0
	addMinutes := 0

	if sec >= 60{
		addMinutes = sec/60
		sec = sec%60
		min += addMinutes
	}
	if min >= 60{
		addHours = min/60
		min = min%60
		hour += addHours
	}
	if hour >= 24{
		addDays = hour/24
		hour = hour%24
		days += addDays
	}

	if input == "add"{
		date = date.AddDate(0, 0, days)
		date = date.Add(time.Hour * time.Duration(hour))
		date = date.Add(time.Minute * time.Duration(min))
		date = date.Add(time.Second * time.Duration(sec))
	}else{
		
		date = date.AddDate(0, 0, -days)
		date = date.Add(-time.Hour * time.Duration(hour))
		date = date.Add(-time.Minute * time.Duration(min))
		date = date.Add(-time.Second * time.Duration(sec))
	}
	fmt.Println("The new date is:", date)
	fmt.Println("The weekday is", date.Weekday())
}

func ageCalc(){

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
