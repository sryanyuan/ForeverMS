package mockmodels

import (
	"fmt"

	"github.com/sryanyuan/ForeverMS/core/models"
)

var (
	charID int64
)

func init() {
	charID = 1
}

func SelectCharacterByAccountIDWorldID(accountID int32, worldID int32) ([]*models.Character, error) {
	res := make([]*models.Character, 0, 3)
	for i := 0; i < 1; i++ {
		var c models.Character
		c.ID = int64(i + 1)
		c.AccountID = 1
		c.World = worldID
		c.Name = fmt.Sprintf("Heart_%d", i+1)
		c.Level = 10
		c.Str = 10
		c.Dex = 10
		c.Luk = 10
		c.Intt = 10
		c.HP = 100
		c.MP = 100
		c.MaxHP = 100
		c.MaxMP = 100
		c.Job = 0
		c.BuddyCapacity = 25
		c.Face = 20001
		c.Hair = 30000

		res = append(res, &c)
	}
	return res, nil
}

func InsertCharacter(ch *models.Character) (int64, error) {
	charID++
	return charID, nil
}

func SelectCharacterNameCount(name string) (int, error) {
	return 0, nil
}
