package pieces

import (
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type Pawn struct {
	CurrPos   helpers.Coord
	Color     string
	Promotion helpers.Coord
	Sprite    *ebiten.Image
}

func InitPawns(y int, w int, color string) []Pawn {
	pawns := []Pawn{}
	for i := 0; i < w; i++ {
		var (
			sprite    *ebiten.Image
			promotion helpers.Coord
		)
		switch color {
		case "white":
			sprite = WhiteSprites.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image)
			promotion = helpers.Coord{float64(i), 8}
		case "black":
			sprite = BlackSprites.SubImage(image.Rect(0, 0, 16, 16)).(*ebiten.Image)
			promotion = helpers.Coord{float64(i), 0}
		}

		pos := helpers.Coord{float64(i), float64(y)}

		pawns = append(pawns, Pawn{pos, color, promotion, sprite})

	}
	return pawns

}

func (b Pawn) Draw(dst *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(5, 5)
	op.GeoM.Translate(float64(cfg.BoardPadding)+b.CurrPos.X*float64(cfg.SquareSize), float64(cfg.BoardPadding)+b.CurrPos.Y*float64(cfg.SquareSize)-8)

	dst.DrawImage(b.Sprite, op)
}

func (b Pawn) GenValidMoves() [][]helpers.Coord {
	x, y := b.CurrPos.X, b.CurrPos.Y
	var moves [][]helpers.Coord
	if b.Color == "white" {
		moves = [][]helpers.Coord{{helpers.Coord{x + 1, y + 1}}, {helpers.Coord{x - 1, y + 1}}, {helpers.Coord{x, y + 1}}}
	}else {
		moves = [][]helpers.Coord{{helpers.Coord{x + 1, y - 1}}, {helpers.Coord{x - 1, y - 1}}, {helpers.Coord{x, y - 1}}}
	}

	return filterNegatives(moves)
}

func (b Pawn) GetPos() *helpers.Coord {
	return &b.CurrPos
}

func (b Pawn) GetColor() string {
	return b.Color
}
