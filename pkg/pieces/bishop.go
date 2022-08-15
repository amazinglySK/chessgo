package pieces

import (
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Bishop struct {
	CurrPos helpers.Coord
	Color   string
	Sprite  *ebiten.Image
}

func InitBishop(pos helpers.Coord, color string) Bishop {
	var sprite *ebiten.Image
	switch color {
	case "white":
		sprite = WhiteSprites.SubImage(image.Rect(16, 16, 32, 32)).(*ebiten.Image)
	case "black":
		sprite = BlackSprites.SubImage(image.Rect(16, 16, 32, 32)).(*ebiten.Image)
	}

	return Bishop{pos, color, sprite}

}

func (b Bishop) Draw(dst *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(5, 5)
	op.GeoM.Translate(float64(cfg.BoardPadding)+b.CurrPos.X*float64(cfg.SquareSize), float64(cfg.BoardPadding)+b.CurrPos.Y*float64(cfg.SquareSize)-8)

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
