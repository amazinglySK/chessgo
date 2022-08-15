package square

import (
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/amazinglySK/chessgo/pkg/pieces"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

type Square struct {
	Pos      helpers.Coord
	Clr      color.Color
	Size     float64
	Occupied bool
	Orig     color.Color
	Piece    pieces.Piece
}

func (sq Square) Draw(dst *ebiten.Image) {
	x := sq.Pos.X*sq.Size + float64(cfg.BoardPadding)
	y := sq.Pos.Y*sq.Size + float64(cfg.BoardPadding)
	ebitenutil.DrawRect(dst, x, y, sq.Size, sq.Size, sq.Clr)
}

func (sq *Square) Activate(prev *Square) {
	if prev != nil {
		prev.Clr = prev.Orig
	}
	sq.Clr = cfg.ActiveColor
}
