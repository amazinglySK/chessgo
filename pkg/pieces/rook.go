package pieces

import (
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Rook struct {
	CurrPos helpers.Coord
	Color   string
	Sprite  *ebiten.Image
}

func (b Rook) Draw(dst *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(5, 5)
	op.GeoM.Translate(float64(cfg.BoardPadding)+b.CurrPos.X*float64(cfg.SquareSize), float64(cfg.BoardPadding)+b.CurrPos.Y*float64(cfg.SquareSize)-8)

	dst.DrawImage(b.Sprite, op)
}


func (b Rook) GenValidMoves() [][]helpers.Coord {
	return filterNegatives(GenStraightMoves(b.CurrPos.X, b.CurrPos.Y))
}

func InitRook(pos helpers.Coord, color string) *Rook {
	var sprite *ebiten.Image
	switch color {
	case "white":
		sprite = WhiteSprites.SubImage(image.Rect(16, 0, 32, 16)).(*ebiten.Image)
	case "black":
		sprite = BlackSprites.SubImage(image.Rect(16, 0, 32, 16)).(*ebiten.Image)
	}
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
