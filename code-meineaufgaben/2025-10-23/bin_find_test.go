package searching

import "fmt"

func ExampleBinFindRek() {
	l1 := []int{1, 6, 8, 10, 25, 42, 125, 277}

	fmt.Println(BinFindRek(l1, 1))
	fmt.Println(BinFindRek(l1, 6))
	fmt.Println(BinFindRek(l1, 25))
	fmt.Println(BinFindRek(l1, 125))
	fmt.Println(BinFindRek(l1, 277))
	fmt.Println(BinFindRek(l1, 3))

	//Output:
	//0
	//1
	//4
	//6
	//7
	//-1
}
