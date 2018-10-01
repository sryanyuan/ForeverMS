package inventory

import (
	"github.com/sryanyuan/ForeverMS/core/consts"
)

type Item struct {
	ItemID      int
	Position    int
	Quantity    int
	Flag        int
	Owner       string
	Expiration  int64
	SN          int
	UID         int64
	SenderID    int
	SendMessage string
}

func NewItem(itemID int, position int, quantity int) *Item {
	item := &Item{
		ItemID:   itemID,
		Position: position,
		Quantity: quantity,
	}
	// TODO: Generate uid for cash item
	return item
}

func (i *Item) GetSenderID() int {
	return i.SenderID
}

func (i *Item) SetSenderID(sid int) {
	i.SenderID = sid
}

func (i *Item) GetSendMessage() string {
	return i.SendMessage
}

func (i *Item) SetSendMessage(msg string) {
	i.SendMessage = msg
}

func (i *Item) GetItemID() int {
	return i.ItemID
}

func (i *Item) SetItemID(id int) {
	i.ItemID = id
}

func (i *Item) GetFlag() int {
	return i.Flag
}

func (i *Item) SetFlag(flag int) {
	i.Flag = flag
}

func (i *Item) GetPosition() int {
	return i.Position
}

func (i *Item) SetPosition(pos int) {
	i.Position = pos
}

func (i *Item) GetOwner() string {
	return i.Owner
}

func (i *Item) SetOwner(o string) {
	i.Owner = o
}

func (i *Item) GetQuantity() int {
	return i.Quantity
}

func (i *Item) SetQuantity(q int) {
	i.Quantity = q
}

func (i *Item) GetExpiration() int64 {
	return i.Expiration
}

func (i *Item) SetExpiration(e int64) {
	i.Expiration = e
}

func (i *Item) GetSN() int {
	return i.SN
}

func (i *Item) SetSN(sn int) {
	i.SN = sn
}

func (i *Item) GetUID() int64 {
	return i.UID
}

func (i *Item) SetUID(uid int64) {
	i.UID = uid
}

func (i *Item) GetType() int {
	return consts.ItemType.Item
}

func (i *Item) Copy() IItem {
	var item Item
	item = *i
	return &item
}
