package cfg

import (
	"image/color"
)

var (
	LightColor   color.Color = color.RGBA{255, 211, 164, 255}
	DarkColor    color.Color = color.RGBA{152, 97, 39, 255}
	ActiveColor  color.Color = color.RGBA{255, 255, 167, 255}
	WindowWidth  int         = 700
	WindowHeight int         = 700
	BoardPadding int         = 30
	SquareSize   int         = 80
	SquareCount  int         = 8
)
