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

func FindAll(l []int, a int) []int {
	result := []int{}

	for i := 0; i < len(l); i++ {
		if l[i] == a {
			result = append(result, i)
		}
	}

	return result

}

func ExampleFind() {
	l1 := []int{17, 5, 42, 25, 3, 4, 8, -23, 5}

	pos1 := FindAll(l1, 42)
	pos2 := FindAll(l1, 100)
	pos3 := FindAll(l1, 5)
	fmt.Println(pos1)
	fmt.Println(pos2)
	fmt.Println(pos3)

	//Output:
	//[2]
	//[]
	//[1 8]
}
