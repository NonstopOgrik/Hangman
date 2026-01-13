package structures

type CurrentGameStruct struct {
	PublicMistakesLeft int
	PublicTurn         int
	PublicWord         string
	PublicGameResult   bool
	PublicGameState    bool
}

type Letter struct {
	PublicWord string
	IsOpen     bool
}
