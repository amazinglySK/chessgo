package board

import (
	"github.com/amazinglySK/chessgo/pkg/cfg"
	"github.com/amazinglySK/chessgo/pkg/events"
	"github.com/amazinglySK/chessgo/pkg/helpers"
	"github.com/amazinglySK/chessgo/pkg/pieces"
	"github.com/amazinglySK/chessgo/pkg/player"
	"github.com/amazinglySK/chessgo/pkg/sounds"
	"github.com/amazinglySK/chessgo/pkg/square"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"image/color"
	"log"
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

// makeSquare makes squares assigned with the right coordinates when given a width and height
func makeSquares(w, h int) [][]*square.Square {
	white := false
	squares := [][]*square.Square{}
	for i := 0; i < h; i++ {
		row := []*square.Square{}
		for j := 0; j < w; j++ {
			var sq_color color.Color
			var sq_color_str string
			if white {
				sq_color = cfg.LightColor
				sq_color_str = "white"
			} else {
				sq_color = cfg.DarkColor
				sq_color_str = "black"
			}

			white = !white
			coords := helpers.Coord{X: float64(j), Y: float64(i)}
			row = append(row, &square.Square{Pos: coords, Size: float64(cfg.SquareSize), Clr: sq_color, Color: sq_color_str, Occupied: false, Orig: sq_color})
		}
		squares = append(squares, row)
		white = !white
	}
	return squares
}

// genWhitePieces initiates all the white pieces with their starting positions
func genWhitePieces(w int) []pieces.Piece {
	pieces_slice := []pieces.Piece{}
	pieces_slice = append(pieces_slice, pieces.InitBishop(helpers.Coord{X: 2, Y: 0}, "white"), pieces.InitBishop(helpers.Coord{X: float64(w - 3), Y: 0}, "white"))
	pieces_slice = append(pieces_slice, pieces.InitRook(helpers.Coord{X: 0, Y: 0}, "white"), pieces.InitRook(helpers.Coord{X: float64(w - 1), Y: 0}, "white"))
	pieces_slice = append(pieces_slice, pieces.InitKnight(helpers.Coord{X: 1, Y: 0}, "white"), pieces.InitKnight(helpers.Coord{X: float64(w - 2), Y: 0}, "white"))
	pieces_slice = append(pieces_slice, pieces.InitQueen(helpers.Coord{X: 3, Y: 0}, "white"))
	pieces_slice = append(pieces_slice, pieces.InitKing(helpers.Coord{X: 4, Y: 0}, "white"))
	for _, pawn := range pieces.InitPawns(1, w, "white") {
		pieces_slice = append(pieces_slice, pawn)
	}
	return pieces_slice
}

// genBlackPieces initiate all the black pieces with their starting positions
func genBlackPieces(w int) []pieces.Piece {
	pieces_slice := []pieces.Piece{}
	pieces_slice = append(pieces_slice, pieces.InitBishop(helpers.Coord{X: 2, Y: float64(w - 1)}, "black"), pieces.InitBishop(helpers.Coord{X: float64(w - 3), Y: float64(w - 1)}, "black"))
	pieces_slice = append(pieces_slice, pieces.InitRook(helpers.Coord{X: 0, Y: float64(w - 1)}, "black"), pieces.InitRook(helpers.Coord{X: float64(w - 1), Y: float64(w - 1)}, "black"))
	pieces_slice = append(pieces_slice, pieces.InitKnight(helpers.Coord{X: 1, Y: float64(w - 1)}, "black"), pieces.InitKnight(helpers.Coord{X: float64(w - 2), Y: float64(w - 1)}, "black"))
	pieces_slice = append(pieces_slice, pieces.InitQueen(helpers.Coord{X: 3, Y: float64(w - 1)}, "black"))
	pieces_slice = append(pieces_slice, pieces.InitKing(helpers.Coord{X: 4, Y: float64(w - 1)}, "black"))
	for _, pawn := range pieces.InitPawns(w-2, w, "black") {
		pieces_slice = append(pieces_slice, pawn)
	}

	return pieces_slice
}

// UpdateInitialSquareState updates the occupied status and the piece on it
func (b *Board) UpdateInitialSquareState(pieces []pieces.Piece) {
	for _, piece := range pieces {
		pos := *piece.GetPos()
		sq := b.GetSquare(pos)
		sq.Piece = piece
		sq.Occupied = true
	}
}

// go:embed move.ogg
var moveRaw []byte

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

// ManageClick manages the click event
func (b *Board) ManageClick() bool {
	clicked, pos := events.CheckMouseEvents(false)
	if clicked {
		// Deactivating all previously activated sqs
		square.Deactivate(b.PrevActives)

		sq := b.GetSquare(pos)
		curr_player := b.Players[b.CurrPlayerIdx]

		// Selecting a piece
		if sq.Occupied && sq.Piece.GetColor() == curr_player.Color {

			piece := sq.Piece
			curr_player.CurrPiece = piece
			coords := piece.GenValidMoves()
			var moves []helpers.Coord
			switch piece.(type) {
			case *pieces.Pawn:
				moves = curr_player.FilterPawnMoves(piece, coords, b.Squares)
				break
			default:
				moves = curr_player.FilterMoves(coords, b.Squares)
			}

			// Acitvating the current piece square
			sq.Activate()
			b.PrevActives = append(b.PrevActives, sq)

			curr_player.ValidMoves = moves

			// Activating all move squares
			for _, move := range moves {
				b.GetSquare(move).Activate()
				sq = b.GetSquare(move)
				b.PrevActives = append(b.PrevActives, sq)
			}

			return false
		}

		// Probably a capture or a normal move
		if curr_player.CurrPiece != nil {
			var king_cap = false
			if (sq.Occupied && sq.Piece.GetColor() != curr_player.Color) || !sq.Occupied {
				// Plays the sounds
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

				prev_sq := b.GetSquare(*curr_player.CurrPiece.GetPos())
				move_sq := b.GetSquare(pos)
				switch move_sq.Piece.(type) {
				case *pieces.King:
					king_cap = true

				}
				if curr_player.MovePiece(curr_player.CurrPiece, move_sq, prev_sq) {
					curr_player.CurrPiece = nil
					// Swapping the current player
					b.CurrPlayerIdx = uint(len(b.Players)) - (b.CurrPlayerIdx + 1)

				} else {
					log.Println("invalid move")
				}

				return king_cap
			}
		}

	}

	return false
}

// Update updates the board's state according to click and other event(s)
func (b *Board) Update() (bool, string) {
	if b.ManageClick() {
		b.CurrPlayerIdx = uint(len(b.Players)) - (b.CurrPlayerIdx + 1)
		return true, b.Players[b.CurrPlayerIdx].Color
	}


	return false, ""
}

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
