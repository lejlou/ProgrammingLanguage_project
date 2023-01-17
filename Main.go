/* package connect4project

import (
    "fmt"
    "github.com/go-board/board"
    "github.com/go-board/grid"
)

type player struct {
    name string
    symbol string
    moves []string
}

const (
    defaultRows    = 6
    defaultColumns = 7
)

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
    board := board.New(rows, columns)
    player1 := player{name: "Player 1", symbol: "⚪", moves: []string{}}
    player2 := player{name: "Player 2", symbol: "⚫", moves: []string{}}
    currentPlayer := &player1
    for !board.IsFull() {
        fmt.Println("Current Board:")
        printBoard(board, player1, player2)
        fmt.Printf("%s, enter column:", currentPlayer.name)
        var col int
        fmt.Scan(&col)
        if !board.IsValidMove(col) {
            fmt.Println("Invalid move, please try again.")
            continue
        }
        row := dropCoin(board, col, currentPlayer)
        if checkWin(board, row, col) {
            fmt.Println(currentPlayer.name, "wins!")
            break
        }
        if currentPlayer == &player1 {
            currentPlayer = &player2
        } else {
            currentPlayer = &player1
        }
    }
    if board.IsFull() {
        fmt.Println("It's a tie!")
    }
    fmt.Println("Player 1 moves:", player1.moves)
    fmt.Println("Player 2 moves:", player2.moves)
    
    */
