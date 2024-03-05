package main

import (
	"bufio"
	"os"
	"strings"
)

type player struct {
	name   string
	symbol string
}

func winner(board [3][3]string) bool {
	if board[0][0] != "" && board[0][0] == board[0][1] && board[0][0] == board[0][2] {
		printBoard(board)
		return true
	}
	if board[1][0] != "" && board[1][0] == board[1][1] && board[1][0] == board[1][2] {
		printBoard(board)
		return true
	}
	if board[2][0] != "" && board[2][0] == board[2][1] && board[2][0] == board[2][2] {
		printBoard(board)
		return true
	}
	if board[0][0] != "" && board[0][0] == board[1][0] && board[0][0] == board[2][0] {
		printBoard(board)
		return true
	}
	if board[0][1] != "" && board[0][1] == board[1][1] && board[0][1] == board[2][1] {
		printBoard(board)
		return true
	}
	if board[0][2] != "" && board[0][2] == board[1][2] && board[0][2] == board[2][2] {
		printBoard(board)
		return true
	}
	if board[0][0] != "" && board[0][0] == board[1][1] && board[0][0] == board[2][2] {
		printBoard(board)
		return true
	}
	if board[0][2] != "" && board[0][2] == board[1][1] && board[0][2] == board[2][0] {
		printBoard(board)
		return true
	}
	return false
}

func printBoard(board [3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == "" {
				print(" ")
			} else {
				print(board[i][j])
			}
			if j < 2 {
				print(" | ")
			}
		}
		println()
		if i < 2 {
			println("---------")
		}
	}

}

// xo game
func main() {
	board := [3][3]string{{"", "", ""}, {"", "", ""}, {"", "", ""}}
	players := [2]player{
		{name: "Player 1", symbol: "X"},
		{name: "Player 2", symbol: "O"},
	}
	turn := 0
	for {
		printBoard(board)
		// get the input
		var x, y int
		print(players[turn%2].name, " turn: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		coordinates := strings.Split(input, " ")

		if len(coordinates) != 2 {
			println("Invalid input. Try again.")
			continue
		}
		x, y = int(coordinates[0][0]-'0'), int(coordinates[1][0]-'0')
		if x < 0 || x > 2 || y < 0 || y > 2 || board[x][y] != "" {
			println("Invalid input. Try again.")
			continue
		}
		board[x][y] = players[turn%2].symbol
		// check if the game is over
		if winner(board) {
			println(players[turn%2].name, "wins!")
			break
		}
		// check if the game is a draw
		if turn == 8 {
			printBoard(board)
			println("It's a draw!")
			break
		}
		turn++
	}
}
