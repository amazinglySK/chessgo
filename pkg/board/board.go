package board

import (
	"image/color"

	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/amazinglySK/chessgo/pkg/events"
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/amazinglySK/chessgo/pkg/pieces"
	"github.com/amazinglySK/chessgo/pkg/player"
	"github.com/amazinglySK/chessgo/pkg/square"
	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	Squares       [][]*square.Square
	CurrPlayerIdx uint
	Players       []*player.Player
	PrevActive    *square.Square
}

func makeSquares(w, h int) [][]*square.Square {
	white := false
	squares := [][]*square.Square{}
	for i := 0; i < h; i++ {
		row := []*square.Square{}
		for j := 0; j < w; j++ {
			var sq_color color.Color
			if white {
				sq_color = cfg.LightColor
			} else {
				sq_color = cfg.DarkColor
			}

			white = !white
			coords := helpers.Coord{X: float64(j), Y: float64(i)}
			row = append(row, &square.Square{Pos: coords, Size: float64(cfg.SquareSize), Clr: sq_color, Occupied: false, Orig: sq_color})
		}
		squares = append(squares, row)
		white = !white
	}
	return squares
}

func genWhitePieces(w int) []pieces.Piece {
	pieces_slice := []pieces.Piece{}
	pieces_slice = append(pieces_slice, pieces.InitBishop(helpers.Coord{X: 2, Y: 0}, "white"), pieces.InitBishop(helpers.Coord{X: float64(w - 3), Y: 0}, "white"))
	pieces_slice = append(pieces_slice, pieces.InitRook(helpers.Coord{X: 0, Y: 0}, "white"), pieces.InitRook(helpers.Coord{X: float64(w - 1), Y: 0}, "white"))
	pieces_slice = append(pieces_slice, pieces.InitKnight(helpers.Coord{X: 1, Y: 0}, "white"), pieces.InitKnight(helpers.Coord{X: float64(w - 2), Y: 0}, "white"))
	pieces_slice = append(pieces_slice, pieces.InitQueen(helpers.Coord{X: 3, Y: 0}, "white"))
	pieces_slice = append(pieces_slice, pieces.InitKing(helpers.Coord{X: 4, Y: 0}, "white"))
	for _, pawn := range pieces.InitPawns(1, w, "white") {
		pieces_slice = append(pieces_slice, pawn)
	}
	return pieces_slice
}

func genBlackPieces(w int) []pieces.Piece {
	pieces_slice := []pieces.Piece{}
	pieces_slice = append(pieces_slice, pieces.InitBishop(helpers.Coord{X: 2, Y: float64(w - 1)}, "black"), pieces.InitBishop(helpers.Coord{X: float64(w - 3), Y: float64(w - 1)}, "black"))
	pieces_slice = append(pieces_slice, pieces.InitRook(helpers.Coord{X: 0, Y: float64(w - 1)}, "black"), pieces.InitRook(helpers.Coord{X: float64(w - 1), Y: float64(w - 1)}, "black"))
	pieces_slice = append(pieces_slice, pieces.InitKnight(helpers.Coord{X: 1, Y: float64(w - 1)}, "black"), pieces.InitKnight(helpers.Coord{X: float64(w - 2), Y: float64(w - 1)}, "black"))
	pieces_slice = append(pieces_slice, pieces.InitQueen(helpers.Coord{X: 3, Y: float64(w - 1)}, "black"))
	pieces_slice = append(pieces_slice, pieces.InitKing(helpers.Coord{X: 4, Y: float64(w - 1)}, "black"))
	for _, pawn := range pieces.InitPawns(w-2, w, "black") {
		pieces_slice = append(pieces_slice, pawn)
	}

	return pieces_slice
}

func (b *Board) UpdateInitialSquareState() {
	for _, player := range b.Players {
		for _, piece := range player.Pieces {
			pos := *piece.GetPos()
			sq := b.GetSquare(pos)
			sq.Piece = piece
			sq.Occupied = true
		}
	}
}

func InitBoard(w int, h int) Board {
	squares := makeSquares(w, h)
	whitePlayer := player.Player{Color: "white", Pieces: genWhitePieces(w)}
	blackPlayer := player.Player{Color: "black", Pieces: genBlackPieces(w)}
	board := Board{Squares: squares, CurrPlayerIdx: 0, Players: []*player.Player{&whitePlayer, &blackPlayer}}
	board.UpdateInitialSquareState()
	return board
}

func (b *Board) GetSquare(pos helpers.Coord) *square.Square {
	x, y := int(pos.X), int(pos.Y)
	return b.Squares[y][x]
}

func (b *Board) Draw(dst *ebiten.Image) {
	for _, row := range b.Squares {
		for _, s := range row {
			s.Draw(dst)
		}
	}
	for _, player := range b.Players {
		for _, piece := range player.Pieces {
			piece.Draw(dst)
		}
	}
}

func (b *Board) ManageClick() {
	clicked, pos := events.CheckMouseEvents()
	if clicked {
		square := b.GetSquare(pos)
		if square.Occupied {
			// TODO : Generate Valid moves for that piece
			piece := square.Piece
			coords := piece.GenValidMoves()
			moves := b.Players[b.CurrPlayerIdx].FilterMoves(coords, b.Squares)

			for _, move := range moves {
				b.GetSquare(move).Activate(b.PrevActive)
			} 

			square.Activate(b.PrevActive)
			b.PrevActive = square
		}
	}
}

func (b *Board) Update() {
	b.ManageClick()
}
