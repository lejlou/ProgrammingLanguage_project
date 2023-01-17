package main

import (
	"fmt"
)

func printBoard(board [6][7]string) {
	fmt.Println("---------------")
	for i := range board {
		for j := range board[i] {
			fmt.Printf("|" + board[i][j])
		}
		fmt.Println("|")
		fmt.Println("---------------")
	}

}

func fillEmptyBoard(board [6][7]string) [6][7]string {
	for i := range board {
		board[i] = [7]string{" ", " ", " ", " ", " ", " ", " "}
	}
	return board
}

func colDrop(board [6][7]string, column int, turn string) [6][7]string {

	for i := 5; i >= 0; i-- {
		if board[i][column] == " " {
			board[i][column] = turn
			break
		}
	}

	return board
}

func isColFull(board [6][7]string, column int) bool {
	for i := 5; i >= 0; i-- {
		if board[i][column] == " " {
			return false
		}
	}
	return true
}

func validateInput(board [6][7]string, column int) bool {
	if column > 6 || column < 0 {
		fmt.Println("That column doesn't exist.")
		return false
	}

	if isColFull(board, column) {
		fmt.Println("The column is full.")
		return false
	}

	return true
}

func horizontalCheck(board [6][7]string, column int, turn string) bool {
	var won bool = false
	var startRow int = 5
	var positions [4]int = [4]int{-3, -2, -1, 0}

	for i := 4; i >= 0; i-- {
		if board[i][column] != " " {
			startRow = i
		}
	}

	for _, position := range positions {
		if column+position < 0 || column+position+3 > 6 {
			continue
		}

		for i := 0; i < 4; i++ {
			if board[startRow][column+position+i] != turn {
				break
			}
			if i == 3 {
				won = true
			}
		}
	}

	return won
}

func verticalCheck(board [6][7]string, column int, turn string) bool {
	var counter int = 0
	for row := 0; row < 6; row++ {
		if board[row][column] != turn {
			counter = 0
		} else {
			counter++
		}
		if counter == 4 {
			return true
		}
	}

	return false
}

func diagonalCheck(board [6][7]string, column int, turn string) bool {
	//var won bool = false
	var startRow int = 5
	var positions [4]int = [4]int{-3, -2, -1, 0}

	for i := 4; i >= 0; i-- {
		if board[i][column] != " " {
			startRow = i
		}
	}

	for _, position := range positions {
		if column+position < 0 || column+position+3 > 6 {
			continue
		}
		// fmt.Printf("Checking position %v\n", position)
		// fmt.Printf("Starting row is %v\n", startRow)

		// Upwards diagonal
		if startRow-position-3 >= 0 && startRow-position <= 5 {
			for i := 0; i < 4; i++ {
				if board[startRow-position-i][column+position+i] != turn {
					break
				}
				if i == 3 {
					return true
				}
			}
		}

		// Downwards diagonal
		if startRow+position >= 0 && startRow+position+3 <= 5 {
			for i := 0; i < 4; i++ {
				if board[startRow+position+i][column+position+i] != turn {
					break
				}
				if i == 3 {
					return true
				}
			}
		}
	}

	return false
}

func isBoardFull(board [6][7]string) bool {
	for i := 0; i < 7; i++ {
		if !isColFull(board, i) {
			return false
		}
	}
	return true
}

func main() {

	var isGameOver bool = false
	var board [6][7]string
	turn := "O"
	var colChoice int

	board = fillEmptyBoard(board)
	fmt.Println("Welcome to Connect-4 in Go!")

	for !isGameOver {

		printBoard(board)

		var validInput bool = false
		fmt.Printf("It's %v's turn. ", turn)

		for !validInput {
			fmt.Println("Choose your column. (0 - 6)")
			fmt.Scan(&colChoice)
			validInput = validateInput(board, colChoice)
		}

		board = colDrop(board, colChoice, turn)

		if horizontalCheck(board, colChoice, turn) || verticalCheck(board, colChoice, turn) || diagonalCheck(board, colChoice, turn) {
			isGameOver = true
			printBoard(board)
			fmt.Printf("Game Over. %v has won!", turn)
		}
		if isBoardFull(board) {
			isGameOver = true
			printBoard(board)
			fmt.Printf("Game Over. It's a tie!")
		}
		if turn == "O" {
			turn = "0"
		} else {
			turn = "O"
		}
	}

}
