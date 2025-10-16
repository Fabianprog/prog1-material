package fizz

import (
	"fmt"
)

func Fizzcompact() {

	for i := 0; i <= 30; i++ {
		if i%3 != 0 && i%5 != 0 {
			fmt.Println(i)
			continue
		}

		if i%3 == 0 {
			fmt.Print("fizz")
		}
		if i%5 == 0 {
			fmt.Print("buzz")
		}

	}
}
