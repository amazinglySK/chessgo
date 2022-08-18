package sounds

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"bytes"
	_ "embed"
)

//go:embed capture.ogg
var CaptureRaw []byte

//go:embed move.ogg
var MoveRaw []byte

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func GetTracks(ctx *audio.Context) (*audio.Player, *audio.Player){
	
	m, err := vorbis.Decode(ctx, bytes.NewReader(MoveRaw))
	check(err)
	c, err := vorbis.Decode(ctx, bytes.NewReader(CaptureRaw))
	check(err)
	capture, err := ctx.NewPlayer(c)
	check(err)
	move, err := ctx.NewPlayer(m)
	check(err)
	
	return capture, move

}
