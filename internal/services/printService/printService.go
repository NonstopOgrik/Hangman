package printservice

import (
	"fmt"
	"hangman/internal/structures"
)

func PrintStatistics(turn, mistakes int, words map[int]structures.Letter) {

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

func PrintScore(myGame structures.CurrentGameStruct) {
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

func PrintChoosenWordInfo(rightWord bool) {
	if rightWord {
		fmt.Println("Такая буква...")
		fmt.Println("Есть в слове!!!")
	} else {
		fmt.Println("Такая буква...")
		fmt.Println("Отсутствует в слове!!!")
	}
}

func TurnBorders() {
	fmt.Println("----------------------------------------------")
}

func ChooseTheWord() {
	fmt.Print("Введите букву: ")
}
func WordRePeakPrint() {
	fmt.Println("Вы уже вводили данную букву!")
}
func StartGamePrint() {
	println("Игра началась!")
}
