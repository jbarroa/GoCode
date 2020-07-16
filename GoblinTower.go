//Jasmine Barroa
//June 17, 2020
//Goblin Tower game

package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

type hero struct {
	health int
	maxHealth int
	attack int
	defense int
	maxDefense int
	potions [5]int 
	gold int
	goblinsKilled int
	level int
}

type goblin struct {
	health int
	attack int
	defense int
}

func (player *hero) createHero(){

	ns := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(ns)

	max := 30
	min := 20
	heal := generator.Intn(max - min) + min
	player.health = heal
	player.maxHealth = heal

	max = 3
	min = 1
	player.attack = generator.Intn(max - min) + min
	
	max = 5
	min = 1
	defend := generator.Intn(max - min) + min
	player.defense = defend
	player.maxDefense = defend

	player.potions[0] = 2
	player.potions[1] = 2
	player.potions[2] = 2
	player.potions[3] = 2
	player.potions[4] = 2

	player.gold = 0
	player.goblinsKilled = 0
	player.level = 1
}

func (g *goblin) createGoblin(){

	ns := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(ns)

	max := 10
	min := 5
	g.health = generator.Intn(max - min) + min

	max = 3
	min = 2
	g.attack = generator.Intn(max - min) + min
	
	max = 2
	min = 1
	g.defense = generator.Intn(max - min) + min
}

func main(){

	var enemy goblin
	var player hero
	var input string
	var steps int
	var rounds int
	var amountOfGold int
	for{
		if rounds == 0{
			player.createHero() //creates a new hero character
		}else{
			player.createHero() //creates a new hero character
			player.gold = amountOfGold //keeps the same amount of gold from last character
		}

		//Prints out the stats for the character 
		fmt.Println()
		fmt.Println("The health of your player is:", player.health)
		fmt.Println("The attack of your player is:", player.attack)
		fmt.Println("The defense of your player is:", player.defense)
		fmt.Println("The amount of gold you have is:", player.gold)
		fmt.Println()

		for{
			steps++ //character takes a step
		//	fmt.Println("You have taken:", steps, "steps")
			meet := rand.Intn(4)

			if meet % 4 == 0{//25% chance of seeing a goblin on every step
				enemy.createGoblin()//creates a goblin to fight the character
				for{
					//hero fights goblin
					if enemy.defense > 0{
						enemy.defense -= player.attack//decreases the enemy's defense first
						if enemy.defense < 0{
							enemy.health += enemy.defense
						}
					}else{
						enemy.health -= player.attack//decreases the enemy's health 
					}

					/*fmt.Println("Goblin defense:", enemy.defense)
					fmt.Println("Goblin health:", enemy.health)*/

					if enemy.health <= 0{//goblin is dead
						fmt.Println("You killed a goblin")
						player.goblinsKilled++ //adds to the hero's goblin kill count
						player.gold += 2 //gives 2 pieces of gold to hero
						fmt.Println("The amount of gold you have is:", player.gold)
						fmt.Println("Your health is at", player.health)
						fmt.Print("Would you like to take a potion for +2 points (yes/no)? ")
						fmt.Scan(&input)

						if input == "yes"{
							numPotions := player.checkPotions()//checks how many potions you have
							if numPotions > 0{
								player.takePotions(numPotions)
								fmt.Println("Your new health after potions is", player.health)
							}else{
								fmt.Println("You do not have any potions")
							}
						}
						break
					}	
					//goblin fights hero
					if player.defense > 0{
						player.defense -= enemy.attack//decreases the hero's defense first
						if player.defense < 0{
							player.health += player.defense
						}
					}else{
						player.health -= enemy.attack//decreases the hero's health
					}

				/*	fmt.Println("Hero defense:", player.defense)
					fmt.Println("Hero health:", player.health)*/

					if player.health <= 0{//hero is dead
						fmt.Println("Oh no! You were killed :(")
						break
					}
				}
			}

			if player.health <= 0{
				break
			}

			if steps % 10 == 0{//every 10 steps, the hero will level up
				fmt.Println("Congratulations! You leveled up!")
				//steps = 0
				player.level++
				player.defense = player.maxDefense
				fmt.Println("You are now on level", player.level)
				num := player.checkPotions()
				if num < 5{//if the hero has less than 5 potions, ask the hero if they would like to go to the potion shop
					Message:
						fmt.Print("Would you like to go to the potion shop to buy more potions (yes/no)? ")
						fmt.Scan(&input)
					if input == "yes"{
						if player.gold >= 4{//if the user has enough gold to buy one potion
							player.potionShop(num)
						}else{
							fmt.Println("You do not have enough gold to buy more potions")
						}
					} else if input != "no"{
						fmt.Println("Please enter either yes or no")
						goto Message
					}
				}else{//cannot go to the potion shop if they still have 5 potions
					fmt.Println("You already have 5 potions")
				}
			}//ends steps if statement
		}//ends for loop once player dies

		fmt.Print("Would you like to play again (yes/no)? ")
		fmt.Scan(&input)

		if input != "yes"{
			fmt.Println("You ended at level", player.level)
			fmt.Println("You killed", player.goblinsKilled, "goblin(s)")
			break
		}

		amountOfGold = player.gold
		rounds++
	}//ends game for loop
}

func (player *hero) takePotions(count int){

	var input string 
	Message1:
		fmt.Println("You have", count, "potions")
		fmt.Print("How many potions would you like to take? ")
		fmt.Scan(&input)
		choice, _ := strconv.Atoi(input)
	

	if choice <= 0 || choice > count{
		fmt.Println("Enter a valid amount")
		goto Message1
	}else{
		k := 0
		for j := 0; j < len(player.potions); j++{
			if k < choice && player.potions[j] == 2{
				if player.health < player.maxHealth{
					player.health += 2
					player.potions[j] = 0
					k++
				}else{
					fmt.Println("You are already at max health")
				}
				if k <= choice && player.health > player.maxHealth{//if the last potion had caused the health to go over the max health
					fmt.Println("You are already at max health. Returning the potion to you")
					player.health -= 2//goes back to health prior to going over maxHealth
					player.potions[j] = 2//returns the potion
					break
				}
			}
		}
	}
}

func (player *hero) potionShop(count int){
	var input string
	Message1:
		fmt.Println("You have", count, "potions")
		fmt.Println("The amount of gold you have is:", player.gold)
		fmt.Print("How many potions would you like to buy at 4 gold pieces? ")
		fmt.Scan(&input)
		choice, _ := strconv.Atoi(input)
	
	if choice <= 0 || choice > 5{
		fmt.Println("Enter a valid amount")
		goto Message1
	}
	cost := choice * 4
	if cost > player.gold{
		fmt.Println("You do not have enough funds to buy", choice, "potions.")
		goto Message1
	}else{
		k := 0
		for j := 0; j < len(player.potions); j++{
			if k < choice && player.potions[j] == 0{
				player.gold -= 4
				player.potions[j] = 2
				k++
			}
		}
		fmt.Println("The amount of gold you have after buying", choice, "postions is:", player.gold)
		fmt.Print("Would you like to take a potion for +2 points (yes/no)? ")
		fmt.Scan(&input)

		if input == "yes"{
			numPotions := player.checkPotions()//checks how many potions you have
			player.takePotions(numPotions)
			fmt.Println("Your new health after potions is", player.health)
		}
	}
}

func (player *hero) checkPotions () int{
	var count int	
	for i := 0; i < len(player.potions); i++{
		if player.potions[i] == 2{
			count++
		}
	}
	return count
}