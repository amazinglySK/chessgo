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
