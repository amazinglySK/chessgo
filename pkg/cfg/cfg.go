package cfg

import (
	"image/color"
)

var (
	// Colors
	LightColor      color.Color = color.RGBA{216, 227, 230, 255}
	DarkColor       color.Color = color.RGBA{123, 157, 178, 255}
	ActiveColor     color.Color = color.RGBA{255, 255, 167, 200}
	BackgroundColor color.Color = color.RGBA{49, 46, 43, 200}

	// Window
	WindowWidth  int = 700
	WindowHeight int = 700

	// Board
	BoardPadding int = 30
	SquareSize   int = 80
	SquareCount  int = 8

	// Pieces
	PieceSize   int            = 200
	ScaleFactor float64        = .4
	SpriteMap   map[string]int = map[string]int{"pawn": 5, "queen": 1, "king": 0, "rook": 4, "knight": 3, "bishop": 2}
)
