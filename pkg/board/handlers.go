package board

import (
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/amazinglySK/chessgo/pkg/pieces"
	"github.com/amazinglySK/chessgo/pkg/player"
	"github.com/amazinglySK/chessgo/pkg/square"
	"log"
)

// HandleSelect handles the selecting of a piece
func (b *Board) HandleSelect(sq *square.Square, curr_player *player.Player) {
	piece := sq.Piece
	curr_player.CurrPiece = piece
	coords := piece.GenValidMoves()
	var moves []helpers.Coord
	switch piece.(type) {
	case *pieces.Pawn:
		moves = curr_player.FilterPawnMoves(piece, coords, b.Squares)
		break
	default:
		moves = curr_player.FilterMoves(coords, b.Squares)
	}

	// Acitvating the current piece square
	sq.Activate()
	b.PrevActives = append(b.PrevActives, sq)

	curr_player.ValidMoves = moves

	// Activating all move squares
	for _, move := range moves {
		b.GetSquare(move).Activate()
		sq = b.GetSquare(move)
		b.PrevActives = append(b.PrevActives, sq)
	}
}

// HandleMove handles the moving of a piece to a position
func (b *Board) HandleMove(curr_player *player.Player, pos helpers.Coord) bool {
	var king_cap = false
	prev_sq := b.GetSquare(*curr_player.CurrPiece.GetPos())
	move_sq := b.GetSquare(pos)
	switch move_sq.Piece.(type) {
	// Captured the king
	case *pieces.King:
		king_cap = true

	}

	// Valid move
	if curr_player.MovePiece(curr_player.CurrPiece, move_sq, prev_sq) {
		curr_player.CurrPiece = nil
		// Swapping the current player
		b.CurrPlayerIdx = uint(len(b.Players)) - (b.CurrPlayerIdx + 1)

	} else {
		log.Println("invalid move")
	}

	return king_cap
}
