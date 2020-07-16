package main 

import (
	"fmt"
	"math/rand"
) 

func main(){

	var dogName string
	var stBernard int
	var chihuahua int
	var goldenRetriver int
	var husky int
	var kingDoberman int
    fmt.Print("What is your dog's name? ")
    fmt.Scan(&dogName)

	stBernard = rand.Intn(100)
	chihuahua = rand.Intn((100 - stBernard))
	goldenRetriver = rand.Intn((100 - stBernard - chihuahua))
	husky = rand.Intn((100 - stBernard - chihuahua - goldenRetriver))
	kingDoberman = (100 - stBernard - chihuahua - goldenRetriver - husky)

	fmt.Println(dogName, "is: ")
	fmt.Println(stBernard, "% St. Bernard")
	fmt.Println(chihuahua, "% Chihuahua")
	fmt.Println(goldenRetriver, "% Golden Retriever")
	fmt.Println(husky, "% Husky")
	fmt.Println(kingDoberman, "% King Doberman")
}