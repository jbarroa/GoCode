//Jasmine Barroa 
//June 29, 2020
//Reading and writing to files

package main

import (
	"fmt"
	"os"
	"bufio"
	"io/ioutil"
	"encoding/json"
	"strconv"
)

type collection struct{
	collection []movies `json:"collection"`
}

type movies struct{
	movieId int `json:"id"`
	title string `json:"movie"`
	genre string `json:"genre"`
	releaseYear int `json:"releaseYear"`
	series bool `json:"series"`
}

func (c *collection) createFile(){

	if _, err := os.Stat("collection.json"); os.IsNotExist(err) {//check if collection.json exists
		os.Create("collection.json")//create file if it does not exist
	}
	jsonFile, err := os.Open("collection.json")//open and use json file
	if err != nil {//error opening file
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)//read the contents of the file
	if err != nil {
		panic(err)
	}

	json.Unmarshal(byteValue, c)//change json variable to interface value
}

func (c *collection) findMovie(id int) (int, movies, bool){//return index, the movie, and if it was found
	var index int
	var movie movies
	for i, v := range c.collection {//run through the whole collection
		if v.movieId == id {//if id is found
			index = i//set index
			movie = v//set the movie to return
			break
		}
	}
	if (movie == movies{}) {//if the id was not found
		return 0, movies{}, false
	}
	return index, movie, true
}

func (c *collection) display(){
	if len(c.collection) == 0 {
		fmt.Println()
		fmt.Println("You do not have any movies in your collection.")
		fmt.Println()
		return
	}
	for _, v := range c.collection {
		fmt.Println()
		fmt.Println("Movie Id:", v.movieId)
		fmt.Println("Title:", v.title)
		fmt.Println("Genre:", v.genre)
		fmt.Println("Release Year:", v.releaseYear)
		fmt.Println("In a series?", v.series)
	}
}

func (c *collection) addMovie(title1 string, genre1 string, year int, series1 bool){
	var id int
	size := len(c.collection)
	//set id number based on the last id number
	//id starts at 1 rather than 0
	if size == 0 {
		id = 1
	} else {
		id = c.collection[size-1].movieId + 1
	}
	//create a new movie in collection
	newMovie := movies{
		movieId: id,
		title: title1,
		genre: genre1,
		releaseYear: year,
		series: series1,
	}
	c.collection = append(c.collection, newMovie)//add to the collection
	c.commit()
}

//edit the title of a movie
func(c *collection) editTitle(id int, index int, movie movies){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What is the new movie title? ")
	scanner.Scan()
	tType := scanner.Text()
	movie.title = tType
	c.collection[index] = movie
	c.commit()
	fmt.Println()
}

//edit the genre of a movie 
func(c *collection) editGenre(id int, index int, movie movies){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What is the new genre? ")
	scanner.Scan()
	gType := scanner.Text()
	movie.genre = gType
	c.collection[index] = movie
	c.commit()
	fmt.Println()
}

//edit the release year of a movie
func(c *collection) editYear(id int, index int, movie movies){
	var input string
	Message:
		fmt.Print("When was the movie released? ")
		fmt.Scan(&input)
		year, err := strconv.Atoi(input)

	if err != nil{
		fmt.Println("Enter a valid year")
		goto Message
	}
	movie.releaseYear = year
	c.collection[index] = movie
	c.commit()
	fmt.Println()
}

//edit if the movie is in a series or not
func(c *collection) editSeries(id int, index int, movie movies){
	var input string
	Message:
		fmt.Print("Is this movie in a series(yes/no)? ")
		fmt.Scan(&input)
	
	if input == "no"{
		movie.series = false
	}else if input == "yes"{
		movie.series = true
	}else{
		fmt.Println("Enter 'yes' or 'no'")
		goto Message
	}
	c.collection[index] = movie
	c.commit()
	fmt.Println()
}

//delete a movie from the collection
func(c *collection) deleteMovie(id int, index int, movie movies){
	size := len(c.collection)
	//swap the movie and the last movie in the collection
	c.collection[size-1], c.collection[index] = c.collection[index], c.collection[size-1]
	//copy entire collection except for the last movie
	c.collection = c.collection[:size-1]
	c.commit()
	fmt.Println("The movie at", id, "has been deleted.")
}

//save file
func (c *collection) commit() error {
	f, err := json.MarshalIndent(c, "", "\t")//format output in json file
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("collection.json", f, 0644)//write it to collection.json file
	return err
}

func main(){
	var c collection
	c.createFile()
	var input string
	//Loop through menu until the user is done 
	//For ease of using program
	for input != "6"{
		fmt.Println("Menu:")
		fmt.Println("1. Display Collection")
		fmt.Println("2. Add Item")
		fmt.Println("3. Update Item")
		fmt.Println("4. Delete Item")
		fmt.Println("5. Testing")
		fmt.Println("6. Exit")
        fmt.Print("Choose a number 1-6: ")
        fmt.Scan(&input)

        switch input {
		case "1":
			c.display()
	
		case "2":
			addItem(&c)
		
		case "3":
			editItem(&c)

		case "4":
			deleteItem(&c)

		case "5":
			test(&c)

		case "6":
			fmt.Println("Thank you for viewing/updating the collection")
	
		default:
			fmt.Println("Please choose a valid value between 1-5 or 6 to exit")
        }
	}
}

func test(c *collection){
	c.addMovie("Iron Man", "Action", 2008, true)
	c.addMovie("Iron Man 2", "Action", 2010, false)//change series
	c.addMovie("Thor", "Action", 2011, true)
	c.addMovie("Star Wars", "Action", 1977, true)//delete from collection
	c.addMovie("Captain America", "Romance", 2012, true)//change genre
	c.addMovie("Justice League", "Action", 2012, true)//change title to Avengers
	c.addMovie("Infinity War", "Action", 2018, true)
	c.addMovie("Endgame", "Action", 2020, true)//change year to 2019
}
//get information to call add movie function
func addItem(c *collection){
	var input string
	var series bool
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What is the title? ")
	scanner.Scan()
	title := scanner.Text()
	fmt.Print("What is the genre? ")
	scanner.Scan()
	genre := scanner.Text()

	Message1:
		fmt.Print("When was the movie released? ")
		fmt.Scan(&input)
		year, err := strconv.Atoi(input)
	
	if err != nil{
		fmt.Println("Enter a valid year")
		goto Message1
	}

	Message2:
		fmt.Print("Is this movie in a series(yes/no)? ")
		fmt.Scan(&input)
	
	if input == "no"{
		series = false
	}else if input == "yes"{
		series = true
	}else{
		fmt.Println("Enter 'yes' or 'no'")
		goto Message2
	}
	c.addMovie(title, genre, year, series)
}

//choose what the user would like to update
func editItem(c *collection){
	var input string
	Message:
		fmt.Print("What is the id number of the movie you would like to update? ")
		fmt.Scan(&input)
		idNum, err := strconv.Atoi(input)
	if err != nil{
		fmt.Println("Please enter a valid number")
		goto Message
	}

	index, movie, isThere := c.findMovie(idNum)
	if isThere{//if the id is within the collection
		//choose what the user would like to edit
		Message1:
			fmt.Println("What whould you like to edit?")
			fmt.Println("1. Edit Title")
			fmt.Println("2. Edit Genre")
			fmt.Println("3. Edit Release Year")
			fmt.Println("4. Edit Series")
			fmt.Print("Choose a number 1-4: ")
			fmt.Scan(&input)
			user, err := strconv.Atoi(input)
		//call appropriate method
		if err != nil{//enter a number
			fmt.Println("Please enter a valid number between 1-4")
			goto Message1
		}else if user == 1{
			c.editTitle(idNum, index, movie)
		}else if user == 2{
			c.editGenre(idNum, index, movie)
		}else if user == 3{
			c.editYear(idNum, index, movie)
		}else if user == 4{
			c.editSeries(idNum, index, movie)
		}else{
			fmt.Println("Please enter a number between 1-4")
			goto Message1
		}
	}else{
		fmt.Println("The movie with id", idNum, "is not in your collection")
	}
}

//choose what the user would like to delete 
func deleteItem(c *collection){
	var input string
	Message:
		fmt.Print("What is the id number of the movie you would like to delete? ")
		fmt.Scan(&input)
		idNum, err := strconv.Atoi(input)
	if err != nil{
		fmt.Println("Please enter a valid number")
		goto Message
	}

	index, movie, isThere := c.findMovie(idNum)
	if isThere{//if the movie is there, delete movie
		c.deleteMovie(idNum, index, movie)
	}else{
		fmt.Println("The movie with id", idNum, "is not in your collection")
	}
}