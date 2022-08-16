package events

import (
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func CheckMouseEvents() (bool, helpers.Coord) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		x, y = (x-cfg.BoardPadding)/cfg.SquareSize, (y-cfg.BoardPadding)/cfg.SquareSize
		return true, helpers.Coord{X: float64(x), Y: float64(y)}
	}
	return false, helpers.Coord{}
}
