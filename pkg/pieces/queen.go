package pieces

import (
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Queen struct {
	CurrPos helpers.Coord
	Color   string
	Sprite  *ebiten.Image
}

func InitQueen(pos helpers.Coord, color string) Queen {
	var sprite *ebiten.Image
	switch color {
	case "white":
		sprite = WhiteSprites.SubImage(image.Rect(0, 32, 16, 48)).(*ebiten.Image)
	case "black":
		sprite = BlackSprites.SubImage(image.Rect(0, 32, 16, 48)).(*ebiten.Image)
	}

	return Queen{pos, color, sprite}

}

func (b Queen) Draw(dst *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(5, 5)
	op.GeoM.Translate(float64(cfg.BoardPadding)+b.CurrPos.X*float64(cfg.SquareSize), float64(cfg.BoardPadding)+b.CurrPos.Y*float64(cfg.SquareSize)-4)

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
