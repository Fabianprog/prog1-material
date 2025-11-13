package chess

import "fmt"

func ExampleChessPiece_MoveAllowed() {
	l1 := ChessPiece{
		pieceType: BISHOP,
		colour:    WHITE,
		row:       2,
		column:    2,
	}
	l2 := ChessPiece{
		pieceType: KNIGHT,
		colour:    WHITE,
		row:       2,
		column:    2,
	}
	l3 := ChessPiece{
		pieceType: QUEEN,
		colour:    WHITE,
		row:       2,
		column:    2,
	}
	l4 := ChessPiece{
		pieceType: KING,
		colour:    WHITE,
		row:       2,
		column:    2,
	}
	l5 := ChessPiece{
		pieceType: ROOK,
		colour:    WHITE,
		row:       2,
		column:    2,
	}

	fmt.Println(l1.MoveAllowed(3, 4))
	fmt.Println(l1.MoveAllowed(0, 0))
	fmt.Println(l1.MoveAllowed(10, 10))
	fmt.Println(l2.MoveAllowed(3, 4))
	fmt.Println(l2.MoveAllowed(0, 0))
	fmt.Println(l2.MoveAllowed(10, 10))
	fmt.Println(l3.MoveAllowed(3, 4))
	fmt.Println(l3.MoveAllowed(0, 0))
	fmt.Println(l3.MoveAllowed(9, 9))
	fmt.Println(l4.MoveAllowed(3, 4))
	fmt.Println(l4.MoveAllowed(0, 0))
	fmt.Println(l4.MoveAllowed(11, 11))
	fmt.Println(l5.MoveAllowed(3, 4))
	fmt.Println(l5.MoveAllowed(0, 0))
	fmt.Println(l5.MoveAllowed(12, 12))

	// Output:
	// false
	// true
	// false
	//
	//
	// false
	//
	//
	// false
	//
	//
	// false
	//
	//
	// false
}
