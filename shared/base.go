package shared

type Coordinate struct {
	X, Y int
}

func (c Coordinate) Equals(d Coordinate) bool {
	return c.X == d.X && c.Y == d.Y
}

type Actor struct {
	Position Coordinate
}
