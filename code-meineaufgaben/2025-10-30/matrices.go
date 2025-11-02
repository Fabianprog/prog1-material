package matrices

import "fmt"

func PrintMatrix(m [][]int) {
	fmt.Println("+---+---+---+")
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Print("|")
			fmt.Printf(" %v ", m[i][j])

		}
		fmt.Println("|")
		fmt.Println("+---+---+---+")

	}
}
