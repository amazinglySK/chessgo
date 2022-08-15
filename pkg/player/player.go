package player

import (
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/amazinglySK/chessgo/pkg/pieces"
	"github.com/amazinglySK/chessgo/pkg/square"
)

type Player struct {
	Pieces []pieces.Piece
	Color  string
}

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
		}
	}

	return valid_moves
}

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
				if sq.Occupied && sq.Piece.GetColor() == piece.GetColor() {
					break
				}
				valid = append(valid, move)
			}
		}
	}

	return valid
}
