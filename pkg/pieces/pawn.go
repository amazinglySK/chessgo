package pieces

import (
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/hajimehoshi/ebiten/v2"
)

type Pawn struct {
	CurrPos helpers.Coord
	Color   string
	Sprite  *ebiten.Image
}

func InitPawns(y int, w int, color string) []*Pawn {
	pawns := []*Pawn{}
	for i := 0; i < w; i++ {
		sprite := generateSprite(color, cfg.SpriteMap["pawn"])
		pos := helpers.Coord{float64(i), float64(y)}
		pawns = append(pawns, &Pawn{pos, color, sprite})
	}
	return pawns
}

func (b Pawn) Draw(dst *ebiten.Image) {
	op := generateDrawingOps(b.CurrPos)
	dst.DrawImage(b.Sprite, op)
}

func (b Pawn) GenValidMoves() [][]helpers.Coord {
	x, y := b.CurrPos.X, b.CurrPos.Y
	var moves [][]helpers.Coord
	if b.Color == "white" {
		moves = [][]helpers.Coord{{helpers.Coord{x + 1, y + 1}}, {helpers.Coord{x - 1, y + 1}}, {helpers.Coord{x, y + 1}}}
	} else {
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

func (b *Pawn) Move(pos helpers.Coord) {
	b.CurrPos = pos
}
