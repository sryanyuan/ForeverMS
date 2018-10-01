package inventory

import (
	"github.com/sryanyuan/ForeverMS/core/consts"
	"github.com/sryanyuan/ForeverMS/core/models"
)

type Equip struct {
	Item
	Model       models.InventoryEquipment
	PartnerUID  int64
	PartnerName string
	ItemExp     int
	ItemLevel   int
	ScrollRes   int
}

func NewEquip(itemID int, position int, ringID int) *Equip {
	var equip Equip
	equip.Item = *NewItem(itemID, position, 1)
	equip.Model.RingID = ringID
	return &equip
}

func (i *Equip) GetType() int {
	return consts.ItemType.Equip
}
