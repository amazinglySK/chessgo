package pieces

import (
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/hajimehoshi/ebiten/v2"
)

type Queen struct {
	CurrPos helpers.Coord
	Color   string
	Sprite  *ebiten.Image
}

func InitQueen(pos helpers.Coord, color string) *Queen {
	sprite := generateSprite(color, cfg.SpriteMap["queen"])
	return &Queen{pos, color, sprite}

}

func (b Queen) Draw(dst *ebiten.Image) {
	op := generateDrawingOps(b.CurrPos)
	dst.DrawImage(b.Sprite, op)
}

func (b Queen) GenValidMoves() [][]helpers.Coord {
	diagonals := GenDiagonalMoves(int(b.CurrPos.X), int(b.CurrPos.Y))
	straights := GenStraightMoves(b.CurrPos.X, b.CurrPos.Y)
	moves := [][]helpers.Coord{}

	for _, i := range diagonals {
		moves = append(moves, i)
	}
	for _, i := range straights {
		moves = append(moves, i)
	}

	return filterNegatives(moves)
}

func (b Queen) GetPos() *helpers.Coord {
	return &b.CurrPos
}

func (b Queen) GetColor() string {
	return b.Color
}

func (b *Queen) Move(pos helpers.Coord) {
	b.CurrPos = pos
}
