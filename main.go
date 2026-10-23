package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	wordsDB := [...]string{
		"fair", "good", "well", "fine",
		"fail", "pass", "yoyo", "earn",
		"wins", "mark", "make", "luck",
		"wind", "fire", "take", "took",
		"king", "eggs", "chan", "made",
		"dare", "hell", "plan", "miss"}

	secret := wordsDB[rand.IntN(len(wordsDB))]

	var guess string
	fmt.Print("guess the word: ")

	for {
		_, err := fmt.Scan(&guess)
		if err != nil {
			break
		}

		if len(guess) != 4{
			fmt.Println("Please enter a 4-letter word.")
			continue
		}
		break
	}

	letters := []rune(secret)
	lettersguess := []rune(guess)

	for i, letter := range letters {

			if letter == lettersguess[i] {
				fmt.Println("this letter is in the correct place", i)
		}else if  {
			
		}

	}

}
