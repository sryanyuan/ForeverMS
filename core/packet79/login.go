package packet79

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sryanyuan/ForeverMS/core/consts"
	"github.com/sryanyuan/ForeverMS/core/consts/opcode"
	"github.com/sryanyuan/ForeverMS/core/maplepacket"
	"github.com/sryanyuan/ForeverMS/core/models"
)

/**
 * Sends a hello packet.
 *
 * @param mapleVersion The maple client version.
 * @param sendIv the IV used by the server for sending
 * @param recvIv the IV used by the server for receiving
 * @param testServer
 * @return
 */
func NewHello(version uint16, ivSend []byte, ivRecv []byte, testServer bool) maplepacket.Packet {
	p := maplepacket.NewPacket()
	// Fixed length: 0x0D
	p.WriteInt16(0x0d)
	p.WriteUint16(version)
	p.WriteBytes(ivRecv)
	p.WriteBytes(ivSend)
	tsV := byte(4)
	if testServer {
		tsV = 5
	}
	p.WriteByte(tsV)
	return p
}

/**
 * Gets a login failed packet.
 *
 * Possible values for <code>reason</code>:<br>
 * 3: ID deleted or blocked<br>
 * 4: Incorrect password<br>
 * 5: Not a registered id<br>
 * 6: System error<br>
 * 7: Already logged in<br>
 * 8: System error<br>
 * 9: System error<br>
 * 10: Cannot process so many connections<br>
 * 11: Only users older than 20 can use this channel<br>
 * 13: Unable to log on as master at this ip<br>
 * 14: Wrong gateway or personal info and weird korean button<br>
 * 15: Processing request with that korean button!<br>
 * 16: Please verify your account through email...<br>
 * 17: Wrong gateway or personal info<br>
 * 21: Please verify your account through email...<br>
 * 23: License agreement<br>
 * 25: Maple Europe notice =[<br>
 * 27: Some weird full client notice, probably for trial versions<br>
 *
 * @param reason The reason logging in failed.
 * @return The login failed packet.
 */
func NewLoginFailed(reason int) maplepacket.Packet {
	p := maplepacket.NewPacketWithOp(opcode.SendOps.LOGIN_STATUS)
	p.WriteInt(reason)
	p.WriteInt16(0)
	return p
}

/**
 * Gets a successful authentication and PIN Request packet.
 *
 * @param account The account name.
 * @return The successful authentication packet.
 */
func NewLoginSuccess(username string, accountId int, gender bool) maplepacket.Packet {
	p := maplepacket.NewPacketWithOp(opcode.SendOps.LOGIN_STATUS)
	p.WriteByte(0)
	p.WriteBool(gender)
	p.WriteInt16(0)
	p.WriteString(username)
	p.WriteBytes([]byte{
		0x00, 0x00, 0x00, 0x03, 0x01, 0x00, 0x00, 0x00, 0xE2, 0xED, 0xA3, 0x7A, 0xFA, 0xC9, 0x01,
	})
	p.WriteInt(0)
	p.WriteInt64(0)
	p.WriteString(strconv.Itoa(accountId))
	p.WriteString(username)
	p.WriteByte(0)
	return p
}

/**
 * Gets a packet detailing a server and its channels.
 *
 * @param serverIndex The index of the server to create information about.
 * @param serverName The name of the server.
 * @param channelLoad Load of the channel - 1200 seems to be max.
 * @return The server info packet.
 */
func NewServerList(serverID byte, serverName string, chInfo []int) maplepacket.Packet {
	p := maplepacket.NewPacketWithOp(opcode.SendOps.SERVERLIST)
	p.WriteByte(serverID)
	p.WriteString(serverName)
	p.WriteByte(0x03) //0: Normal 1: hot 2: very hot 3: new
	p.WriteByte(0x64)
	p.WriteByte(0x00)
	p.WriteByte(0x64)
	p.WriteByte(0x00)

	p.WriteByte(byte(len(chInfo)))
	p.WriteInt(500)

	// Writing loads
	for i, v := range chInfo {
		p.WriteString(fmt.Sprintf("%s-%d", serverName, i+1))
		p.WriteInt(v)
		p.WriteByte(serverID)
		p.WriteInt16(int16(i))
	}
	p.WriteInt16(0)
	return p
}

/**
 * Gets a packet saying that the server list is over.
 *
 * @return The end of server list packet.
 */
func NewEndOfServerList() maplepacket.Packet {
	p := maplepacket.NewPacketWithOp(opcode.SendOps.SERVERLIST)
	p.WriteByte(0xff)
	return p
}

/**
 * Gets a packet with a list of characters.
 *
 * @param c The MapleClient to load characters of.
 * @param serverId The ID of the server requested.
 * @return The character list packet.
 */
func NewCharlist(charlist []*models.Character, maxCharacterLimit int) maplepacket.Packet {
	p := maplepacket.NewPacketWithOp(opcode.SendOps.CHARLIST)
	p.WriteByte(0)
	p.WriteInt(0)
	p.WriteByte(byte(len(charlist)))
	if nil != charlist {
		for _, chr := range charlist {
			addCharEntry(&p, chr)
		}
	}
	p.WriteInt16(3)
	p.WriteInt(maxCharacterLimit)
	return p
}

func addCharEntry(p *maplepacket.Packet, chr *models.Character) {
	addCharStats(p, chr)
	addCharLook(p, chr, false)
	p.WriteByte(0)
	if chr.Job == consts.MapleJobs.GM {
		p.WriteByte(0x02)
	}
}

/**
 * Adds character stats to an existing MaplePacketLittleEndianWriter.
 *
 * @param mplew The MaplePacketLittleEndianWrite instance to write the stats
 * to.
 * @param chr The character to add the stats of.
 */
func addCharStats(p *maplepacket.Packet, chr *models.Character) {
	p.WriteInt(int(chr.ID))
	p.WriteString(chr.Name)
	for i := len(chr.Name); i < 13; i++ {
		p.WriteByte(0)
	}

	genderV := false
	if 0 != chr.Gender {
		genderV = true
	}
	p.WriteBool(genderV)
	p.WriteByte(byte(chr.SkinColor))
	p.WriteInt(chr.Face)
	p.WriteInt(chr.Hair)
	p.WriteInt64(0)
	p.WriteInt64(0)
	p.WriteInt64(0)
	p.WriteByte(byte(chr.Level))
	p.WriteInt16(int16(chr.Job))
	p.WriteInt16(int16(chr.Str))
	p.WriteInt16(int16(chr.Dex))
	p.WriteInt16(int16(chr.Intt))
	p.WriteInt16(int16(chr.Luk))
	p.WriteInt16(int16(chr.HP))
	p.WriteInt16(int16(chr.MaxHP))
	p.WriteInt16(int16(chr.MP))
	p.WriteInt16(int16(chr.MaxMP))
	p.WriteInt16(int16(chr.Ap))
	p.WriteInt16(int16(chr.Sp))
	p.WriteInt(chr.Exp)
	p.WriteInt16(int16(chr.Fame))
	p.WriteInt(0)
	p.WriteInt64(time.Now().UnixNano() / 1e6)
	if 0 == chr.MapID {
		p.WriteInt(10000)
	} else {
		p.WriteInt(chr.MapID)
	}
	p.WriteByte(byte(chr.SpawnPoint))
}

/**
 * Adds the aesthetic aspects of a character to an existing
 * MaplePacketLittleEndianWriter.
 *
 * @param mplew The MaplePacketLittleEndianWrite instance to write the stats
 * to.
 * @param chr The character to add the looks of.
 * @param mega Unknown
 */
func addCharLook(p *maplepacket.Packet, chr *models.Character, mega bool) {
	genderV := false
	if 0 != chr.Gender {
		genderV = true
	}
	p.WriteBool(genderV)
	p.WriteByte(byte(chr.SkinColor))
	p.WriteInt(chr.Face)
	p.WriteBool(!mega)
	p.WriteInt(chr.Hair)
	// TODO: Add equip here

	p.WriteInt(0)
	p.WriteInt64(0)
}

/**
 * Gets a packet detailing a server status message.
 *
 * Possible values for <code>status</code>:<br>
 * 0 - Normal<br>
 * 1 - Highly populated<br>
 * 2 - Full
 *
 * @param status The server status.
 * @return The server status packet.
 */
func NewServerStatus(status int) maplepacket.Packet {
	p := maplepacket.NewPacketWithOp(opcode.SendOps.SERVERSTATUS)
	p.WriteByte(byte(status))
	return p
}

func NewLicenseResult() maplepacket.Packet {
	p := maplepacket.NewPacketWithOp(opcode.SendOps.LICENSE_RESULT)
	p.WriteByte(1)
	return p
}

func NewGenderSet(username, accountID string) maplepacket.Packet {
	p := maplepacket.NewPacketWithOp(opcode.SendOps.GENDER_SET)
	p.WriteByte(0)
	p.WriteString(username)
	p.WriteString(accountID)
	return p
}

func NewLicenseRequest() maplepacket.Packet {
	p := maplepacket.NewPacketWithOp(opcode.SendOps.LOGIN_STATUS)
	p.WriteByte(0x16)
	return p
}

func NewCharNameResponse(charName string, used bool) maplepacket.Packet {
	p := maplepacket.NewPacketWithOp(opcode.SendOps.CHAR_NAME_RESPONSE)
	p.WriteString(charName)
	p.WriteBool(used)
	return p
}

func NewAddNewCharEntry(chr *models.Character, worked bool) maplepacket.Packet {
	p := maplepacket.NewPacketWithOp(opcode.SendOps.ADD_NEW_CHAR_ENTRY)
	p.WriteBool(worked)
	addCharEntry(&p, chr)
	return p
}

const (
	SMNotice              = 0x00
	SMPopup               = 0x01
	SMMegaphone           = 0x02
	SMSuperMegaphone      = 0x03
	SMTopScrollingMessage = 0x04
	SMPinkText            = 0x05
	SMLightBlueText       = 0x06
	SMHeart               = 0x0B
	SMBones               = 0x0C
)

/**
 * Gets a server message packet.
 *
 * Possible values for <code>type</code>:<br>
 * 0: [Notice]<br>
 * 1: Popup<br>
 * 2: Megaphone<br>
 * 3: Super Megaphone<br>
 * 4: Scrolling message at top<br>
 * 5: Pink Text<br>
 * 6: Lightblue Text B: 心脏 C: 白骨
 *
 * @param type The type of the notice.
 * @param channel The channel this notice was sent on.
 * @param message The message to convey.
 * @param servermessage Is this a scrolling ticker?
 * @return The server notice packet.
 */
func NewServerMessage(tp int, channel int, message string, serverMessage bool, megaEar bool) maplepacket.Packet {
	p := maplepacket.NewPacketWithOp(opcode.SendOps.SERVERMESSAGE)
	p.WriteByte(byte(tp))
	if serverMessage {
		p.WriteBool(true)
	}
	p.WriteString(message)
	if tp == SMSuperMegaphone ||
		tp == SMHeart ||
		tp == SMBones {
		p.WriteByte(byte(channel))
		p.WriteBool(megaEar)
	}
	if tp == SMLightBlueText {
		p.WriteInt(0)
	}
	return p
}

/**
 * Gets a server notice packet.
 *
 * Possible values for <code>type</code>:<br>
 * 0: [Notice]<br>
 * 1: Popup<br>
 * 2: Megaphone<br>
 * 3: Super Megaphone<br>
 * 4: Scrolling message at top<br>
 * 5: Pink Text<br>
 * 6: Lightblue Text
 *
 * @param type The type of the notice.
 * @param message The message to convey.
 * @return The server notice packet.
 */
func NewServerNoticeTM(tp int, message string) maplepacket.Packet {
	return NewServerMessage(tp, 0, message, false, false)
}

func NewServerIP(ip []byte, port int16, clientID int) maplepacket.Packet {
	p := maplepacket.NewPacketWithOp(opcode.SendOps.SERVER_IP)
	p.WriteInt16(0)
	p.WriteBytes(ip)
	p.WriteInt16(port)
	p.WriteInt(clientID)
	p.WriteBytes([]byte{1, 0, 0, 0, 0})
	return p
}
