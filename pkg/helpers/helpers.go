package helpers

type Coord struct {
	X float64
	Y float64
}

func (c Coord) Equal(comp Coord) bool {
	return c.X == comp.X && c.Y == comp.Y
}
