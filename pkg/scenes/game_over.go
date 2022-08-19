package scenes

import (
	_ "embed"
	"fmt"
	"github.com/amazinglySK/chessgo/pkg/events"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image/color"
	"log"
)


//go:embed scene_font.ttf
var scene_font []byte

type GameOverScene struct {
	font   font.Face
	winner string
}


func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitGameOverScene(_w, _h int) GameOverScene {
	tt, err := opentype.Parse(scene_font)
	check(err)

	const dpi = 80 
	mplusNormalFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	check(err)


	return GameOverScene{font: mplusNormalFont}

}

func (g GameOverScene) Draw(dst *ebiten.Image) {
	const x = 20
	text.Draw(dst, "Game Over", g.font, x, 60, color.White)
	text.Draw(dst, fmt.Sprintf("%v won", g.winner), g.font, x, 150, color.White)
	text.Draw(dst, "Press R to restart", g.font, x, 240, color.White)

}

func (g *GameOverScene) SetWinner(winner string) {
	g.winner = winner
}

func (g GameOverScene) Update() bool {
	if events.CheckKbEvents("R") {
		return true
	}

	return false
}
