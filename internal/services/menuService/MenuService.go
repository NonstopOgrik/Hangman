package menuService

import (
	"fmt"
	gameService "hangman/internal/services/gameService"
	"os"
)

func Start() {
	for {
		var state bool
		fmt.Println("1.Начать игру")
		fmt.Println("2.Выйти")
		fmt.Fscan(os.Stdin, &state)
		if state {
			gameService.GameLogic()
		} else {
			break
		}
	}
}
