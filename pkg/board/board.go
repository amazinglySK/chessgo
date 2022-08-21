package board

import (
	"github.com/amazinglySK/chessgo/pkg/events"
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/amazinglySK/chessgo/pkg/pieces"
	"github.com/amazinglySK/chessgo/pkg/player"
	"github.com/amazinglySK/chessgo/pkg/sounds"
	"github.com/amazinglySK/chessgo/pkg/square"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

// Board is the main struct holding data for the entire game
type Board struct {
	width, height int
	Squares       [][]*square.Square
	CurrPlayerIdx uint
	Players       []*player.Player
	PrevActives   []*square.Square
	CaptureSound  *audio.Player
	MoveSound     *audio.Player
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func join(a []pieces.Piece, b []pieces.Piece) []pieces.Piece {
	for _, i := range a {
		b = append(b, i)
	}

	return b
}

// InitBoard initiates a normal board of given width and height
func InitBoard(w int, h int, actx *audio.Context) Board {
	// Squares
	squares := makeSquares(w, h)
	all_pieces := join(genWhitePieces(w), genBlackPieces(w))
	whitePlayer := player.Player{Color: "white"}
	blackPlayer := player.Player{Color: "black"}
	capture, move := sounds.GetTracks(actx)
	board := Board{Squares: squares, CurrPlayerIdx: 0, Players: []*player.Player{&whitePlayer, &blackPlayer}, PrevActives: []*square.Square{}, MoveSound: move, CaptureSound: capture, width:w, height:h}

	board.UpdateInitialSquareState(all_pieces)

	return board
}

// GetSquare gets the square at a given position
func (b *Board) GetSquare(pos helpers.Coord) *square.Square {
	x, y := int(pos.X), int(pos.Y)
	return b.Squares[y][x]
}

// Draw draws the board on the screen
func (b Board) Draw(dst *ebiten.Image) {
	for _, row := range b.Squares {
		for _, s := range row {
			s.Draw(dst)
			if s.Occupied {
				s.Piece.Draw(dst)
			}
		}
	}

}


// PlaySound plays the necessary sounds
func (b *Board) PlaySound(sq *square.Square) {

	switch sq.Occupied {
	// Capture
	case true:
		b.CaptureSound.Rewind()
		b.CaptureSound.Play()

	// Move
	case false:
		b.MoveSound.Rewind()
		b.MoveSound.Play()
	}

}

// ManageClick manages the click event return if a player won or not
func (b *Board) ManageClick() bool {
	clicked, pos := events.CheckMouseEvents(false)


	if clicked {
		// Deactivating all previously activated sqs
		square.Deactivate(b.PrevActives)

		// Out of bounds
		if int(pos.X) < 0 || int(pos.X) >= b.width || int(pos.Y) < 0 || int(pos.Y) >= b.height {
			return false
		}

		sq := b.GetSquare(pos)
		curr_player := b.Players[b.CurrPlayerIdx]

		// Selecting a piece
		if sq.Occupied && sq.Piece.GetColor() == curr_player.Color {
			b.HandleSelect(sq, curr_player)
			return false
		}
	
		// Probably a capture or a normal move
		if curr_player.CurrPiece != nil {
			if (sq.Occupied && sq.Piece.GetColor() != curr_player.Color) || !sq.Occupied {
				// Plays the sounds
				b.PlaySound(sq)
				king_cap := b.HandleMove(curr_player, pos)
				return king_cap
			}
		}
	}

	return false
}

// Update updates the board's state according to click and other event(s)
func (b *Board) Update() (bool, string) {
	// King is captured
	if b.ManageClick() {
		b.CurrPlayerIdx = uint(len(b.Players)) - (b.CurrPlayerIdx + 1)
		return true, b.Players[b.CurrPlayerIdx].Color
	}

	return false, ""
}

// Reset resets the board's state
func (b *Board) Reset() {
	for _, i := range b.Squares {
		for _, j := range i {
			j.Piece = nil
			j.Occupied = false
		}
	}

	for _, player := range b.Players {
		player.ValidMoves = nil
		player.CurrPiece = nil
	}

	b.UpdateInitialSquareState(join(genWhitePieces(b.width), genBlackPieces(b.width)))
	b.CurrPlayerIdx = 0
	b.PrevActives = nil
}
