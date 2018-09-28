package loginsvr

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/ngaut/log"
	"github.com/sryanyuan/ForeverMS/core/maplepacket"
	"github.com/sryanyuan/ForeverMS/core/models"
	"github.com/sryanyuan/ForeverMS/core/msconn"
	"github.com/sryanyuan/ForeverMS/core/packets"
)

type packetHandler func(msconn.IConn, *maplepacket.Reader) error

func (s *LoginServer) initPacketDispatchMap() {
	s.packetDispatchMap = map[int]packetHandler{
		maplepacket.RecvReturnToLoginScreen:  s.handleReturnToLoginScreen,
		maplepacket.RecvLoginRequest:         s.handleLoginRequest,
		maplepacket.RecvLoginCheckLogin:      s.handleLoginCheckLogin,
		maplepacket.RecvLoginWorldSelect:     s.handleWorldSelect,
		maplepacket.RecvLoginChannelSelect:   s.handleChannelSelect,
		maplepacket.RecvLoginNameCheck:       s.handleNameCheck,
		maplepacket.RecvLoginNewCharacter:    s.handleNewCharacter,
		maplepacket.RecvLoginDeleteChar:      s.handleDeleteCharacter,
		maplepacket.RecvLoginSelectCharacter: s.handleSelectCharacter,
	}
}

func (s *LoginServer) handleReturnToLoginScreen(conn msconn.IConn, reader *maplepacket.Reader) error {
	conn.Send(packets.LoginReturnFromChannel())
	return nil
}

func (s *LoginServer) handleLoginRequest(conn msconn.IConn, reader *maplepacket.Reader) error {
	// Convert interface
	lConn, ok := conn.(*loginConn)
	if !ok {
		return errors.New("Convert interface failed")
	}

	username := reader.ReadStringInt16()
	password := reader.ReadStringInt16()

	hasher := sha512.New()
	hasher.Write([]byte(password))
	hashedPassword := hex.EncodeToString(hasher.Sum(nil))

	// Get account from mysql
	res := byte(0x00)
	account, err := models.SelectAccountByUsername(username)
	if nil != err {
		log.Warnf("Select account error: %v", err)
		res = 0x05
	} else {
		// Banned = 2, Deleted or Blocked = 3, Invalid Password = 4, Not Registered = 5, Sys Error = 6,
		// Already online = 7, System error = 9, Too many requests = 10, Older than 20 = 11, Master cannot login on this IP = 13
		if hashedPassword != account.Password {
			res = 0x04
		} else if account.IsBanned != 0 {
			res = 0x02
		}
	}

	if res <= 0x01 {
		lConn.admin = account.IsAdmin
		lConn.gender = account.Gender
		lConn.userID = account.ID
		lConn.logined = 1
		// TODO: set global login status in database
	}

	conn.Send(packets.LoginResponce(res, int32(account.ID), account.Gender, account.IsAdmin, account.Username, account.IsBanned))

	return nil
}

func (s *LoginServer) handleLoginCheckLogin(conn msconn.IConn, reader *maplepacket.Reader) error {
	// Convert interface
	lConn, ok := conn.(*loginConn)
	if !ok {
		return errors.New("Convert interface failed")
	}

	account, err := models.SelectAccountByID(lConn.userID)
	if nil != err {
		log.Errorf("Check user error: %v, userID: %v", err, lConn.userID)
	}

	hasher := sha512.New()
	hasher.Write([]byte(account.Username + account.Password)) // should be unique
	hash := hex.EncodeToString(hasher.Sum(nil))
	lConn.sessionHash = hash

	// TODO:
	const maxNumberOfWorlds = 14

	for i := maxNumberOfWorlds; i > -1; i-- {
		conn.Send(packets.LoginWorldListing(byte(i))) // hard coded for now
	}
	conn.Send(packets.LoginEndWorldList())

	return nil
}

func (s *LoginServer) handleWorldSelect(conn msconn.IConn, reader *maplepacket.Reader) error {
	// Convert interface
	lConn, ok := conn.(*loginConn)
	if !ok {
		return errors.New("Convert interface failed")
	}

	worldID := reader.ReadInt16()
	lConn.worldID = worldID

	conn.Send(packets.LoginWorldInfo(0, 0))

	return nil
}

func (s *LoginServer) handleChannelSelect(conn msconn.IConn, reader *maplepacket.Reader) error {
	// Convert interface
	lConn, ok := conn.(*loginConn)
	if !ok {
		return errors.New("Convert interface failed")
	}

	selectWorldID := int16(reader.ReadByte())
	worldID := lConn.worldID
	lConn.channelID = reader.ReadByte()

	var chars []*models.Character
	var err error

	if worldID == selectWorldID {
		// Read characters
		chars, err = models.SelectCharacterByUserIDWorldID(int32(lConn.userID), int32(worldID))
		if nil != err {
			log.Errorf("Get characters error: %v, userID: %v, worldID: %v",
				err, lConn.userID, worldID)
		}
	}

	conn.Send(packets.LoginDisplayCharacters(chars))

	return nil
}

func (s *LoginServer) handleNameCheck(conn msconn.IConn, reader *maplepacket.Reader) error {
	name := reader.ReadStringInt16()
	cnt, err := models.SelectCharacterNameCount(name)
	if nil != err {
		cnt = 1
		log.Errorf("handleNameCheck: error: %v", err)
	}
	conn.Send(packets.LoginNameCheck(name, cnt))
	return nil
}

func (s *LoginServer) handleNewCharacter(conn msconn.IConn, reader *maplepacket.Reader) error {
	// Convert interface
	lConn, ok := conn.(*loginConn)
	if !ok {
		return errors.New("Convert interface failed")
	}

	var newChar models.Character

	newChar.SetName(reader.ReadStringInt16())
	newChar.SetFace(reader.ReadInt32())
	hair := reader.ReadInt32()
	hairColor := reader.ReadInt32()
	newChar.SetHair(hair + hairColor)
	newChar.SetSkin(byte(reader.ReadInt32()))

	top := reader.ReadInt32()
	bottom := reader.ReadInt32()
	shoes := reader.ReadInt32()
	weapon := reader.ReadInt32()

	newChar.SetStr(int16(reader.ReadByte()))
	newChar.SetDex(int16(reader.ReadByte()))
	newChar.SetInt(int16(reader.ReadByte()))
	newChar.SetLuk(int16(reader.ReadByte()))

	allowedEyes := []int32{20000, 20001, 20002, 21000, 21001, 21002, 20100, 20401, 20402, 21700, 21201, 21002}
	allowedHair := []int32{30000, 30020, 30030, 31000, 31040, 31050}
	allowedHairColour := []int32{0, 7, 3, 2}
	allowedBottom := []int32{1060002, 1060006, 1061002, 1061008, 1062115}
	allowedTop := []int32{1040002, 1040006, 1040010, 1041002, 1041006, 1041010, 1041011, 1042167}
	allowedShoes := []int32{1072001, 1072005, 1072037, 1072038, 1072383}
	allowedWeapons := []int32{1302000, 1322005, 1312004, 1442079}
	allowedSkinColour := []int32{0, 1, 2, 3}

	inSlice := func(val int32, s []int32) bool {
		for _, b := range s {
			if b == val {
				return true
			}
		}
		return false
	}

	valid := inSlice(newChar.GetFace(), allowedEyes) && inSlice(hair, allowedHair) && inSlice(hairColor, allowedHairColour) &&
		inSlice(bottom, allowedBottom) && inSlice(top, allowedTop) && inSlice(shoes, allowedShoes) &&
		inSlice(weapon, allowedWeapons) && inSlice(int32(newChar.GetSkin()), allowedSkinColour) //&& (counter == 0)

	if lConn.admin != 0 {
		newChar.SetName("[GM]" + newChar.GetName())
	} else if strings.ContainsAny(newChar.GetName(), "[]") {
		valid = false // hacked client or packet editting
	}

	if valid {
		characterID, err := models.InsertCharacter(&newChar)
		newChar.SetCharID(int32(characterID))

		if err != nil {
			panic(err.Error())
		}

		if lConn.admin != 0 {
			models.InsertCharacterItem(characterID, 1002140, -1) // Hat
			models.InsertCharacterItem(characterID, 1032006, -4) // Earrings
			models.InsertCharacterItem(characterID, 1042003, -5)
			models.InsertCharacterItem(characterID, 1062007, -6)
			models.InsertCharacterItem(characterID, 1072004, -7)
			models.InsertCharacterItem(characterID, 1082002, -8)  // Gloves
			models.InsertCharacterItem(characterID, 1102054, -9)  // Cape
			models.InsertCharacterItem(characterID, 1092008, -10) // Shield
			models.InsertCharacterItem(characterID, 1322013, -11)
		} else {
			models.InsertCharacterItem(characterID, top, -5)
			models.InsertCharacterItem(characterID, bottom, -6)
			models.InsertCharacterItem(characterID, shoes, -7)
			models.InsertCharacterItem(characterID, weapon, -11)
		}

		if err != nil {
			panic(err.Error())
		}

		//characters := character.GetCharacters(conn.GetUserID(), conn.GetWorldID())
		//newCharacter = characters[len(characters)-1]
	}

	conn.Send(packets.LoginCreatedCharacter(valid, &newChar))

	return nil
}

func (s *LoginServer) handleDeleteCharacter(conn msconn.IConn, reader *maplepacket.Reader) error {
	// Convert interface
	lConn, ok := conn.(*loginConn)
	if !ok {
		return errors.New("Convert interface failed")
	}

	dob := reader.ReadInt32()
	charID := reader.ReadInt32()

	deleted := false
	hacking := false

	defer func() {
		conn.Send(packets.LoginDeleteCharacter(charID, deleted, hacking))
	}()

	deleteCharacter, err := models.SelectCharacterByUserID(charID)
	if nil != err {
		log.Errorf("Can't find delete user, character id: %v",
			charID)
		return nil
	}

	if deleteCharacter.GetUserID() != int32(lConn.userID) {
		hacking = true
		return nil
	}

	// Do delete action
	account, err := models.SelectAccountByID(lConn.userID)
	if nil != err {
		log.Errorf("Get account %v error: %v", lConn.userID, err)
		return nil
	}
	if account.Dob == int64(dob) {
		if err = models.DeleteCharacter(int64(charID)); nil != err {
			log.Errorf("Delete character %v error: %v", charID, err)
			return nil
		}
		deleted = true
	}
	return nil
}

func (s *LoginServer) handleSelectCharacter(conn msconn.IConn, reader *maplepacket.Reader) error {
	charID := reader.ReadInt32()
	_, err := models.SelectCharacterByCharacterID(charID)
	if nil != err {
		// Convert interface
		lConn, ok := conn.(*loginConn)
		if !ok {
			return errors.New("Convert interface failed")
		}
		log.Errorf("Can't get character %v of user %v, error: %v", charID, lConn.userID, err)
		return nil
	}

	// TODO: Register channel server to login server
	conn.Send(packets.LoginMigrateClient([]byte{192, 168, 1, 240}, int16(8686), charID))
	return nil
}
