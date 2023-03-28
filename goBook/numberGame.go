package goBook

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func NumGame() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chose a random number between 1 and 100")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)
	success := false

	for i := 0; i < 10; i++ {
		fmt.Println("You haz: ", 10-i, "guesses left.")
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess > target {
			fmt.Println("Too high you bum!")
		} else if guess < target {
			fmt.Println("Too low you bum!")
		} else {
			fmt.Println("Good job you dip! You guessed it!")
			success = true
			break
		}
	}

	if !success {
		fmt.Println("You dip you didn't guess the magic number: ", target)
	}

}
