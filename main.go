package main

import (
	"github.com/amazinglySK/chessgo/pkg/board"
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)


func check(err error) {
	if err != nil {
		panic(err)
	}
}

var game_board board.Board

func init() {
	game_board = board.InitBoard(8, 8)
}

type Game struct {
	PlayerIdx uint
	Board     board.Board
}

func (g *Game) Update() error {
	g.Board.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{175, 129, 28, 2})
	g.Board.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return cfg.WindowWidth, cfg.WindowHeight
}

func main() {
	ebiten.SetWindowSize(cfg.WindowWidth, cfg.WindowHeight)
	ebiten.SetWindowTitle("Chess In Go")
	if err := ebiten.RunGame(&Game{PlayerIdx:1, Board : game_board}); err != nil {
		log.Fatal(err)
	}
}
