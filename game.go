package connect4project

import (
	"fmt"
	"math"
)

type Game struct {
	grid   [][]string
	turn   int
	gameOn bool
	cols   int
	rows   int
}

func NewGame() Game {
	newGame := Game{}
	newGame.grid = [][]string{
		{"e", "e", "e", "e", "e", "e", "e"},
		{"e", "e", "e", "e", "e", "e", "e"},
		{"e", "e", "e", "e", "e", "e", "e"},
		{"e", "e", "e", "e", "e", "e", "e"},
		{"e", "e", "e", "e", "e", "e", "e"},
		{"e", "e", "e", "e", "e", "e", "e"},
	}
	// REPLACE eXO with a "Stone" struct

	newGame.turn = 1
	newGame.gameOn = true
	newGame.cols = 7
	newGame.rows = 6
	return newGame
}

func (game *Game) EvenTurn() bool {
	if math.Mod(float64(game.turn), 2.0) == 0 {
		return true
	} else {
		return false
	}
}

func (game *Game) PrintGrid() {
	for i := 0; i <= 6; i++ {
		fmt.Print(i, " ")
	}
	fmt.Print("\n")
	for k := 0; k <= 13; k++ {
		fmt.Print("-")
	}
	fmt.Print("\n")
	for i := range game.grid {
		for k := range game.grid[i] {
			fmt.Print(game.grid[i][k], " ")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func (game *Game) PlaceStone(col int) {
	for i := 0; i <= game.cols-2; i++ {
		switch game.grid[i][col] {
		case "X", "O":
			if game.EvenTurn() {
				game.grid[i-1][col] = "O"
				break
			} else {
				game.grid[i-1][col] = "X"
				break
			}
		default:
			if i == game.cols-2 {
				if game.EvenTurn() {
					game.grid[i][col] = "O"
					break
				} else {
					game.grid[i][col] = "X"
					break
				}
			}
			continue
		}
		break
	}
	game.turn += 1
}

func (game *Game) CheckWin() string {
	// horizontal
	for _, line := range game.grid {
		for i := 0; i <= len(line)-4; i++ {
			if compare(line[i:i+4], "X") {
				game.gameOn = false
				return "X"
			} else if compare(line[i:i+4], "O") {
				game.gameOn = false
				return "O"
			}
		}
	}
	// vertical
	for k := 0; k < game.rows; k++ {
		line := game.GetHorizontalLine(k)
		for l := 0; l <= len(line)-4; l++ {
			if compare(line[l:l+4], "X") {
				game.gameOn = false
				return "X"
			} else if compare(line[l:l+4], "O") {
				game.gameOn = false
				return "O"
			}
		}
	}
	// diagonal l-r PROBLEM HERE
	for m := 0; m < game.cols; m++ {
		for n := 0; n < game.rows; n++ {
			if compare(game.GetDiagonalLR(m, n), "X") {
				game.gameOn = false
				return "X"
			} else if compare(game.GetDiagonalLR(m, n), "O") {
				game.gameOn = false
				return "O"
			}
		}
	}
	return ""
}

func (game *Game) GetHorizontalLine(col int) []string {
	slice := []string{}
	for i := 0; i < game.rows; i++ {
		slice = append(slice, game.grid[i][col])
	}
	return slice
}

func (game *Game) GetDiagonalLR(col int, row int) []string {
	slice := []string{}
	if (col+4 >= game.cols) || (row+4 >= game.rows) {
		return slice
	}
	slice = append(slice,
		game.grid[col][row],
		game.grid[col+1][row+1],
		game.grid[col+2][row+2],
		game.grid[col+3][row+3])
	return slice
}

func (game *Game) GetDiagonalRL(col int, row int) []string {
	slice := []string{}
	if (col-4 >= game.cols) || (row+4 >= game.rows) {
		return slice
	}
	slice = append(slice,
		game.grid[col][row],
		game.grid[col-1][row+1],
		game.grid[col-2][row+2],
		game.grid[col-3][row+3])
	return slice
}

func compare(slice []string, stone string) bool {
	if slice == nil {
		return false
	}
	for _, char := range slice {
		if char != stone {
			return false
		}
	}
	return true
}
