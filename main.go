package main

import (
	"log"

	"github.com/amazinglySK/chessgo/pkg/board"
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/amazinglySK/chessgo/pkg/scenes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type Game struct {
	currScene scenes.Scene
	Board     board.Board
	GameOver  scenes.GameOverScene
}

func (g *Game) Update() error {
	switch g.currScene.(type) {
	case board.Board:
		game_over, winner := g.Board.Update()
		if game_over {
			g.GameOver.SetWinner(winner)
			g.currScene = g.GameOver
		}
	case scenes.GameOverScene:
		replay := g.GameOver.Update()
		if replay {
			g.Board.Reset()
			g.currScene = g.Board
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(cfg.BackgroundColor)
	g.currScene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return cfg.WindowWidth, cfg.WindowHeight
}

func main() {
	ebiten.SetWindowSize(cfg.WindowWidth, cfg.WindowHeight)
	ebiten.SetWindowTitle("Chess In Go")
	var audioContext = audio.NewContext(32000)
	board := board.InitBoard(8, 8, audioContext)
	game_over := scenes.InitGameOverScene(cfg.WindowWidth, cfg.WindowHeight)
	if err := ebiten.RunGame(&Game{Board: board, GameOver: game_over, currScene: board}); err != nil {
		log.Fatal(err)
	}
}
