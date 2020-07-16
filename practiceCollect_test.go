//Jasmine Barroa
//July 6, 2020
//

package main

import (
   "testing"
   "fmt"
   "strconv"
)

//Test Case 1: 
//Create order with no errors
func TestAdd(t *testing.T) {
   var c collection
	c.createFile()
   ans := c.addMovie("Iron Man", "Action", 2008, true)
   c.addMovie("Iron Man 2", "Action", 2010, true)
   c.addMovie("Thor", "Action", 2011, true)
   c.addMovie("Captain America", "Action", 2012, true)
   c.addMovie("Avengers", "Action", 2012, true)
   c.addMovie("Infinity War", "Action", 2018, true)
   c.addMovie("Endgame", "Action", 2019, true)
   if ans != 1 {
     // we use t to record the testing error
     t.Errorf("addItem(&c) = %d; Should be 0", ans)    
   }
}

func TestFind(t *testing.T){
   var c collection
	c.createFile()
   c.addMovie("Iron Man", "Action", 2008, true)
   c.addMovie("Iron Man 2", "Action", 2010, true)
   c.addMovie("Thor", "Action", 2011, true)
   c.addMovie("Captain America", "Action", 2012, true)
   c.addMovie("Avengers", "Action", 2012, true)
   c.addMovie("Infinity War", "Action", 2018, true)
   c.addMovie("Endgame", "Action", 2019, true)
	ans, movie, there := c.findMovie(2)
	if there == false{
      t.Errorf("findMovie(2) = %d; Should be 2", ans)
      fmt.Println("Movie:", movie)
   }
}

func TestDelete(t *testing.T){
   var c collection
	c.createFile()
   c.addMovie("Iron Man", "Action", 2008, true)
   c.addMovie("Iron Man 2", "Action", 2010, true)
   c.addMovie("Thor", "Action", 2011, true)
   c.addMovie("Captain America", "Action", 2012, true)
   c.addMovie("Star Wars", "Action", 1977, true)
   c.addMovie("Avengers", "Action", 2012, true)
   c.addMovie("Infinity War", "Action", 2018, true)
   c.addMovie("Endgame", "Action", 2019, true)
   index, movie, there := c.findMovie(5)
   if there{
      ans := c.deleteMovie(5, index, movie)
      if ans != "The movie at " + strconv.Itoa(5) + " has been deleted."{
         t.Errorf("c.deleteMovie(5, index, movie) = %s; should be The movie at 5 has been deleted.", ans)
      }
   }
}
/*func TestEditTitle(t *testing.T){
   var c collection
	c.createFile()
   c.addMovie("Iron Man", "Action", 2008, true)
	c.addMovie("Iron Man 2", "Action", 2010, true)
	c.addMovie("Thor", "Action", 2011, true)
	c.addMovie("Captain America", "Action", 2012, true)
	c.addMovie("Justice League", "Action", 2012, true)//change title to Avengers
	c.addMovie("Infinity War", "Action", 2018, true)
   c.addMovie("Endgame", "Action", 2020, true)//change year to 2019
   c.editTitle(5, 4, {"Justice League", "Action", 2012, true})
}

func TestEditYear(t *testing.T){
   var c collection
	c.createFile()
   c.addMovie("Iron Man", "Action", 2008, true)
	c.addMovie("Iron Man 2", "Action", 2010, true)
	c.addMovie("Thor", "Action", 2011, true)
	c.addMovie("Captain America", "Action", 2012, true)
	c.addMovie("Justice League", "Action", 2012, true)//change title to Avengers
	c.addMovie("Infinity War", "Action", 2018, true)
	c.addMovie("Endgame", "Action", 2020, true)//change year to 2019
}

func TestEditGenre(t *testing.T){
   var c collection
	c.createFile()
   c.addMovie("Iron Man", "Action", 2008, true)
	c.addMovie("Iron Man 2", "Romance", 2010, true)
	c.addMovie("Thor", "Action", 2011, true)
	c.addMovie("Captain America", "Action", 2012, true)
	c.addMovie("Justice League", "Action", 2012, true)//change title to Avengers
	c.addMovie("Infinity War", "Action", 2018, true)
	c.addMovie("Endgame", "Action", 2020, true)//change year to 2019
}*/

