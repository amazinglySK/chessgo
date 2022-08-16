package pieces

import (
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type King struct {
	CurrPos helpers.Coord
	Color   string
	Sprite  *ebiten.Image
}

func InitKing(pos helpers.Coord, color string) *King {
	var sprite *ebiten.Image
	switch color {
	case "white":
		sprite = WhiteSprites.SubImage(image.Rect(16, 32, 32, 48)).(*ebiten.Image)
	case "black":
		sprite = BlackSprites.SubImage(image.Rect(16, 32, 32, 48)).(*ebiten.Image)
	}

	return &King{pos, color, sprite}

}

func (b King) Draw(dst *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(5, 5)
	op.GeoM.Translate(float64(cfg.BoardPadding)+b.CurrPos.X*float64(cfg.SquareSize), float64(cfg.BoardPadding)+b.CurrPos.Y*float64(cfg.SquareSize)-4)

	dst.DrawImage(b.Sprite, op)
}

func (b King) GenValidMoves() [][]helpers.Coord {
	x, y := b.CurrPos.X, b.CurrPos.Y
	moves := [][]helpers.Coord{{helpers.Coord{x + 1, y}}, {helpers.Coord{x - 1, y}}, {helpers.Coord{x, y + 1}}, {helpers.Coord{x, y - 1}}}
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
