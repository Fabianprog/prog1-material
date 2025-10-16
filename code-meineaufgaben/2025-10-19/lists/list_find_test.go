package lists

import "fmt"

func Find(l []int, a int) int {
	for i := 0; i < len(l); i++ {
		if l[i] == a {
			return i

		}

	}
	return -1
}

func ExampleFind() {
	l1 := []int{17, 5, 42, 25, 3, 4, 8, -23}

	pos1 := Find(l1, 42)
	pos2 := Find(l1, 100)
	fmt.Println(pos1)
	fmt.Println(pos2)

	//Output:
	//2
	//-1
}
