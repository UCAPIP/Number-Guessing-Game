package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

var timesToTry int

func scanDifficult() error {

	difficults := map[int]int{
		1: 10,
		2: 5,
		3: 3,
	}

	fmt.Println("\n" + `Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)`)

	fmt.Print("\nEnter your choice: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()
		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Reading error:", err)
			return err
		}
		if num > 3 || num < 1 {
			fmt.Println("Enter num of difficult")
			scanDifficult()
		}
		timesToTry = difficults[num]
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Reading error:", err)
		return err
	}

	switch timesToTry {
	case 10:
		fmt.Println("\nGreat! You have selected the Easy difficulty level.")
	case 5:
		fmt.Println("\nGreat! You have selected the Medium difficulty level.")
	case 3:
		fmt.Println("\nGreat! You have selected the Hard difficulty level.")
	}
	return nil
}

func endMessage() {
	fmt.Println("\nIf you want to play again - type Y or N for exit")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := strings.ToUpper(scanner.Text())
		switch text {
		case "Y":
			gameCycle()
		case "N":
			os.Exit(1)
		default:
			fmt.Println("UNKNOWN COMMAND")
			endMessage()
		}
	}
}

func gameCycle() {
	err := scanDifficult()
	if err != nil {
		fmt.Println(err)
		return
	}
	randomInt := rand.IntN(100)
	fmt.Println("Let's start the game!")
	start := time.Now()
	for i := 1; i <= timesToTry; i++ {
		fmt.Println("Answer: ", randomInt) // -- Ответ
		fmt.Print("Enter your guess: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			text := scanner.Text()
			num, err := strconv.Atoi(text)
			if err != nil {
				fmt.Println("Reading error:", err)
				return
			}
			if num == randomInt {
				duration := time.Since(start)
				fmt.Printf("Congratulations! You guessed the correct number in %v attempts. Time spent %s", i, duration)
				endMessage()
			}
			if num < randomInt {
				fmt.Printf("Incorrect! The number is greater than %v.\n", num)
				continue
			}
			if num > randomInt {
				fmt.Printf("Incorrect! The number is less than %v.\n", num)
				continue
			}
		}
	}
	fmt.Println("\nGAME OVER! You have run out of attempts.")
	endMessage()
}

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")
	gameCycle()
}
