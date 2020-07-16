package main 

import (
	"fmt"
	"strconv"
) 

func main() {

	var ageS string;
	var age int;
	var targetMin float32;
	var targetMax float32;
	fmt.Print("What is your age? ")
	fmt.Scan(&ageS)

	age, _ = strconv.Atoi(ageS)//convert age to int

	max := (220 - age)
	targetMin = float32(max) * .5
	targetMax = float32(max) * .85

	fmt.Println("Your maximum heart rate should be", max, "beats per minute")
	fmt.Println("Your target HR Zone is" ,int(targetMin), "-" ,int(targetMax), "beats per minute")
}