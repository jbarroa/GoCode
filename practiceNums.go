package main
import (
	"fmt"
)

func main(){
	var temp int
	a := 6
	b := 4
	c := 2

	if a > b{
		temp = a
		a = b 
		b = temp
		fmt.Println("A", a)
		fmt.Println("B", b)
	}
	if b > c{
		temp = b 
		b = c
		c = temp
		fmt.Println("B", b)
		fmt.Println("C", c)
	}
	if a > b{
		temp = a
		a = b 
		b = temp
		fmt.Println("A", a)
		fmt.Println("B", b)
	}
	if (b-a) == (c-b){
		fmt.Println("This is true")
	}else{
		fmt.Println("This false")
	}
}