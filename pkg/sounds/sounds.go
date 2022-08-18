package sounds

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"bytes"
	_ "embed"
)

// go:embed capture.ogg
var CaptureRaw []byte

// go:embed move.ogg
var MoveRaw []byte

func GetTracks(ctx *audio.Context) (*audio.Player){
	
	s, err := vorbis.Decode(ctx, bytes.NewReader(MoveRaw))
	if err != nil {
		log.Println("Paniced in decoding")
		panic(err)
	}
	capture, err := ctx.NewPlayer(s)
	if err != nil {
		log.Println("Paniced in creating a new player")
		panic(err)
	}
	
	return capture

}
