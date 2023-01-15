package connect4project

import (
	"fmt"
)

func main() {
	match := NewGame()
	input := 0
	for match.gameOn {
		match.PrintGrid()
		fmt.Print("Which col?: ")
		_, err := fmt.Scanf("%d", &input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		match.PlaceStone(int(input))
		if match.CheckWin() != "" {
			fmt.Println(match.CheckWin(), "wins!")
		}
	}
}
