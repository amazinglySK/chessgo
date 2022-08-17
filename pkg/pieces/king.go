package pieces

import (
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/hajimehoshi/ebiten/v2"
)

type King struct {
	CurrPos helpers.Coord
	Color   string
	Sprite  *ebiten.Image
}

func InitKing(pos helpers.Coord, color string) *King {
	sprite := generateSprite(color, cfg.SpriteMap["king"])
	return &King{pos, color, sprite}

}

func (b King) Draw(dst *ebiten.Image) {
	op := generateDrawingOps(b.CurrPos)
	dst.DrawImage(b.Sprite, op)
}

func (b King) GenValidMoves() [][]helpers.Coord {
	x, y := b.CurrPos.X, b.CurrPos.Y
	moves := [][]helpers.Coord{{helpers.Coord{x + 1, y}}, {helpers.Coord{x - 1, y}}, {helpers.Coord{x, y + 1}}, {helpers.Coord{x, y - 1}}, {helpers.Coord{x - 1, y - 1}}, {helpers.Coord{x + 1, y + 1}}, {helpers.Coord{x - 1, y + 1}}, {helpers.Coord{x + 1, y - 1}}}
	return filterNegatives(moves)
}

func (b King) GetPos() *helpers.Coord {
	return &b.CurrPos
}

func (b King) GetColor() string {
	return b.Color
}

func (b *King) Move(pos helpers.Coord) {
	b.CurrPos = pos
}
