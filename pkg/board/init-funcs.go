package board

import (
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/amazinglySK/chessgo/pkg/square"
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/amazinglySK/chessgo/pkg/pieces"
	"image/color"

)

// makeSquare makes squares assigned with the right coordinates when given a width and height
func makeSquares(w, h int) [][]*square.Square {
	white := true
	squares := [][]*square.Square{}
	for i := 0; i < h; i++ {
		row := []*square.Square{}
		for j := 0; j < w; j++ {
			var sq_color color.Color
			var sq_color_str string
			if white {
				sq_color = cfg.LightColor
				sq_color_str = "white"
			} else {
				sq_color = cfg.DarkColor
				sq_color_str = "black"
			}

			white = !white
			coords := helpers.Coord{X: float64(j), Y: float64(i)}
			row = append(row, &square.Square{Pos: coords, Size: float64(cfg.SquareSize), Clr: sq_color, Color: sq_color_str, Occupied: false, Orig: sq_color})
		}
		squares = append(squares, row)
		white = !white
	}
	return squares
}

// genWhitePieces initiates all the white pieces with their starting positions
func genWhitePieces(w int) []pieces.Piece {
	pieces_slice := []pieces.Piece{}
	pieces_slice = append(pieces_slice, pieces.InitBishop(helpers.Coord{X: 2, Y: 0}, "white"), pieces.InitBishop(helpers.Coord{X: float64(w - 3), Y: 0}, "white"))
	pieces_slice = append(pieces_slice, pieces.InitRook(helpers.Coord{X: 0, Y: 0}, "white"), pieces.InitRook(helpers.Coord{X: float64(w - 1), Y: 0}, "white"))
	pieces_slice = append(pieces_slice, pieces.InitKnight(helpers.Coord{X: 1, Y: 0}, "white"), pieces.InitKnight(helpers.Coord{X: float64(w - 2), Y: 0}, "white"))
	pieces_slice = append(pieces_slice, pieces.InitQueen(helpers.Coord{X: 4, Y: 0}, "white"))
	pieces_slice = append(pieces_slice, pieces.InitKing(helpers.Coord{X: 3, Y: 0}, "white"))
	for _, pawn := range pieces.InitPawns(1, w, "white") {
		pieces_slice = append(pieces_slice, pawn)
	}
	return pieces_slice
}

// genBlackPieces initiate all the black pieces with their starting positions
func genBlackPieces(w int) []pieces.Piece {
	pieces_slice := []pieces.Piece{}
	pieces_slice = append(pieces_slice, pieces.InitBishop(helpers.Coord{X: 2, Y: float64(w - 1)}, "black"), pieces.InitBishop(helpers.Coord{X: float64(w - 3), Y: float64(w - 1)}, "black"))
	pieces_slice = append(pieces_slice, pieces.InitRook(helpers.Coord{X: 0, Y: float64(w - 1)}, "black"), pieces.InitRook(helpers.Coord{X: float64(w - 1), Y: float64(w - 1)}, "black"))
	pieces_slice = append(pieces_slice, pieces.InitKnight(helpers.Coord{X: 1, Y: float64(w - 1)}, "black"), pieces.InitKnight(helpers.Coord{X: float64(w - 2), Y: float64(w - 1)}, "black"))
	pieces_slice = append(pieces_slice, pieces.InitQueen(helpers.Coord{X: 4, Y: float64(w - 1)}, "black"))
	pieces_slice = append(pieces_slice, pieces.InitKing(helpers.Coord{X: 3, Y: float64(w - 1)}, "black"))
	for _, pawn := range pieces.InitPawns(w-2, w, "black") {
		pieces_slice = append(pieces_slice, pawn)
	}

	return pieces_slice
}

// UpdateInitialSquareState updates the occupied status and the piece on it
func (b *Board) UpdateInitialSquareState(pieces []pieces.Piece) {
	for _, piece := range pieces {
		pos := *piece.GetPos()
		sq := b.GetSquare(pos)
		sq.Piece = piece
		sq.Occupied = true
	}
}
