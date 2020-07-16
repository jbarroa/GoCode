package main
import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)
var jobs = make(chan int, 10)
//var results = make(chan string, 10)

func inStore(customer int){
	fmt.Println("Customer", customer, "is in the store")
	num := randNumbergenerator(5, 60)
	time.Sleep(time.Duration(num) * time.Second)
	fmt.Println("Customer", customer, "left the store")
}
func worker(wg *sync.WaitGroup){
	for job := range jobs {
		inStore(job)
    //  results <- "Done"
    }
    wg.Done()
}
func createWorkerPool(noOfWorkers int) {  
    var wg sync.WaitGroup
    for i := 0; i < noOfWorkers; i++ {
        wg.Add(1)
        go worker(&wg)
    }
    wg.Wait()
    //close(results)
}
func allocate(noOfJobs int) {  
    for i := 1; i <= noOfJobs; i++ {
		fmt.Println(i, "is waiting in line")
		jobs <- i
		if i >= 4{
			num := randNumbergenerator(5, 10)
			time.Sleep(time.Duration(num) * time.Second)
		}	
    }
    close(jobs)
}
/*func result(done chan bool) {  
    for _ = range results {
		
    }
    done <- true
}*/
func digits(number int) int {  
    sum := 0
    no := number
    for no != 0 {
        digit := no % 10
        sum += digit
        no /= 10
    }
    time.Sleep(2 * time.Second)
    return sum
}
func randNumbergenerator(min int, max int) int {
	ns := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(ns)
	return generator.Intn(max - min) + min
}
func main(){
//	startTime := time.Now()
    noOfJobs := 15
    go allocate(noOfJobs)
 //   done := make(chan bool)
  //  go result(done)
    noOfWorkers := 5
    createWorkerPool(noOfWorkers)
//	<-done
	fmt.Println("We're closed")
}

