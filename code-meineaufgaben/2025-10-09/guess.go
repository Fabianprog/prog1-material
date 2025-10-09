package main

import (
	"fmt"
	"math/rand"
)

func main() {
	my_number := rand.Intn(100) + 1
	fmt.Println(my_number)
	for i := 0; i < 3; i++ {
		guess := ReadNumber()
		if my_number < guess {
			fmt.Println("Die gesuchte Zahl ist kleiner")
		}
		if my_number > guess {
			fmt.Println("Die gesuchte zahl ist größer")
		}
		if my_number == guess {
			fmt.Print("Das war richtig!")
			return
		}

	}
	fmt.Print("Game Over! ")
}

func ReadNumber() int {
	var n int

	fmt.Print("Bitte gib eine zahl ein: ")
	fmt.Scan(&n)

	return n

}
