package main

import (
	"log"
	"github.com/amazinglySK/chessgo/pkg/board"
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Game struct {
	Board board.Board
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
	var audioContext = audio.NewContext(32000)
	if err := ebiten.RunGame(&Game{Board: board.InitBoard(8, 8, audioContext)}); err != nil {
		log.Fatal(err)
	}
}
