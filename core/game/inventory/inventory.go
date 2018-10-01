package inventory

import (
	"github.com/sryanyuan/ForeverMS/core/consts"
)

const (
	SlotNA = 0xff
)

type Inventory struct {
	typ       int
	slots     int
	inventory map[byte]IItem
}

func NewInventory(typ int, slots int) *Inventory {
	v := &Inventory{
		typ:       typ,
		slots:     slots,
		inventory: make(map[byte]IItem),
	}
	return v
}

func (i *Inventory) GetSlots() int {
	return i.slots
}

func (i *Inventory) GetType() int {
	return i.typ
}

func (i *Inventory) CountByItemID(itemID int) int {
	cnt := 0
	for _, v := range i.inventory {
		if v.GetItemID() == itemID {
			cnt++
		}
	}
	return cnt
}

func (i *Inventory) FindByUID(uid int64) IItem {
	for _, v := range i.inventory {
		if v.GetUID() == uid {
			return v
		}
	}
	return nil
}

func (i *Inventory) FindByItemID(itemID int) IItem {
	for _, v := range i.inventory {
		if v.GetItemID() == itemID {
			return v
		}
	}
	return nil
}

func (i *Inventory) ListByItemID(itemID int) []IItem {
	res := make([]IItem, 0, 8)
	for _, v := range i.inventory {
		if v.GetItemID() == itemID {
			res = append(res, v)
		}
	}
	if len(res) == 0 {
		// Avoid copying
		return nil
	}
	// TODO: Sort items
	return res
}

func (i *Inventory) List() []IItem {
	if len(i.inventory) == 0 {
		return nil
	}
	res := make([]IItem, 0, len(i.inventory))
	for _, v := range i.inventory {
		res = append(res, v)
	}
	return res
}

func (i *Inventory) AddItem(item IItem) byte {
	slot := i.getNextFreeSlot()
	if slot == SlotNA {
		return SlotNA
	}
	i.inventory[slot] = item
	item.SetPosition(int(slot))
	return slot
}

func (i *Inventory) AddItemFromDB(item IItem) {
	if item.GetPosition() > 127 && item.GetType() != consts.InventoryType.Equipped {
		// TODO: Warnning
	}

	i.inventory[byte(item.GetPosition())] = item
}

func (i *Inventory) Swap(source IItem, target IItem) {
	delete(i.inventory, byte(source.GetPosition()))
	delete(i.inventory, byte(target.GetPosition()))
	sourcePosition := byte(source.GetPosition())
	source.SetPosition(target.GetPosition())
	target.SetPosition(int(sourcePosition))
	i.inventory[byte(source.GetPosition())] = source
	i.inventory[byte(target.GetPosition())] = target
}

func (i *Inventory) GetItem(slot byte) IItem {
	item, ok := i.inventory[slot]
	if !ok || nil == item {
		return nil
	}
	return item
}

func (i *Inventory) RemoveItem(slot byte) {
	i.RemoveItemWithQuantity(slot, 1, false)
}

func (i *Inventory) RemoveItemWithQuantity(slot byte, quantity int, allowZero bool) {
	item, ok := i.inventory[slot]
	if !ok || nil == item {
		return
	}
	item.SetQuantity(item.GetQuantity() - quantity)
	if item.GetQuantity() < 0 {
		item.SetQuantity(0)
	}
	if item.GetQuantity() == 0 && !allowZero {
		i.RemoveSlot(slot)
	}
}

func (i *Inventory) RemoveSlot(slot byte) {
	delete(i.inventory, slot)
}

func (i *Inventory) getNextFreeSlot() byte {
	if i.IsFull(0) {
		return SlotNA
	}
	for pos := 1; pos <= i.slots; pos++ {
		if _, ok := i.inventory[byte(pos)]; !ok {
			return byte(pos)
		}
	}
	return SlotNA
}

func (i *Inventory) IsFull(margin int) bool {
	return len(i.inventory)+margin >= i.slots
}

func (i *Inventory) Walk(fn func(k byte, v IItem) error) {
	for ik, iv := range i.inventory {
		if err := fn(ik, iv); nil != err {
			break
		}
	}
}
