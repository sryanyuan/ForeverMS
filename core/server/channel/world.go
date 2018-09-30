package channel

// MSWorld represents a world in a channel
type MSWorld struct {
	channelID int
	players   map[int64]*MSPlayer
}

func NewMSWorld(channelID int) *MSWorld {
	return &MSWorld{
		channelID: channelID,
		players:   make(map[int64]*MSPlayer),
	}
}

func (w *MSWorld) GetPlayer(id int64) *MSPlayer {
	return w.players[id]
}

func (w *MSWorld) AddPlayer(p *MSPlayer) {
	w.players[p.charModel.ID] = p
}
