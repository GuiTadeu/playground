package main

import (
	"fmt"
	"os"
	"os/exec"
)

var board [3][3]string

func main() {
	startBoard()
	printBoard()
	gameLoop()
}

type Player struct {
	number   uint8
	nickname string
	mark     string
}

func gameLoop() {

	var move uint8

	player1 := Player{1, "Player 01", "X"}
	player2 := Player{2, "Player 02", "O"}
	players := []Player{player1, player2}

	possibleWinner := Player{}

	for i := 0; i < 4; i++ {
		for _, selectedPlayer := range players {

			fmt.Print(selectedPlayer.nickname + ": ")
			fmt.Scan(&move)

			updateBoard(move, selectedPlayer.mark)
			printBoard()

			possibleWinner = checkWinner(selectedPlayer)
			if (possibleWinner != Player{}) {
				endGame(selectedPlayer)
			}
		}
	}

	fmt.Print(player1.nickname + ": ")
	fmt.Scan(&move)

	updateBoard(move, player1.mark)
	printBoard()

	if(possibleWinner == Player{}) {
		fmt.Println("Draw!")
	}
}

func endGame(winnerPlayer Player) {
	fmt.Println(winnerPlayer.nickname + " wins!!")
	os.Exit(3)
}

func checkWinner(selectedPlayer Player) (winnerPlayer Player) {

	possibleWinner := checkVertical(selectedPlayer)
	if (possibleWinner != Player{}) {
		winnerPlayer = possibleWinner
		return
	}

	possibleWinner = checkHorizontal(selectedPlayer)
	if (possibleWinner != Player{}) {
		winnerPlayer = possibleWinner
		return
	}

	possibleWinner = checkDiagonal(selectedPlayer)
	if (possibleWinner != Player{}) {
		winnerPlayer = possibleWinner
		return
	}

	return
}

func checkVertical(selectedPlayer Player) (winnerPlayer Player) {

	for y := 0; y < 3; y++ {

		count := 0
		for x := 0; x < 3; x++ {
			if board[x][y] == selectedPlayer.mark {
				count++
			}
		}

		if count == 3 {
			return selectedPlayer
		}
	}

	return Player{}
}

func checkHorizontal(selectedPlayer Player) (winnerPlayer Player) {

	for x := 0; x < 3; x++ {

		count := 0
		for y := 0; y < 3; y++ {
			if board[x][y] == selectedPlayer.mark {
				count++
			}
		}

		if count == 3 {
			return selectedPlayer
		}
	}

	return Player{}
}

func checkDiagonal(selectedPlayer Player) (winnerPlayer Player) {

	// Diagonal 01

	var count = 0
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if x == y {
				if board[x][y] == selectedPlayer.mark {
					count++
				}
			}
		}

		if count == 3 {
			return selectedPlayer
		}
	}

	// Diagonal 02

	var x = 0
	var y = 2
	count = 0
	for x <= 2 {

		if board[x][y] == selectedPlayer.mark {
			count++
		}

		if count == 3 {
			return selectedPlayer
		}

		x++
		y--
	}

	return Player{}
}

func startBoard() {
	var count = 1
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			board[x][y] = fmt.Sprint(count)
			count++
		}
	}
}

func printBoard() {

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			print(board[x][y] + "\t")
		}
		println()
	}
}

func updateBoard(move uint8, mark string) {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if board[x][y] == fmt.Sprint(move) {
				board[x][y] = mark
			}
		}
	}
}
