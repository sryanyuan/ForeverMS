package world

type World struct {
	players map[int64]*Player
}

func NewWorld() *World {
	return &World{
		players: make(map[int64]*Player),
	}
}

func (w *World) AddPlayer(p *Player) {
	w.players[p.charModel.ID] = p
}

func (w *World) GetPlayer(id int64) *Player {
	return w.players[id]
}
