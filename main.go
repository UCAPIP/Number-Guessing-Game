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

func scanDifficult() int {

	difficults := map[int]int{
		1: 10, // Easy
		2: 5,  // Medium
		3: 3,  // Hard
	}

	var timesToTry int

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
			return 0
		}
		if num > 3 || num < 1 {
			fmt.Println("Enter num of difficult")
			scanDifficult()
		}
		timesToTry = difficults[num]
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Reading error:", err)
		return 0
	}

	switch timesToTry {
	case 10:
		fmt.Println("\nGreat! You have selected the Easy difficulty level.")
		return 10
	case 5:
		fmt.Println("\nGreat! You have selected the Medium difficulty level.")
		return 5
	case 3:
		fmt.Println("\nGreat! You have selected the Hard difficulty level.")
		return 3
	}
	return 0
}

func playAgain() bool {
	for {
		fmt.Println("\nIf you want to play again - type Y or N for exit")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			text := strings.ToUpper(scanner.Text())
			switch text {
			case "Y":
				return true
			case "N":
				return false
			default:
				fmt.Println("UNKNOWN COMMAND")
			}
		}
	}
}

func gameCycle() {
	maxAttempts := scanDifficult()
	randomInt := rand.IntN(100) + 1
	fmt.Println("Let's start the game!")
	start := time.Now()
	for i := 1; i <= maxAttempts; i++ {
		//fmt.Println("Answer: ", randomInt) -- Ответ
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
				return
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
}

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")
	for {
		gameCycle()
		if !playAgain() {
			return
		}
	}
}
