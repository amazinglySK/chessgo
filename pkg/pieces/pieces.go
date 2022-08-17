package pieces

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"

	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/hajimehoshi/ebiten/v2"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}


//go:embed sprites.png
var sprites []byte

var (
	WhiteSprites *ebiten.Image
	BlackSprites *ebiten.Image
)

func init() {
	sprite, _, err := image.Decode(bytes.NewReader(sprites))
	check(err)
	SpriteSheet := ebiten.NewImageFromImage(sprite)
	WhiteSprites = SpriteSheet.SubImage(image.Rect(0, 0, 1200, 200)).(*ebiten.Image)
	BlackSprites = SpriteSheet.SubImage(image.Rect(0, 200, 1200, 400)).(*ebiten.Image)
}

type Piece interface {
	Draw(dst *ebiten.Image)
	GenValidMoves() [][]helpers.Coord
	GetPos() *helpers.Coord
	GetColor() string
	Move(helpers.Coord)
}

func generateDrawingOps(pos helpers.Coord) *ebiten.DrawImageOptions {
	op := &ebiten.DrawImageOptions{}
	t_x := float64(cfg.BoardPadding) + (pos.X)*float64(cfg.SquareSize)
	t_y := float64(cfg.BoardPadding) + (pos.Y)*float64(cfg.SquareSize)
	op.GeoM.Scale(cfg.ScaleFactor, cfg.ScaleFactor)
	op.GeoM.Translate(t_x, t_y)

	return op
}

func generateSprite(color string, SpriteIndex int) *ebiten.Image {
	var sprite *ebiten.Image
	var color_idx int
	if color == "white" {
		color_idx = 0
	} else {
		color_idx = 1
	}
	SpriteRect := image.Rect(SpriteIndex*cfg.PieceSize, cfg.PieceSize*color_idx, (SpriteIndex+1)*(cfg.PieceSize), cfg.PieceSize*(color_idx+1))
	switch color {
	case "white":
		sprite = WhiteSprites.SubImage(SpriteRect).(*ebiten.Image)
	case "black":
		sprite = BlackSprites.SubImage(SpriteRect).(*ebiten.Image)
	}

	return sprite
}

func filterNegatives(coord [][]helpers.Coord) [][]helpers.Coord {
	filtered := [][]helpers.Coord{}
	for _, set := range coord {
		ver_sets := []helpers.Coord{}
		for _, c := range set {
			if c.X < 0 || c.Y < 0 || c.X >= float64(cfg.SquareCount) || c.Y >= float64(cfg.SquareCount) {
				continue
			}
			ver_sets = append(ver_sets, c)
		}

		if len(ver_sets) >= 1 {
			filtered = append(filtered, ver_sets)
		}
	}

	return filtered
}

func GenDiagonalMoves(X, Y int) [][]helpers.Coord {
	moves := [][]helpers.Coord{}
	set := []helpers.Coord{}
	x := X
	y := Y
	for x < cfg.SquareCount-1 && y < cfg.SquareCount {
		x += 1
		y += 1
		set = append(set, helpers.Coord{float64(x), float64(y)})
	}

	moves = append(moves, set)
	set = nil
	x = int(X)
	y = int(Y)

	for x < cfg.SquareCount-1 && y > 0 {
		y -= 1
		x += 1
		set = append(set, helpers.Coord{float64(x), float64(y)})
	}

	moves = append(moves, set)
	set = nil
	x = int(X)
	y = int(Y)

	for x > 0 && y < cfg.SquareCount-1 {
		x -= 1
		y += 1
		set = append(set, helpers.Coord{float64(x), float64(y)})
	}

	moves = append(moves, set)
	set = nil
	x = int(X)
	y = int(Y)

	for x > 0 && y > 0 {
		x -= 1
		y -= 1
		set = append(set, helpers.Coord{float64(x), float64(y)})
	}

	moves = append(moves, set)

	return filterNegatives(moves)

}

func GenStraightMoves(X, Y float64) [][]helpers.Coord {
	moves := [][]helpers.Coord{}
	set := []helpers.Coord{}
	// The left side
	for i := int(X) - 1; i >= 0; i-- {
		set = append(set, helpers.Coord{float64(i), Y})
	}

	moves = append(moves, set)
	set = nil

	// The top side
	for i := int(Y) - 1; i >= 0; i-- {
		set = append(set, helpers.Coord{X, float64(i)})
	}

	moves = append(moves, set)
	set = nil

	// The right side
	for i := int(X) + 1; i <= cfg.SquareCount; i++ {
		set = append(set, helpers.Coord{float64(i), Y})
	}

	moves = append(moves, set)
	set = nil

	// The bottom side
	for i := int(Y) + 1; i <= cfg.SquareCount; i++ {
		set = append(set, helpers.Coord{X, float64(i)})
	}

	moves = append(moves, set)
	set = nil

	return moves
}
