package main

import (
	"log"

	"bytes"
	"github.com/amazinglySK/chessgo/pkg/board"
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
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

// go:embed sound.ogg
var moveRaw []byte

func main() {
	ebiten.SetWindowSize(cfg.WindowWidth, cfg.WindowHeight)
	ebiten.SetWindowTitle("Chess In Go")
	var audioContext = audio.NewContext(32000)
	s, err := vorbis.Decode(audioContext, bytes.NewReader(moveRaw))
	check(err)
	p, err := audioContext.NewPlayer(s)
	check(err)
	if err := ebiten.RunGame(&Game{Board: board.InitBoard(8, 8, p)}); err != nil {
		log.Fatal(err)
	}
}
