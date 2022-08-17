package player

import (
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/amazinglySK/chessgo/pkg/pieces"
	"github.com/amazinglySK/chessgo/pkg/square"
)

// Player is an object store the player color white/black and a set of valid moves of the current piece in focus and all the piece
type Player struct {
	Pieces     []pieces.Piece
	Color      string
	ValidMoves []helpers.Coord
	CurrPiece  pieces.Piece
}

// FilterMoves filters moves for any normal moving pieces without any odds
func (p Player) FilterMoves(moves [][]helpers.Coord, squares [][]*square.Square) []helpers.Coord {
	valid_moves := []helpers.Coord{}
	for _, set := range moves {
		for _, move := range set {
			x, y := int(move.X), int(move.Y)
			sq := squares[y][x]
			if sq.Occupied && sq.Piece.GetColor() == p.Color {
				break
			}
			valid_moves = append(valid_moves, move)
			if sq.Occupied && sq.Piece.GetColor() != p.Color {
				break
			}
		}
	}

	return valid_moves
}

// FilterPawnMoves is filtering for pawns so that they don't end up always showing a diagonal move available
func (p Player) FilterPawnMoves(piece pieces.Piece, moves [][]helpers.Coord, squares [][]*square.Square) []helpers.Coord {
	valid := []helpers.Coord{}
	for _, set := range moves {
		for _, move := range set {
			x, y := int(move.X), int(move.Y)
			curr_x, curr_y := int(piece.GetPos().X), int(piece.GetPos().Y)
			sq := squares[y][x]
			curr_sq := squares[curr_y][curr_x]

			// This means it's a pawns diagonal move
			if sq.Color == curr_sq.Color {
				// There's a piece which is of the opponent
				if sq.Occupied && sq.Piece.GetColor() != piece.GetColor() {
					valid = append(valid, move)
				}
			} else {
				// An allied piece
				if sq.Occupied {
					break
				}
				valid = append(valid, move)
			}
		}
	}

	return valid
}



// MovePiece moves a given piece to a valid square. If the move's a valid move then returns true else returns false.
func (p Player) MovePiece(piece pieces.Piece, sq *square.Square, prev_sq *square.Square) bool {
	for _, c := range p.ValidMoves {
		if c == sq.Pos {
			prev_sq.Piece = nil
			prev_sq.Occupied = false
			sq.Piece = piece
			sq.Occupied = true
			piece.Move(sq.Pos)
			return true
		}
	}
	return false
}
