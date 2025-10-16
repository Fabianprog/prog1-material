package fizz

import (
	"fmt"
)

func FizzImproved(x, y, n int) {

	for i := 0; i <= 30; i++ {

		if i%x == 0 && i%y == 0 {
			fmt.Println("fizzbuzz")
		} else if i%x == 0 {
			fmt.Println("fizz")
		} else if i%y == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}

	}
}
