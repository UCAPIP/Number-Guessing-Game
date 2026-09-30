package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var difficult int

func scanDifficult() {
	fmt.Print("\nEnter your choice: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		text := scanner.Text()
		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Reading error:", err)
			return
		}
		if num > 3 || num < 1 {
			fmt.Println("Enter num of difficult")
			scanDifficult()
		}
		difficult = num
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Reading error:", err)
	}

	switch difficult {
	case 1:
		fmt.Println("\nGreat! You have selected the Easy difficulty level.")
	case 2:
		fmt.Println("\nGreat! You have selected the Medium difficulty level.")
	case 3:
		fmt.Println("\nGreat! You have selected the Hard difficulty level.")
	}

}

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")
	fmt.Println("You have 5 chances to guess the correct number.")

	fmt.Println("\n" + `Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)`)

	scanDifficult()
	fmt.Println("Let's start the game!")

	for {

	}
}
