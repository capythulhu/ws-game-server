package shared

type Profile struct {
	Nick rune
}

type Player struct {
	MatchMap *Map `json:"-"`

	Profile
	Position Coordinate
	Velocity uint
}

func (p *Player) Move(direction Coordinate) {
	switch true {
	case direction.X < 0:
		p.Position.X = max(p.Position.X-int(p.Velocity), 0)
	case direction.X > 0:
		p.Position.X = min(p.Position.X+int(p.Velocity), p.MatchMap.Size.X-1)
	}
	switch true {
	case direction.Y < 0:
		p.Position.Y = max(p.Position.Y-int(p.Velocity), 0)
	case direction.Y > 0:
		p.Position.Y = min(p.Position.Y+int(p.Velocity), p.MatchMap.Size.Y-1)
	}
}
