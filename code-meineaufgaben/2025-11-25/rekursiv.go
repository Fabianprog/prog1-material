package rekursiv

import "fmt"

func CountDown(n int) {
	fmt.Println(n)
	if n <= 0 {
		return
	}
	CountDown(n - 1)

}

func CountUp(n int) {
	if n <= 101 {
		return
	}
	CountDown(n + 1)
	fmt.Println(n)
}
