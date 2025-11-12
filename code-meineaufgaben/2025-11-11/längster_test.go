package duration

import "fmt"

func ExampleLängster() {
	a := []string{"hallo", "hall", "hal", "ha"}
	fmt.Println(Längster(a))
	//Output:
	//4
}
