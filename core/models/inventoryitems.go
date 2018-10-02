package models

import "github.com/juju/errors"

/*
-- ----------------------------
-- Table structure for `inventoryitems`
-- ----------------------------
DROP TABLE IF EXISTS `inventoryitems`;
CREATE TABLE `inventoryitems` (
  `inventoryitemid` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `characterid` int(11) DEFAULT NULL,
  `storageid` int(10) unsigned DEFAULT NULL,
  `itemid` int(11) NOT NULL DEFAULT '0',
  `inventorytype` int(11) NOT NULL DEFAULT '0',
  `position` int(11) NOT NULL DEFAULT '0',
  `quantity` int(11) NOT NULL DEFAULT '0',
  `owner` tinytext NOT NULL,
  `petid` int(11) NOT NULL DEFAULT '-1',
  `expiration` bigint(20) unsigned DEFAULT '0',
  PRIMARY KEY (`inventoryitemid`),
  KEY `inventoryitems_ibfk_1` (`characterid`),
  KEY `characterid` (`characterid`),
  KEY `inventorytype` (`inventorytype`),
  KEY `storageid` (`storageid`),
  KEY `characterid_2` (`characterid`,`inventorytype`),
  CONSTRAINT `inventoryitems_ibfk_1` FOREIGN KEY (`characterid`) REFERENCES `characters` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
*/

type InventoryItem struct {
	InventoryItemID int64
	CharacterID     int64
	StorageID       int
	ItemID          int
	InventoryType   int
	Position        int
	Quantity        int
	Owner           string
	Expiration      int64
	UniqueID        int64
	PetSlot         int
}

type CharacterInventoryItem struct {
	InventoryItem
	InventoryEquipment
}

func SelectCharacterInventoryItemIDsAndPositionByInventoryType(charID int64, inventoryType int) ([]*InventoryItem, error) {
	rows, err := GetGlobalDB().Query(`SELECT
	itemid,
	position,
	FROM
	inventoryitems WHERE
	characterid = ? AND
	inventorytype = ?`,
		charID, inventoryType)
	if nil != err {
		return nil, errors.Trace(err)
	}
	defer rows.Close()

	res := make([]*InventoryItem, 0, 32)
	for rows.Next() {
		var item InventoryItem
		if err = rows.Scan(&item); nil != err {
			return nil, errors.Trace(err)
		}
		res = append(res, &item)
	}

	return res, nil
}

func SelectCharacterInventoryItem(charID int64) ([]*CharacterInventoryItem, error) {
	rows, err := GetGlobalDB().Query(`SELECT
	inventoryitems.inventoryitemid,
	inventoryitems.storageid,
	inventoryitems.itemid,
	inventoryitems.inventorytype,
	inventoryitems.position,
	inventoryitems.quantity,
	inventoryitems.owner,
	inventoryitems.petid,
	inventoryitems.expiration,
	inventoryitems.uniqueid,
	inventoryequipment.inventoryequipmentid,
	inventoryequipment.inventoryitemid,
	inventoryequipment.upgradeslots,
	inventoryequipment.level,
	inventoryequipment.str,
	inventoryequipment.dex,
	inventoryequipment.intt,
	inventoryequipment.luk,
	inventoryequipment.hp,
	inventoryequipment.mp,
	inventoryequipment.watk,
	inventoryequipment.matk,
	inventoryequipment.wdef,
	inventoryequipment.mdef,
	inventoryequipment.acc,
	inventoryequipment.avoid,
	inventoryequipment.hands,
	inventoryequipment.speed,
	inventoryequipment.jump,
	inventoryequipment.ringid,
	inventoryequipment.locked,
	inventoryequipment.vicious,
	inventoryequipment.flag
	FROM inventoryitems LEFT JOIN inventoryequipment USING (inventoryitemid) WHERE characterid = ?`,
		charID)
	if nil != err {
		return nil, errors.Trace(err)
	}
	defer rows.Close()

	res := make([]*CharacterInventoryItem, 0, 32)
	for rows.Next() {
		var item CharacterInventoryItem
		if err = rows.Scan(
			&item.InventoryItem.InventoryItemID,
			&item.InventoryItem.CharacterID,
			&item.InventoryItem.StorageID,
			&item.InventoryItem.ItemID,
			&item.InventoryItem.InventoryType,
			&item.InventoryItem.Position,
			&item.InventoryItem.Quantity,
			&item.InventoryItem.Owner,
			&item.InventoryItem.Expiration,
			&item.InventoryItem.UniqueID,
			&item.InventoryEquipment.InventoryEquipmentID,
			&item.InventoryEquipment.InventoryItemID,
			&item.InventoryEquipment.UpgradeSlots,
			&item.InventoryEquipment.Level,
			&item.InventoryEquipment.Str,
			&item.InventoryEquipment.Dex,
			&item.InventoryEquipment.Intt,
			&item.InventoryEquipment.Luk,
			&item.InventoryEquipment.HP,
			&item.InventoryEquipment.MP,
			&item.InventoryEquipment.WAtk,
			&item.InventoryEquipment.MAtk,
			&item.InventoryEquipment.WDef,
			&item.InventoryEquipment.MDef,
			&item.InventoryEquipment.Acc,
			&item.InventoryEquipment.Avoid,
			&item.InventoryEquipment.Hands,
			&item.InventoryEquipment.Speed,
			&item.InventoryEquipment.Jump,
			&item.InventoryEquipment.RingID,
			&item.InventoryEquipment.Locked,
			&item.InventoryEquipment.Vicious,
			&item.InventoryEquipment.Flag,
		); nil != err {
			return nil, errors.Trace(err)
		}
		item.InventoryItem.CharacterID = charID
		res = append(res, &item)
	}
	return res, nil
}
