package channel

import (
	"time"

	"github.com/sryanyuan/ForeverMS/core/maplepacket"
	"github.com/sryanyuan/ForeverMS/core/netio"
)

const (
	_ = iota
	loginStatusLoggedIn
)

// TODO: Using map to accelerate the interface conversion
func must2chanConn(conn netio.IConn) *chanConn {
	return conn.(*chanConn)
}

func (s *ChannelServer) handlePong(conn netio.IConn, reader *maplepacket.Reader) error {
	cc := must2chanConn(conn)
	cc.lastPongTime = time.Now().Unix()
	return nil
}

func (s *ChannelServer) handlePlayerLoggedIn(conn netio.IConn, reader *maplepacket.Reader) error {
	/*cc := must2chanConn(conn)

	charID := reader.ReadInt()
	// TODO: Read connection status from redis/Mysql, current just pass
	if cc.loginStatus == loginStatusLoggedIn {
		return nil
	}*/

	// Read character data from DB
	//var err error

	// TODO: Set the loggedIn channel status

	// Add player into world
	/*prevPlayer := s.world.GetPlayer(int64(charID))
	if nil != prevPlayer {
		// Already logined?
		return errors.New("Already logged in")
	}
	player := &MSPlayer{
		chanConn: cc,
	}
	if err = player.LoadCharacterData(charID); nil != err {
		return errors.Trace(err)
	}
	s.world.AddPlayer(player)
	log.Infof("Player [%s] logged in", player.charModel.Name)
	// Send character information to player
	*/
	return nil
}
