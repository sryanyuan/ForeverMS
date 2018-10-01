package mockmodels

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/sryanyuan/ForeverMS/core/models"
)

func SelectAccountForLogin(username string) (*models.Account, error) {
	account := &models.Account{}
	account.ID = 1
	account.Name = username

	hasher := sha512.New()
	hasher.Write([]byte("111111"))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))
	account.Password = hashedPassword

	return account, nil
}

func UpdateAccountSetLoggedIn(username string, loggedIn bool) error {
	return nil
}
