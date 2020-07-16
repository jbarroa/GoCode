//Jasmine Barroa, Matthew Kobilas, and Joel Karei
//June 12, 2020
//Tic Tac Toe game

package main

import (
    "fmt"
    "math"
    "strconv"
    "strings"
)
var board = [3][3]string{{"0", "1", "2"}, {"3", "4", "5"}, {"6", "7", "8"}}
var player1 string = "Player1"
var player2 string = "Player2"
var winner bool = false
func main() {
    var guessString string
    for winner == false {
        //Player 1 guesses
        printBoard()
        fmt.Println(player1, "please enter the number of your guess:")
        fmt.Scanln(&guessString)
        guessInt, _ := strconv.Atoi(guessString)
        makeGuess(player1, guessInt)
        // printBoard()
        //fmt.Println(boardFull())
        if checkWin(player1) {
            break
        }
		if boardFull() {
			fmt.Println("It's a tie")
			break
		}

        //Player 2 guesses
        printBoard()
        fmt.Println(player2, "please enter the number of your guess:")
        fmt.Scanln(&guessString)
        guessInt, _ = strconv.Atoi(guessString)
        makeGuess(player2, guessInt)
        // printBoard()
        //fmt.Println(boardFull())
        if checkWin(player2) {
            break
        }
		if boardFull() {
			fmt.Println("It's a tie")
			break
		}

    }
}
func makeGuess(player string, guess int) {
    var marker string
    if player == "Player1" {
        marker = "x"
    }
    if player == "Player2" {
        marker = "o"
    }
    x := int(math.Trunc(float64(guess) / 3.0))
    y := (guess % 3)
    board[x][y] = marker
}
func printBoard() {
    println(board[0][0], board[0][1], board[0][2])
    println(board[1][0], board[1][1], board[1][2])
    println(board[2][0], board[2][1], board[2][2])
}
func checkWin(player string) bool {
    //Jasmine "just counter variables..."
    var colOne int
    var colTwo int
    var colThree int
    var row int
    var rightDiagonal int
    var leftDiagonal int
    //check columns
    for i := 0; i < len(board); i++ {
        if player == "Player1" {
            if board[i][0] == "x" {
                colOne++
            }
            if board[i][1] == "x" {
                colTwo++
            }
            if board[i][2] == "x" {
                colThree++
            }
        } else {
            if board[i][0] == "o" {
                colOne++
            }
            if board[i][1] == "o" {
                colTwo++
            }
            if board[i][2] == "o" {
                colThree++
            }
        }
    }
    //fmt.Println(colOne, colTwo, colThree)
    if colOne >= 3 || colTwo >= 3 || colThree >= 3 {
        fmt.Println(player, "has won!!")
        return true
    }
    //check rows
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            if player == "Player1" {
                if board[i][j] == "x" {
                    row++
                }
            } else {
                if board[i][j] == "o" {
                    row++
                }
            }
        }
        //fmt.Println(row)
        if row == 3 {
            fmt.Println(player, "has won!!")
            return true
        }
        row = 0
    }
    //check diagonals
    for i := 0; i < len(board); i++ {
        if player == "Player1" {
            if board[i][i] == "x" {
                rightDiagonal++
            }
            if board[i][2-i] == "x" {
                leftDiagonal++
            }
        } else {
            if board[i][i] == "o" {
                rightDiagonal++
            }
            if board[i][2-i] == "o" {
                leftDiagonal++
            }
        }
    }
    if rightDiagonal == 3 || leftDiagonal == 3 {
        fmt.Println(player, "has won!!")
        return true
    }
    return false
}
func boardFull() bool {
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board); j++ {
            if strings.ContainsAny(board[i][j], "012345678") {
                return false
            }
        }
    }
    return true
}
