package main

import (
	"github.com/amazinglySK/chessgo/pkg/board"
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)


func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Game struct {
	Board     board.Board
}

func (g *Game) Update() error {
	g.Board.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(cfg.BackgroundColor)
	g.Board.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return cfg.WindowWidth, cfg.WindowHeight
}

func main() {
	ebiten.SetWindowSize(cfg.WindowWidth, cfg.WindowHeight)
	ebiten.SetWindowTitle("Chess In Go")
	if err := ebiten.RunGame(&Game{Board : board.InitBoard(8, 8)}); err != nil {
		log.Fatal(err)
	}
}
