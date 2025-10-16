package fizz

import (
	"fmt"
)

func Fizzcompact() {

	for i := 0; i <= 30; i++ {

		if i%5 == 0 && i%3 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}

	}
}
