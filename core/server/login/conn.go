package login

import (
	"github.com/sryanyuan/ForeverMS/core/netio"
)

const (
	loginStatusNone = iota
	loginStatusServerMigrate
	loginStatusLoggedIn
	loginStatusWaitting
	loginStatusEnteringPin
	loginStatusPinCorrect
	loginStatusViewAllChar
)

type loginConn struct {
	*netio.Conn
	// Private fields
	username string

	gender      byte
	admin       byte
	userID      int64
	logined     int8
	sessionHash string

	isGM      bool
	accountID int

	worldID   int
	channelID int

	loginStatus int

	lastPongTm int64
}
