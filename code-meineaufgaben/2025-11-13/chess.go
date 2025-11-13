package chess

type PType int

const (
	BISHOP PType = iota
	KNIGHT
	QUEEN
	KING
	ROOK
	PAWN
)

type Field int

const (
	Outofcol = 9
	Outofrow = 9
)

type Colour int

const (
	WHITE Colour = iota
	BLACK
)

// ChessPiece repräsentiert eine Schachfigur auf einem Spielfeld.
type ChessPiece struct {
	pieceType PType
	colour    Colour
	row       int
	column    int
}

// Methoden: TODO

// MoveAllowed erwartet eine Feld-Angabe und liefert true,
// falls die Figur nach den Bewegungsregeln beim Schach
// auf dieses Feld ziehen darf.
// Besonderheiten wie Rochade oder im Weg stehende Figuren
// Spielen keine Rolle.
func (p ChessPiece) MoveAllowed(row, col int) bool {

	// switch p.pieceType {
	// case BISHOP: ...
	// case KNIGHT: ...
	// }

	if p.pieceType == BISHOP {
		return p.MoveAllowedBishop(row, col)
	}
	if p.pieceType == KNIGHT {
		return p.MoveAllowedKnight(row, col)
	}
	if p.pieceType == QUEEN {
		return p.MoveAllowedQueen(row, col)
	}
	if p.pieceType == ROOK {
		return p.MoveAllowedRook(row, col)
	}
	if p.pieceType == KING {
		return p.MoveAllowedKing(row, col)
	}

	return false
}

func (p ChessPiece) MoveAllowedBishop(row, col int) bool {
	if row >= Outofrow || col >= Outofcol {
		return false
	}
	// Diagonale von links unten nach rechts oben.
	if row-col == p.row-p.column {
		return true
	}
	// Diagonale von links oben nach rechts unten.
	if row+col == p.row+p.column {
		return true
	}
	return false
}
func (p ChessPiece) MoveAllowedQueen(row, col int) bool {
	if row >= Outofrow || col >= Outofcol {
		return false
	}
	return p.MoveAllowedBishop(row, col) || p.MoveAllowedRook(row, col)
}
func (p ChessPiece) MoveAllowedKnight(row, col int) bool {
	if row >= Outofrow || col >= Outofcol {
		return false
	}
	if (row == p.row+1 || row == p.row-1) && (col == p.column-2 || col == p.column+2) {
		return true
	}
	if (row == p.row+2 || row == p.row-2) && (col == p.column-1 || col == p.column+1) {
		return true
	}
	return false

}
func (p ChessPiece) MoveAllowedRook(row, col int) bool {
	if row >= Outofrow || col >= Outofcol {
		return false
	}
	if row == p.row {
		return true
	}
	if col == p.column {
		return true
	}
	return false
}
func (p ChessPiece) MoveAllowedKing(row, col int) bool {
	if row >= Outofrow || col >= Outofcol {
		return false
	}
	if row == p.row && (col == p.column-1 || col == p.column+1) {
		return true
	}
	if col == p.column && (row == p.row-1 || row == p.row+1) {
		return true
	}
	if (col == p.column-1 || col == p.column+1) && (row == p.row-1 || row == p.row+1) {
		return true
	}
	return false
}
