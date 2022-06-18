package shared

type Profile struct {
	Nick rune
}

type Player struct {
	Profile
	Position Coordinate
	Velocity uint
}

func (p *Player) Move(direction Coordinate, matchMap *Map) {
	switch true {
	case direction.X < 0:
		p.Position.X = max(p.Position.X-int(p.Velocity), 0)
	case direction.X > 0:
		p.Position.X = min(p.Position.X+int(p.Velocity), matchMap.Size.X-1)
	}
	switch true {
	case direction.Y < 0:
		p.Position.Y = max(p.Position.Y-int(p.Velocity), 0)
	case direction.Y > 0:
		p.Position.Y = min(p.Position.Y+int(p.Velocity), matchMap.Size.Y-1)
	}
}
