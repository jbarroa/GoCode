//Jasmine Barroa
//June 15, 2020
//Practice slices

package main

import (
	"fmt"
	"strings"
//	"math"
)
func main(){
	words := [10]string{}
	var wordLonger []string
	var wordShorter []string

	var length int
	var count int
	var average int
	var input string

	for i := 0; i < len(words); i++{
		fmt.Print("Enter a word: ")
		fmt.Scan(&input)
		input = strings.ToLower(input)
		words[i] = input
		length += len(input)
		count++
	}

	//average = int(math.Trunc(float64(length)/float64(count)))
	average = length/count
	fmt.Println("The average length of a word is:", average)
	
	for j := 0; j < len(words); j++{
		if len(words[j]) < average{
			wordShorter = append(wordShorter, words[j])
		}else if len(words[j]) > average{
			wordLonger = append(wordLonger, words[j])
		}
	}

	fmt.Println("The words longer than the average are:", wordLonger)
	fmt.Println("The words shorter than the average are:", wordShorter)
}