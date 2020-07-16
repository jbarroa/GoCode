//Jasmine Barroa 
//June 19, 2020
//Concurrency practice

package main
import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)
func ready(name string, channel chan string) {
	fmt.Println(name, "started getting ready")
	time1 := randNumbergenerator(6,9)
	time.Sleep(time.Duration(time1) * time.Second)
	fmt.Println(name, "Spent " +  strconv.Itoa(time1) + " seconds getting ready")
	channel <- "Ready"
}
func shoes(name string, channel chan string) {
	fmt.Println(name, "started putting on shoes")
	time1 := randNumbergenerator(3,4)
	time.Sleep(time.Duration(time1) * time.Second)
	fmt.Println(name, "Spent " +  strconv.Itoa(time1) + " seconds putting on shoes")
	channel <- "Ready"
}
func alarm(channel chan string) {
	fmt.Println("Arming Alarm")
	time.Sleep(6 * time.Second)
	fmt.Println("Alarm is armed")
	channel <- "Ready"
}
func randNumbergenerator(low int, high int) int {
	ns := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(ns)
	return generator.Intn(high - low+1) + low
}
func main() {
	ch := make(chan string)
	go ready("Alice",ch)
	go ready("Bob",ch)
	alice:= <-ch
	bob:= <-ch
	time.Sleep(4 * time.Second)
	if alice == "Ready" && bob == "Ready" {
		go alarm(ch)
		go shoes("Alice",ch)
		go shoes("Bob",ch)
		alice= <-ch
		bob= <-ch
		alarm:= <-ch
		if alarm == "Ready" {
		}
	}
}