package scenes

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/amazinglySK/chessgo/pkg/events"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"image"
	"image/color"
	_ "image/png"
	"log"
)

//go:embed reload.png
var replayImg []byte

//go:embed scene_font.ttf
var scene_font []byte

type GameOverScene struct {
	font   font.Face
	winner string
	replay *ebiten.Image
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitGameOverScene(w, h int) GameOverScene {
	tt, err := opentype.Parse(scene_font)
	check(err)

	const dpi = 80 
	mplusNormalFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	check(err)

	img, _, err := image.Decode(bytes.NewReader(replayImg))
	check(err)

	sprite := ebiten.NewImageFromImage(img)

	return GameOverScene{font: mplusNormalFont, replay: sprite}

}

func (g GameOverScene) Draw(dst *ebiten.Image) {
	const x = 20
	text.Draw(dst, "Game Over", g.font, x, 60, color.White)
	text.Draw(dst, fmt.Sprintf("%v won", g.winner), g.font, x, 150, color.White)

	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Translate(x, 200)
	dst.DrawImage(g.replay, ops)
}

func (g *GameOverScene) SetWinner(winner string) {
	g.winner = winner
}

func (g GameOverScene) Update() bool {
	click, pos := events.CheckMouseEvents(true)

	if click {
		if image.Pt(int(pos.X), int(pos.Y)).In(g.replay.Bounds()) {
			return true
		}
	}

	return false
}
