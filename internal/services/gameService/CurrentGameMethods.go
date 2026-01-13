package gameservice

import (
	"bufio"
	"fmt"
	"hangman/internal/structures"
	"math/rand"
	"os"
	"strings"
)

func initCurrentGame() structures.CurrentGameStruct {
	game := structures.CurrentGameStruct{PublicMistakesLeft: 6, PublicTurn: 0, PublicWord: chooseTheWord()}
	println("Игра началась!")
	return game
}

func chooseTheWord() string {
	file, err := os.Open("internal/packets/words.txt")
	if err != nil {
		return ""
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	data := make([]string, 64)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	resultWord := string(data[rand.Intn(len(data))])
	for {
		if len(resultWord) <= 5 {
			resultWord = string(data[rand.Intn(len(data))])
		} else {
			break
		}
	}
	return resultWord
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
		fmt.Println("----------------------------------------------")
		printStatistics(myGame.PublicTurn, myGame.PublicMistakesLeft, wordAsMap)
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
		fmt.Println("----------------------------------------------")
	}
	printScore(myGame)
}

func printStatistics(turn, mistakes int, words map[int]structures.Letter) {

	fmt.Println("Ход номер: ", turn, "Ошибок осталось: ", mistakes)
	printedWord := ""
	for i := 0; i < len(words); i++ {
		if words[i].IsOpen {
			printedWord += words[i].PublicWord
		} else {
			printedWord += "*"
		}
		printedWord += " "
	}
	fmt.Println("Ваше слово: ", printedWord)

}
func turn(wordAsMap *map[int]structures.Letter) (bool, bool) {
	fmt.Print("Введите букву: ")
	var word string
	fmt.Fscan(os.Stdin, &word)
	word = strings.ToLower(word)
	rightWord := false
	gameState := true
	for i := 0; i < len(*wordAsMap); i++ {
		if (*wordAsMap)[i].PublicWord == word && (*wordAsMap)[i].IsOpen == true {
			fmt.Println("Вы уже вводили данную букву!")
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
	if rightWord {
		fmt.Println("Такая буква...")
		fmt.Println("Есть в слове!!!")
		return true, gameState
	} else {
		fmt.Println("Такая буква...")
		fmt.Println("Отсутствует в слове!!!")
		return false, gameState
	}

}

func printScore(myGame structures.CurrentGameStruct) {
	fmt.Println("----------------------------------------------")
	fmt.Println("Итог Игры!")
	if myGame.PublicGameResult {
		fmt.Println("Вы победили!!!")
	} else {
		fmt.Println("Вы проиграли!!!")
	}
	fmt.Println("Количество ходов: ", myGame.PublicTurn)
	fmt.Println("Ошибок осталось: ", myGame.PublicMistakesLeft)
	fmt.Println("Слово: ", myGame.PublicWord)
	fmt.Println("----------------------------------------------")
}
