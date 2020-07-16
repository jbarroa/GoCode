//Jasmine Barroa
//June 21, 2020
//Tourist cafe

package main
import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)
func online(name string, channel chan string) {
	fmt.Println(name, "is online")
	time1 := randNumbergenerator(15,120)
	time.Sleep(time.Duration(time1) * time.Millisecond + time.Second)
	fmt.Println(name, "is done, having spent " +  strconv.Itoa(time1) + " minutes online.")
	channel <- name
}
func waiting(use chan string, done chan string, stop chan bool){
	computer := 0
	waitList := make([]string, 0)
	for{
	/*	tourist := <- use 
		if computer < 8{
			computer++
			go online(use, done)
		}else{
			time.Sleep(3 * time.Second)
			fmt.Println(tourist, "waiting for turn.")
			waitList = append(waitList, tourist)
		}
		ready := <- done
		if ready == "Ready"{
			computer --
			if len(waitList) > 0{
				next := waitList[0]
				waitList = waitList[1:]
				go func(){
					use <- next
				}()
			} else if computer == 0{
				return
			}
		}*/
		select {
		case tourist := <-use:
			if computer < 8 {
				computer++
				go online(tourist, done)
			} else {
				time.Sleep(time.Millisecond)
				fmt.Println(tourist, "waiting for turn.")
				waitList = append(waitList, tourist)
			}
		case <-done:
			computer--
			if len(waitList) > 0 {
				next := waitList[0]
				waitList = waitList[1:]
				go func() {
					use <- next
				}()
			} else if computer == 0 {
				stop <- true
				close(stop)
				return
			}

		}
	}
}
func shuffleArray(people []int)[]int {

	shuffled := make([]int, len(people))
	n := len(people)
	for i := 0; i < n; i++ {
		randIndex := rand.Intn(len(people))
		shuffled[i] = people[randIndex]
		people = append(people[:randIndex], people[randIndex+1:]...)
	}
	return shuffled
}
func randNumbergenerator(min int, max int) int {
	ns := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(ns)
	return generator.Intn(max - min) + min
}
func main() {
	touristList := make([]int, 25)
	for i := 0; i < len(touristList); i++{
		touristList[i] = i+1
	}
	touristList = shuffleArray(touristList)
	done := make(chan string)
	name := make(chan string)
	stop := make(chan bool)
	go waiting(name, done, stop)
	var allDone bool
	
	for i := 0; i < len(touristList); i++{
		name <- "Tourist " + strconv.Itoa(touristList[i])
	}
	allDone = <- stop
	
	if allDone == true {
		fmt.Println("The place is empty, let's close up and go to the beach!")
	}
}