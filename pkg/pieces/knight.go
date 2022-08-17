package pieces

import (
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/hajimehoshi/ebiten/v2"
)

type Knight struct {
	CurrPos helpers.Coord
	Color   string
	Sprite  *ebiten.Image
}

func (b Knight) Draw(dst *ebiten.Image) {
	op := generateDrawingOps(b.CurrPos)
	dst.DrawImage(b.Sprite, op)
}

func (b Knight) GenValidMoves() [][]helpers.Coord {
	moves := [][]helpers.Coord{}
	x, y := b.CurrPos.X, b.CurrPos.Y
	coord := [][]float64{{x - 2, y + 1}, {x - 2, y - 1}, {x + 2, y + 1}, {x + 2, y - 1}, {x - 1, y + 2}, {x + 1, y + 2}, {x - 1, y - 2}, {x + 1, y - 2}}

	for _, c := range coord {
		x, y := c[0], c[1]
		set := []helpers.Coord{helpers.Coord{x, y}}
		moves = append(moves, set)
	}

	return filterNegatives(moves)

}

func InitKnight(pos helpers.Coord, color string) *Knight {
	sprite := generateSprite(color, cfg.SpriteMap["knight"])
	return &Knight{pos, color, sprite}
}

func (b Knight) GetPos() *helpers.Coord {
	return &b.CurrPos
}
func (b Knight) GetColor() string {
	return b.Color
}

func (b *Knight) Move(pos helpers.Coord) {
	b.CurrPos = pos
}
