package scenes

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface{ 
	Draw(dst *ebiten.Image)
}
