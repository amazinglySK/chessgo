package pieces

import (
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/hajimehoshi/ebiten/v2"
)

type Rook struct {
	CurrPos helpers.Coord
	Color   string
	Sprite  *ebiten.Image
}

func (b Rook) Draw(dst *ebiten.Image) {
	op := generateDrawingOps(b.CurrPos)
	dst.DrawImage(b.Sprite, op)
}

func (b Rook) GenValidMoves() [][]helpers.Coord {
	return filterNegatives(GenStraightMoves(b.CurrPos.X, b.CurrPos.Y))
}

func InitRook(pos helpers.Coord, color string) *Rook {
	sprite := generateSprite(color, cfg.SpriteMap["rook"])
	return &Rook{pos, color, sprite}
}

func (b Rook) GetPos() *helpers.Coord {
	return &b.CurrPos
}

func (b Rook) GetColor() string {
	return b.Color
}

func (b *Rook) Move(pos helpers.Coord) {
	b.CurrPos = pos
}
