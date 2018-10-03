package world

import (
	"time"

	"github.com/sryanyuan/ForeverMS/core/netio"

	"github.com/juju/errors"
	"github.com/sryanyuan/ForeverMS/core/consts"
	"github.com/sryanyuan/ForeverMS/core/game/inventory"
	"github.com/sryanyuan/ForeverMS/core/models"
)

type localStatInfo struct {
	MaxHP          int
	MaxMP          int
	Dex            int
	Intt           int
	Str            int
	Luk            int
	Watk           int
	Matk           int
	Wdef           int
	Mdef           int
	Magic          int
	SpeedMod       float64
	JumpMod        float64
	MaxBasedDamage int
}

type Player struct {
	netio.IConn
	charModel *models.Character
	// Inventorys
	inventorys []*inventory.Inventory
	// Local status
	localStat localStatInfo
}

func NewPlayer(conn netio.IConn) *Player {
	player := &Player{
		IConn: conn,
	}
	return player
}

func (p *Player) LoadCharacterData(charID int) error {
	var err error
	p.charModel, err = models.SelectCharacterByCharacterID(int32(charID))
	if nil != err {
		return errors.Trace(err)
	}

	if err = p.loadInventorys(); nil != err {
		return errors.Trace(err)
	}
	if err = p.loadQuestStatus(); nil != err {
		return errors.Trace(err)
	}
	if err = p.loadSkills(); nil != err {
		return errors.Trace(err)
	}
	if err = p.loadSkillMacros(); nil != err {
		return errors.Trace(err)
	}
	if err = p.loadKeyMap(); nil != err {
		return errors.Trace(err)
	}
	if err = p.loadSavedLocations(); nil != err {
		return errors.Trace(err)
	}
	if err = p.loadFameLog(); nil != err {
		return errors.Trace(err)
	}
	if err = p.loadCharAresInfo(); nil != err {
		return errors.Trace(err)
	}
	if err = p.loadAchievements(); nil != err {
		return errors.Trace(err)
	}

	return nil
}

func (p *Player) getInventory(typ int) *inventory.Inventory {
	return p.inventorys[typ]
}

func (p *Player) loadInventorys() error {
	// Load inventorys
	p.inventorys = make([]*inventory.Inventory, consts.InventoryType.Total)
	p.inventorys[consts.InventoryType.Equip] = inventory.NewInventory(consts.InventoryType.Equip, p.charModel.EquipSlots)
	p.inventorys[consts.InventoryType.Use] = inventory.NewInventory(consts.InventoryType.Use, p.charModel.UseSlots)
	p.inventorys[consts.InventoryType.Setup] = inventory.NewInventory(consts.InventoryType.Setup, p.charModel.SetupSlots)
	p.inventorys[consts.InventoryType.Etc] = inventory.NewInventory(consts.InventoryType.Etc, p.charModel.EtcSlots)
	p.inventorys[consts.InventoryType.Cash] = inventory.NewInventory(consts.InventoryType.Cash, 100)
	p.inventorys[consts.InventoryType.Equipped] = inventory.NewInventory(consts.InventoryType.Equipped, 24)

	// Load inventory equipments
	items, err := models.SelectCharacterInventoryItem(p.charModel.ID)
	if nil != err {
		return errors.Trace(err)
	}
	for _, item := range items {
		if item.InventoryType == consts.InventoryType.Equip ||
			item.InventoryType == consts.InventoryType.Equipped {
			equip := inventory.NewEquip(item.ItemID, item.Position, 0)
			if item.RingID > 0 {
				// TODO: Load ring ?
			} else {
				equip.SetOwner(item.Owner)
				equip.SetQuantity(item.Quantity)
				equip.Model = item.InventoryEquipment
				equip.SetUID(item.UniqueID)
			}
			if item.Expiration != 0 {
				if time.Now().Unix() > item.Expiration {
					// TODO: expire item
					continue
				}
				equip.SetExpiration(item.Expiration)
			}
			// Add to inventory
			p.getInventory(item.InventoryType).AddItemFromDB(equip)
			continue
		}
		// Add none-quip item
		nitem := inventory.NewItem(item.ItemID, item.Position, item.Quantity)
		nitem.SetOwner(item.Owner)
		nitem.SetUID(item.UniqueID)
		if item.Expiration != 0 {
			if time.Now().Unix() > item.Expiration {
				// TODO: expire item
				continue
			}
			nitem.SetExpiration(item.Expiration)
		}
		// Add to inventory
		p.getInventory(item.InventoryType).AddItemFromDB(nitem)
		// TODO: petslot ?
	}

	return nil
}

// TODO: load quest
func (p *Player) loadQuestStatus() error {
	return nil
}

// TODO: load kills
func (p *Player) loadSkills() error {
	return nil
}

// TODO: load kills macros
func (p *Player) loadSkillMacros() error {
	return nil
}

// TODO: load key map
func (p *Player) loadKeyMap() error {
	return nil
}

// TODO: load saved locations
func (p *Player) loadSavedLocations() error {
	return nil
}

// TODO: load fame log
func (p *Player) loadFameLog() error {
	return nil
}

// TODO: load char ares info
func (p *Player) loadCharAresInfo() error {
	return nil
}

// TODO: load achievements
func (p *Player) loadAchievements() error {
	return nil
}

func (p *Player) updateLocalStats() {
	oldMaxHP := p.localStat.MaxHP
	p.localStat.MaxHP = p.charModel.MaxHP
	p.localStat.MaxMP = p.charModel.MaxMP
	p.localStat.Dex = p.charModel.Dex
	p.localStat.Intt = p.charModel.Intt
	p.localStat.Str = p.charModel.Str
	p.localStat.Luk = p.charModel.Luk
	p.localStat.Magic = p.localStat.Intt
	p.localStat.Watk = 0

	speed := 100
	jump := 100

	p.getInventory(consts.InventoryType.Equipped).Walk(func(k byte, v inventory.IItem) error {
		equip := v.(*inventory.Equip)
		p.localStat.MaxHP += equip.Model.HP
		p.localStat.MaxMP += equip.Model.MP
		p.localStat.Dex += equip.Model.Dex
		p.localStat.Intt += equip.Model.Intt
		p.localStat.Str += equip.Model.Str
		p.localStat.Luk = equip.Model.Luk
		p.localStat.Magic += equip.Model.MAtk + equip.Model.Intt
		p.localStat.Watk += equip.Model.WAtk
		speed += equip.Model.Speed
		jump += equip.Model.Jump
		return nil
	})

	weapon := p.getInventory(consts.InventoryType.Equipped).GetItem(0xf5)
	if nil != weapon && p.charModel.Job == consts.MapleJobs.Priate {
		// Barefists
		p.localStat.Watk += 8
	}
	if p.localStat.Magic > consts.AbilityLimit.MaxMagic {
		p.localStat.Magic = consts.AbilityLimit.MaxMagic
	}
	// TODO: add buffer status

	if p.localStat.MaxHP > consts.AbilityLimit.MaxHP {
		p.localStat.MaxHP = consts.AbilityLimit.MaxHP
	}
	if p.localStat.MaxMP > consts.AbilityLimit.MaxMP {
		p.localStat.MaxMP = consts.AbilityLimit.MaxMP
	}
	if speed > consts.AbilityLimit.MaxSpeed {
		speed = consts.AbilityLimit.MaxSpeed
	}
	if jump > consts.AbilityLimit.MaxJump {
		jump = consts.AbilityLimit.MaxJump
	}
	p.localStat.SpeedMod = float64(speed) / 100.0
	p.localStat.JumpMod = float64(jump) / 100.0

	if oldMaxHP != 0 &&
		oldMaxHP != p.localStat.MaxHP {
		p.UpdatePartyMemberHP()
	}
}

// TODO:
func (p *Player) UpdatePartyMemberHP() {

}
