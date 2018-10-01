package world

type MapEffect struct {
	active bool
	itemID int
	msg    string
}

func NewMapEffect(msg string, itemID int) *MapEffect {
	return &MapEffect{
		msg:    msg,
		itemID: itemID,
	}
}
