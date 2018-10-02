package mockmodels

import (
	"github.com/sryanyuan/ForeverMS/core/models"
)

func SelectCharacterInventoryItemIDsAndPositionByInventoryType(charID int64, inventoryType int) ([]*models.InventoryItem, error) {
	res := []*models.InventoryItem{
		&models.InventoryItem{},
	}

	return res, nil
}
