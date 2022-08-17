package pieces

import (
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/hajimehoshi/ebiten/v2"
)

type Bishop struct {
	CurrPos helpers.Coord
	Color   string
	Sprite  *ebiten.Image
}

func InitBishop(pos helpers.Coord, color string) *Bishop {
	sprite := generateSprite(color, cfg.SpriteMap["bishop"])
	return &Bishop{pos, color, sprite}
}

func (b Bishop) Draw(dst *ebiten.Image) {
	op := generateDrawingOps(b.CurrPos)
	dst.DrawImage(b.Sprite, op)
}

func (b Bishop) GenValidMoves() [][]helpers.Coord {

	x := int(b.CurrPos.X)
	y := int(b.CurrPos.Y)

	return filterNegatives(GenDiagonalMoves(x, y))

}

func (b Bishop) GetPos() *helpers.Coord {
	return &b.CurrPos
}

func (b Bishop) GetColor() string {
	return b.Color
}

func (b *Bishop) Move(pos helpers.Coord) {
	b.CurrPos = pos
}
