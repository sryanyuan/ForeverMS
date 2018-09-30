package login

import (
	"crypto/sha512"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/juju/errors"
	"github.com/ngaut/log"
	"github.com/sryanyuan/ForeverMS/core/maplepacket"
	"github.com/sryanyuan/ForeverMS/core/models"
	"github.com/sryanyuan/ForeverMS/core/netio"
	"github.com/sryanyuan/ForeverMS/core/packet79"
)

const (
	LoginResSuccess           = 0
	LoginResBlocked           = 3
	LoginResIncorrectPassword = 4
	LoginResUnFind            = 5
	LoginResBanned            = 6
	LoginResIsLogged          = 7
	LoginResServerBusy        = 10
	LoginResGenderNeeded      = 11
)

// TODO: Using map to accelerate the interface conversion
func must2loginConn(conn netio.IConn) *loginConn {
	return conn.(*loginConn)
}

func (s *LoginServer) handleDummy(conn netio.IConn, reader *maplepacket.Reader) error {
	return nil
}

func (s *LoginServer) handlePong(conn netio.IConn, reader *maplepacket.Reader) error {
	lc := must2loginConn(conn)
	lc.lastPongTm = time.Now().UnixNano() / 1e6
	return nil
}

func (s *LoginServer) handleLoginPassword(conn netio.IConn, reader *maplepacket.Reader) error {
	lc := must2loginConn(conn)
	if lc.loginStatus == loginStatusLoggedIn {
		return errors.Errorf("Client %s already logged in", conn.GetRemoteAddress())
	}

	username := reader.ReadStringInt16()
	password := reader.ReadStringInt16()

	var err error
	logRes := LoginResUnFind

	var account *models.Account
	account, err = models.SelectAccountForLogin(username)
	if nil != err {
		return errors.Trace(err)
	}

	defer func() {
		if nil != err {
			return
		}
		// Send response to client
		var p maplepacket.Packet
		if logRes == LoginResSuccess {
			p = packet79.NewLoginSuccess(account.Name, int(account.ID), account.Gender != 0)
		} else {
			p = packet79.NewLoginFailed(logRes)
		}
		conn.Send(p)

		// Send serverlist if success
		if logRes != LoginResSuccess {
			return
		}
		// TODO: Get server list from channel servers
		channelServers := []int{1}
		conn.Send(packet79.NewServerList(byte(s.config.ServerID), s.config.ServerName, channelServers))
		conn.Send(packet79.NewEndOfServerList())
	}()

	lc.username = username
	// TODO: Check banned or not

	// Check password
	hasher := sha512.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum([]byte(account.Salt)))

	if account.Password != hashedPassword {
		logRes = LoginResIncorrectPassword
		return nil
	}

	if account.Banned != 0 {
		logRes = LoginResBlocked
		return nil
	}
	if account.LoggedIn != 0 {
		logRes = LoginResIsLogged
		return nil
	}

	// Login ok
	lc.loginStatus = loginStatusLoggedIn
	lc.accountID = int(account.ID)
	lc.gender = account.Gender
	lc.isGM = account.GM != 0
	models.UpdateAccountSetLoggedIn(username, true)

	return nil
}

func (s *LoginServer) handleServerListRequest(conn netio.IConn, reader *maplepacket.Reader) error {
	// TODO: Get server list from channel servers
	channelServers := []int{1}
	conn.Send(packet79.NewServerList(byte(s.config.ServerID), s.config.ServerName, channelServers))
	conn.Send(packet79.NewEndOfServerList())
	return nil
}

func (s *LoginServer) handleCharlistRequest(conn netio.IConn, reader *maplepacket.Reader) error {
	lc := must2loginConn(conn)
	if lc.loginStatus != loginStatusLoggedIn {
		log.Warnf("Client %s request charlist without login", lc.GetRemoteAddress())
		return nil
	}

	worldID := int(reader.ReadByte())
	channelID := int(reader.ReadByte())
	lc.worldID = worldID
	lc.channelID = channelID
	log.Debugf("User %s request charlist with worldID %d, channelID %d",
		worldID, channelID)
	// Get characters from DB
	chars, err := models.SelectCharacterByAccountIDWorldID(int32(lc.accountID), int32(worldID))
	if nil != err {
		return errors.Trace(err)
	}
	conn.Send(packet79.NewCharlist(chars, s.config.MaxCharactersLimit))

	return nil
}

func (s *LoginServer) handleServerStatusRequest(conn netio.IConn, reader *maplepacket.Reader) error {
	// TODO: Read all channel server status
	conn.Send(packet79.NewServerStatus(0))
	return nil
}

func (s *LoginServer) handleLicenseRequest(conn netio.IConn, reader *maplepacket.Reader) error {
	if 1 == reader.ReadByte() {
		conn.Send(packet79.NewLicenseResult())
		lc := must2loginConn(conn)
		lc.loginStatus = loginStatusNone
	} else {
		return errors.Errorf("Client %s not accept the license", conn.GetRemoteAddress())
	}
	return nil
}

func (s *LoginServer) handleSetGender(conn netio.IConn, reader *maplepacket.Reader) error {
	lc := must2loginConn(conn)

	gender := reader.ReadByte()
	username := reader.ReadStringInt16()
	if lc.username != username {
		return errors.Errorf("Username not equal when set gender, request username: %v, account: %v",
			username, lc.username)
	}
	if err := models.UpdateAccountSetGender(lc.username, int(gender)); nil != err {
		return errors.Trace(err)
	}
	lc.gender = gender
	// Request for license again
	conn.Send(packet79.NewGenderSet(username, strconv.Itoa(lc.accountID)))
	conn.Send(packet79.NewLicenseRequest())

	return nil
}

func (s *LoginServer) handleCheckCharName(conn netio.IConn, reader *maplepacket.Reader) error {
	charName := reader.ReadStringInt16()
	count, err := models.SelectCharacterNameCount(charName)
	if nil != err {
		return errors.Trace(err)
	}
	conn.Send(packet79.NewCharNameResponse(charName, count != 0))
	return nil
}

func (s *LoginServer) handleCreateChar(conn netio.IConn, reader *maplepacket.Reader) error {
	lc := must2loginConn(conn)

	name := reader.ReadStringInt16()
	job := reader.ReadInt()
	face := reader.ReadInt()
	hair := reader.ReadInt()
	hairColor := 0
	var skinColor byte

	if job == 0 {
		skinColor = 10
	} else if job == 2 {
		skinColor = 11
	} else {
		skinColor = 0
	}

	top := reader.ReadInt()
	bottom := reader.ReadInt()
	shoes := reader.ReadInt()
	weapon := reader.ReadInt()
	_ = top
	_ = bottom
	_ = shoes
	_ = weapon

	var newChar models.Character
	if lc.isGM {
		newChar.GM = 1
	}
	newChar.World = int32(lc.worldID)
	newChar.Face = face
	newChar.Hair = hair + hairColor
	newChar.Gender = int(lc.gender)

	if job == 2 {
		newChar.Str = 11
		newChar.Dex = 6
		newChar.Intt = 4
		newChar.Luk = 4
		newChar.Ap = 0
	} else {
		newChar.Str = 4
		newChar.Dex = 4
		newChar.Intt = 4
		newChar.Luk = 4
		newChar.Ap = 9
	}
	newChar.Name = name
	newChar.SkinColor = int(skinColor)

	if job == 1 {
		newChar.Job = 0
	} else if job == 0 {
		newChar.Job = 1000
	} else if job == 2 {
		newChar.Job = 2000
	}

	if _, err := models.InsertCharacter(&newChar); nil != err {
		return errors.Trace(err)
	}
	// Send success response
	conn.Send(packet79.NewAddNewCharEntry(&newChar, true))

	return nil
}

func (s *LoginServer) handleCharSelect(conn netio.IConn, reader *maplepacket.Reader) error {
	lc := must2loginConn(conn)
	if lc.loginStatus != loginStatusLoggedIn {
		return nil
	}

	// TODO: Dynamic send channel server address
	channelServerIP := []byte{127, 0, 0, 1}
	channelServerPort := int16(7575)

	charID := reader.ReadInt()
	lc.loginStatus = loginStatusServerMigrate
	conn.Send(packet79.NewServerIP(channelServerIP, channelServerPort, charID))
	return nil
}

func (s *LoginServer) handleErrorLog(conn netio.IConn, reader *maplepacket.Reader) error {
	log.Errorf("Client %s error: %s",
		conn.GetRemoteAddress(), reader.ReadStringInt16())
	return nil
}

func (s *LoginServer) handlePlayerUpdate(conn netio.IConn, reader *maplepacket.Reader) error {
	// TODO: Save player data here
	return nil
}
