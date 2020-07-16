//Jasmine Barroa
//June 17, 2020
//Blackjack game using structs, methods, functions 

package main
import (
    "fmt"
    "math/rand"
    "time"
)
var gameCardCounter = 52
var theDeck [52]card
type card struct {
    name  int
    suit  string
    value int
}
type player struct {
    name         string
    hiddenValue  int
    visibleValue int
}
func main() {
    player1 := player{"Matt", 0, 0}
    //array to hold deck
    theDeck = generateDeck()
    fmt.Println(theDeck)
    theDeck = shuffleDeck(theDeck)
    fmt.Println(theDeck)
    player1.dealCard()
    player1.dealCard()
    player1.dealCard()
    player1.dealCard()
}
func generateDeck() [52]card {
    var theDeck [52]card
    cardCounter := 52
    //assign cards
    for i := 0; i < 4; i++ {
        for j := 0; j < 13; j++ {
            //thisCard = theDeck[j]
            var thisCard card
            thisCard.name = j + 1
            if cardCounter < 13 {
                thisCard.suit = "Spades"
            }else if cardCounter < 26 {
                thisCard.suit = "Hearts"
            }else if cardCounter < 39 {
                thisCard.suit = "Diamonds"
            }else if cardCounter >= 39 {
                thisCard.suit = "Clubs"
            }
            if thisCard.name > 10 {
                thisCard.value = 10
            }
            if thisCard.name <= 10 {
                thisCard.value = thisCard.name
            }
            theDeck[cardCounter-1] = thisCard
            cardCounter--
        }
    }
    return theDeck
}
func shuffleDeck(theDeck [52]card) [52]card {
    rand.Seed(time.Now().UnixNano())
    rand.Shuffle(len(theDeck), func(i, j int) { theDeck[i], theDeck[j] = theDeck[j], theDeck[i] })
    return theDeck
}
func (thisPlayer *player) dealCard() {
    var thisCard card
    thisCard = theDeck[gameCardCounter-1]
    thisPlayer.hiddenValue += thisCard.value
    fmt.Println("Your card is the", thisCard.name, "of", thisCard.suit)
    fmt.Println(thisPlayer.hiddenValue)
    gameCardCounter--
}
/*generate card function
func generateCard(name int, value int, suit string) card {
    return card{name, value, suit}
}
*/
