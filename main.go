package main

import (
	"fmt"
	"math/rand/v2"
)

var WordsDB = [...]string{
	"fair", "good", "well", "fine",
	"fail", "pass", "yoyo", "earn",
	"wins", "mark", "make", "luck",
	"wind", "fire", "take", "took",
	"king", "eggs", "chan", "made",
	"dare", "hell", "plan", "miss"}

func engine() {

	secret := WordsDB[rand.IntN(len(WordsDB))]
	letters := []rune(secret)
	won := false

	for i := 0; i < 6; i++ {

		var guess string
		fmt.Print("guess the word: ")

		for {
			_, err := fmt.Scan(&guess)
			if err != nil {
				fmt.Println("\nInput ended. Exiting game.")
				return

			}

			if len(guess) != 4 {
				fmt.Println("Please enter a 4-letter word.")
				continue
			}
			break
		}

		if guess == secret {
			fmt.Println("Congrate You Win")
			won = true
			break
		}

		lettersguess := []rune(guess)
		Used := make([]bool, len(letters))

		for i, guessletter := range lettersguess {

			if guessletter == letters[i] {
				fmt.Println(guessletter, "is in the correct place")
				Used[i] = true
				continue
			}

			found := false
			for j, letter := range letters {
				if guessletter == letter && !Used[j] {
					found = true
					Used[j] = true
					break
				}
			}

			if found {
				fmt.Println(guessletter, "is in the word but not in the correct place")
			} else {
				fmt.Println(guessletter, "is not in the word")
			}
		}

		fmt.Println("Wrong guess. Try another word.")

	}

	if !won {
		fmt.Println("GG buddy but you lost.")
		fmt.Println("The Word Was", secret)
	}

}

func main() {
	fmt.Println("Welcome to the Wordle game!")

	engine()

	fmt.Println("Thank you for playing!")

}
