package main
import (
	"fmt"
	"time"
)

type philsopher struct{
	name string
	left int
	right int
	haveFork bool
} 
type fork bool

type spot struct{

}

func (p *philsopher) eat(forks *[5]fork, ch chan string){
	_ = <- ch
	
	fmt.Println(p.name, "is waiting")
	fmt.Println(*forks)

	for !(forks[p.left]) && !(forks[p.right]){
		if forks[p.left] == true{
			forks[p.left] = false
			p.haveFork = true
		}else if forks[p.right] == true{
			forks[p.right] = false
			p.haveFork = true
		}
	}
	p.haveFork = true
	forks[p.left] = false
	forks[p.right] = false
	fmt.Println(*forks)
	fmt.Println(p.name, "is eating")
	time.Sleep(60* time.Second)
	fmt.Println("Philosopher ate")
	p.haveFork = false
	forks[p.left] = true
	forks[p.right] = true

	ch <- "Done"
}

func main(){
	
	var forks = [5]fork{true, true, true, true, true}
	var a = philsopher{"A", 0, 1, false}
	var b = philsopher{"B", 1, 2, false}
	var c = philsopher{"C", 2, 3, false}
	var d = philsopher{"D", 3, 4, false}
	var e = philsopher{"E", 4, 0, false}
	var philos = [5]philsopher{a, b, c, d, e}
	
	cha := make(chan string)
	cha <- "Done"
	chb := make(chan string)
	chb <- "Done"
	chc := make(chan string)
	chc <- "Done"
	chd := make(chan string)
	chd <- "Done"
	che := make(chan string)
	che <- "Done"

	var eatPhilos = [5]chan string{cha, chb, chc, chd, che}

	for !(a.haveFork == true && b.haveFork == true && c.haveFork == true && d.haveFork == true && e.haveFork == true){
		fmt.Println("Beginning outer loop")
		for i := 0; i < 5; i++{
			fmt.Println(i)
			val := <- eatPhilos[i]
			if val == "Done"{
				go philos[i].eat(&forks, eatPhilos[i])
			}
		}
		
	}
	fmt.Println("Done running")
}