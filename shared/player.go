package shared

type Player struct {
	Map      *Map
	Position Coordinate
	Velocity int
	Nick     rune
}

func (p *Player) Move(direction Coordinate) {
	switch true {
	case direction.X < 0:
		p.Position.X = max(p.Position.X-p.Velocity, 0)
	case direction.X > 0:
		p.Position.X = min(p.Position.X+p.Velocity, p.Map.Size.X-1)
	}
	switch true {
	case direction.Y < 0:
		p.Position.Y = max(p.Position.Y-p.Velocity, 0)
	case direction.Y > 0:
		p.Position.Y = min(p.Position.Y+p.Velocity, p.Map.Size.Y-1)
	}
}
