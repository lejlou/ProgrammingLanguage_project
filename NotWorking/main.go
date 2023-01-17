package main

import (
	"fmt"
)

const (
	rows    = 6
	columns = 7
)

// Piece represents a single piece on the game board.
type Piece int

const (
	Empty Piece = iota
	Blue
	Red
)

func (p Piece) String() string {
	switch p {
	case Blue:
		return "B"
	case Red:
		return "R"
	default:
		return " "
	}
}

// Game represents a game of 4 in a row.
type Game struct {
	board  *Board
	player Piece
}

// NewGame returns a new game.
func NewGame() *Game {
	return &Game{
		board:  NewBoard(),
		player: Blue,
	}
}

// Board represents the game board.
type Board [rows][columns]Piece

// NewBoard returns a new game board.
func NewBoard() *Board {
	return &Board{}
}

// Drop drops a piece into the given column.
func (b *Board) Drop(p Piece, col int) error {
	if col < 0 || col >= columns {
		return fmt.Errorf("invalid column: %d", col)
	}

	for i := rows - 1; i >= 0; i-- {
		if b[i][col] == Empty {
			b[i][col] = p
			return nil
		}
	}

	return fmt.Errorf("column %d is full", col)
}

// String returns a string representation of the board.
func (b *Board) String() string {
	var s string
	for i := rows - 1; i >= 0; i-- {
		s += "|"
		for j := 0; j < columns; j++ {
			s += b[i][j].String() + "|"
		}
		s += "\n"
	}
	s += "---------------\n"
	s += " 0 1 2 3 4 5 6\n"
	return s
}

// IsWin checks if the given player has won the game.
func (b *Board) IsWin(p Piece) bool {
	// Check for horizontal wins.
	for i := 0; i < rows; i++ {
		for j := 0; j < columns-3; j++ {
			if b[i][j] == p && b[i][j+1] == p && b[i][j+2] == p && b[i][j+3] == p {
				return true
			}
		}
	}

	// Check for vertical wins.
	for i := 0; i < rows-3; i++ {
		for j := 0; j < columns; j++ {
			if b[i][j] == p && b[i+1][j] == p && b[i+2][j] == p && b[i+3][j] == p {
				return true
			}
		}
	}

	// Check for diagonal wins (top left to bottom right).
	for i := 0; i < rows-3; i++ {
		for j := 0; j < columns-3; j++ {
			if b[i][j] == p && b[i+1][j+1] == p && b[i+2][j+2] == p && b[i+3][j+3] == p {
				return true
			}
		}
	}

	// Check for diagonal wins (top right to bottom left).
	for i := 0; i < rows-3; i++ {
		for j := columns - 1; j >= 3; j-- {
			if b[i][j] == p && b[i+1][j-1] == p && b[i+2][j-2] == p && b[i+3][j-3] == p {
				return true
			}
		}
	}

	// If no wins were found, return false.
	return false
}

// Play plays a game of 4 in a row.
func (g *Game) Play() {
	for {
		fmt.Println(g.board)
		fmt.Printf("Player %s's turn. Enter column number: ", g.player)

		var col int
		_, err := fmt.Scanf("%d", &col)
		if err != nil {
			fmt.Println("Invalid input. Try again.")
			continue
		}

		err = g.board.Drop(g.player, col)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if g.board.IsWin(g.player) {
			fmt.Println(g.board)
			fmt.Printf("Player %s wins!\n", g.player)
			break
		}

		// Check if the game is a draw.
		isDraw := true
		for i := 0; i < rows; i++ {
			for j := 0; j < columns; j++ {
				if g.board[i][j] == Empty {
					isDraw = false
					break
				}
			}
		}
		if isDraw {
			fmt.Println(g.board)
			fmt.Println("It's a draw!")
			break
		}

		// Switch to the other player.
		if g.player == Blue {
			g.player = Red
		} else {
			g.player = Blue
		}
	}
}

func main() {

	fmt.Println("Welcome to Connect 4!")
	var rows, columns int
	fmt.Print("Enter number of rows (default 6): ")
	fmt.Scan(&rows)
	if rows == 0 {
		rows = defaultRows
	}
	fmt.Print("Enter number of columns (default 7): ")
	fmt.Scan(&columns)
	if columns == 0 {
		columns = defaultColumns
	}

	game := NewGame()
	game.Play()
}
