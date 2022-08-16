package helpers

type Coord struct {
	X float64
	Y float64
}

func (c Coord) Equal(comp Coord) bool {
	return int(c.X) == int(comp.X) && int(c.Y) == int(comp.Y)
}
