package matrices

import "fmt"

func PrintMatrix(m [][]int) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			fmt.Print(m[i][j])
			if j < len(m[i])-1 {
				fmt.Print(" ")
			}

		}
		fmt.Println()

	}
}
