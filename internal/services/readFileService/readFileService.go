package readfileservice

import (
	"bufio"
	"math/rand"
	"os"
)

func ChooseTheWord() string {
	file, err := os.Open("internal/resources/words.txt")
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
