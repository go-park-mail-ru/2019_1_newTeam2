package game

type Game struct {
	MaxRooms int
}

func NewGame() *Game{
	return &Game{
		MaxRooms: 2,
	}
}

func (g *Game) Run() {
}