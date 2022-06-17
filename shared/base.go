package shared

type Coordinate struct {
	x, y int
}

func (c Coordinate) Equals(d Coordinate) bool {
	return c.x == d.x && c.y == d.y
}

type Actor struct {
	position Coordinate
}
