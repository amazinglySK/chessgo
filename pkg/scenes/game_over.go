package scenes

import (
	"github.com/amazinglySK/chessgo/pkg/events"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	_ "image/png"
	"image"
	"image/color"
	"log"
)

type GameOverScene struct {
	font   font.Face
	winner string
	replay *ebiten.Image
}

func InitGameOverScene(w, h int) GameOverScene {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	
	img, _, err := ebitenutil.NewImageFromFile("./pkg/scenes/reload.png")
	if err != nil {
		log.Fatal(err)
	}

	return GameOverScene{font: mplusNormalFont, replay: img}

}

func (g GameOverScene) Draw(dst *ebiten.Image) {
	const x = 20
	text.Draw(dst, "Game Over", g.font, x, 40, color.White)
	text.Draw(dst, g.winner, g.font, x, 60, color.White)

	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Translate(x, 80)
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
