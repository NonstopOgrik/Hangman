package gameService

import (
	printService "hangman/internal/services/printService"
	readFileService "hangman/internal/services/readFileService"
	readFromConsoleService "hangman/internal/services/readFromConsoleService"
	"hangman/internal/structures"
)

func initCurrentGame() structures.CurrentGameStruct {
	game := structures.CurrentGameStruct{PublicMistakesLeft: 6, PublicTurn: 0, PublicWord: readFileService.ChooseTheWord()}
	printService.StartGamePrint()
	return game
}

func wordToMap(stringWord string) map[int]structures.Letter {
	resultMap := make(map[int]structures.Letter, len(stringWord))
	wordIndex := 0
	for _, n := range stringWord {
		tempLetter := structures.Letter{IsOpen: false, PublicWord: string(n)}
		resultMap[wordIndex] = tempLetter
		wordIndex++
	}
	return resultMap
}
func GameLogic() {
	myGame := initCurrentGame()
	wordAsMap := wordToMap(myGame.PublicWord)
	for {
		printService.TurnBorders()
		printService.PrintStatistics(myGame.PublicTurn, myGame.PublicMistakesLeft, wordAsMap)
		turnResult, GameState := turn(&wordAsMap)
		if turnResult {
			myGame.PublicTurn++
		} else {
			myGame.PublicMistakesLeft--
			myGame.PublicTurn++
		}
		if myGame.PublicMistakesLeft == 0 && !GameState {
			myGame.PublicGameResult = false
			break
		}
		if GameState {
			myGame.PublicGameResult = true
			break
		}
		printService.TurnBorders()
	}
	printService.PrintScore(myGame)
}

func turn(wordAsMap *map[int]structures.Letter) (bool, bool) {
	printService.ChooseTheWord()
	var word = readFromConsoleService.ConsoleReadWord()
	rightWord := false
	gameState := true
	for i := 0; i < len(*wordAsMap); i++ {
		if (*wordAsMap)[i].PublicWord == word && (*wordAsMap)[i].IsOpen == true {
			printService.WordRePeakPrint()
			return true, false
		}
		if (*wordAsMap)[i].PublicWord == word && (*wordAsMap)[i].IsOpen != true {
			letter := (*wordAsMap)[i]
			letter.IsOpen = true
			(*wordAsMap)[i] = letter
			rightWord = true
		}
		gameState = gameState && (*wordAsMap)[i].IsOpen

	}
	printService.PrintChoosenWordInfo(rightWord)
	if rightWord {
		return true, gameState
	} else {
		return false, gameState
	}

}
