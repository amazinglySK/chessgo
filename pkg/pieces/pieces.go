package pieces

import (
	"bytes"
	_ "embed"
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

//go:embed white_pieces.png
var whiteSpritesheet []byte

//go:embed black_pieces.png
var blackSpritesheet []byte

var (
	WhiteSprites *ebiten.Image
	BlackSprites *ebiten.Image
)

func init() {

	white_img, _, err := image.Decode(bytes.NewReader(whiteSpritesheet))
	check(err)

	black_img, _, err := image.Decode(bytes.NewReader(blackSpritesheet))
	check(err)

	WhiteSprites = ebiten.NewImageFromImage(white_img)
	BlackSprites = ebiten.NewImageFromImage(black_img)

}

type Piece interface {
	Draw(dst *ebiten.Image)
	GenValidMoves() [][]helpers.Coord
	GetPos() *helpers.Coord
	GetColor() string
	Move(helpers.Coord)
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
	for i := int(X) - 1; i>= 0; i-- {
		set = append(set, helpers.Coord{float64(i), Y})
	}

	moves = append(moves, set)
	set = nil

	// The top side
	for i := int(Y) - 1; i>=0 ; i-- {
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
