//Jasmine Barroa
//June 16, 2020
//Data Collections activities for module 4

package main

import (
	"fmt"
//	"strings"
	"math/rand"
	"time"
)
func main(){
	
	var max int
	var maxI int //max index
	var min int
	var minI int //min index
//	var avg float64 
//	var median int 

	random := randomArray(10)

	for i := 0; i < len(random); i++{
		fmt.Println(random[i])
	}
	
	max = maxArray(random)
	fmt.Println("The max of the array is:", max)

	maxI = maxIndex(random)
	fmt.Println("The max index of the array is:", maxI)

	min = minArray(random)
	fmt.Println("The min of the array is", min)

	minI = minIndex(random)
	fmt.Println("The min index of the array is:", minI)
}

func randomArray(size int) []int{
	
	newArr := make([]int, size)
	max := 100
	min := -100

	ns := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(ns)


	for i := 0; i < size; i++{
		newArr[i] = generator.Intn(max - min) + min
	}

	return newArr
}

func maxArray(findArr []int) int{

	max := findArr[0]

	for i := 0; i < len(findArr); i++{
		if findArr[i] > max{
			max = findArr[i]
		}
	}

	return max
}

func maxIndex(findArr []int) int{

	max := findArr[0]
	index := 0

	for i := 0; i < len(findArr); i++{
		if findArr[i] > max{
			max = findArr[i]
			index = i
		}
	}

	return index
}

func minArray(findArr []int) int{

	min := findArr[0]

	for i := 0; i < len(findArr); i++{
		if findArr[i] < min{
			min = findArr[i]
		}
	}

	return min
}

func minIndex(findArr []int) int{
	
	min := findArr[0]
	index := 0 

	for i := 0; i < len(findArr); i++{
		if findArr[i] < min{
			min = findArr[i]
			index = i
		}
	}

	return index
}

/*func sortDesc(){

}

func sortAsc(){

}

func meanArray() float64{

}

func medianArray() int{

}*/