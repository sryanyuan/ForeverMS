package models

import (
	"time"

	"github.com/juju/errors"
	"github.com/sryanyuan/ForeverMS/core/consts"
)

/*
-- ----------------------------
-- Table structure for `accounts`
-- ----------------------------
DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(13) NOT NULL DEFAULT '',
  `password` varchar(128) NOT NULL DEFAULT '',
  `salt` varchar(32) DEFAULT NULL,
  `loggedin` tinyint(4) NOT NULL DEFAULT '0',
  `lastlogin` TIMESTAMP NULL DEFAULT NULL,
  `createdat` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `birthday` date  NULL DEFAULT '0000-00-00',
  `banned` tinyint(1) NOT NULL DEFAULT '0',
  `banreason` text,
  `gm` tinyint(1) NOT NULL DEFAULT '0',
  `email` tinytext,
  `emailcode` varchar(40) DEFAULT NULL,
  `forumaccid` int(11) NOT NULL DEFAULT '0',
  `macs` tinytext,
  `lastpwemail` TIMESTAMP NOT NULL DEFAULT '2002-12-31 17:00:00',
  `tempban` TIMESTAMP NULL DEFAULT '0000-00-00 00:00:00',
  `greason` tinyint(4) DEFAULT NULL,
  `paypalNX` int(11) DEFAULT NULL,
  `mPoints` int(11) DEFAULT NULL,
  `cardNX` int(11) DEFAULT NULL,
  `donatorPoints` tinyint(1) DEFAULT NULL,
  `guest` tinyint(1) NOT NULL DEFAULT '0',
  `LastLoginInMilliseconds` bigint(20) unsigned NOT NULL DEFAULT '0',
  `LeaderPoints` tinyint(1) DEFAULT NULL,
  `pqPoints` tinyint(1) DEFAULT NULL,
  `lastknownip` tinytext NOT NULL,
  `pin` varchar(13) DEFAULT NULL,
  `NomePessoal` varchar(40) NOT NULL,
  `fb` varchar(40) NOT NULL,
  `twt` varchar(40) NOT NULL,
  `BetaPoints` int(11) DEFAULT NULL,
  `sitelogged` text,
  `webadmin` int(1) DEFAULT '0',
  `nick` varchar(20) DEFAULT NULL,
  `mute` int(1) DEFAULT '0',
  `ip` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `forumaccid` (`forumaccid`),
  KEY `ranking1` (`id`,`banned`,`gm`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;
*/

// Account -
type Account struct {
	ID                      int64  `json:"id"`
	Name                    string `json:"name"`
	Password                string `json:"password"`
	Salt                    string
	LoggedIn                byte `json:"status"`
	LastLogin               *int64
	CreateDat               *int64
	Birthday                *int64
	Banned                  byte
	BanReason               string
	GM                      byte
	Email                   string
	EmailCode               string
	ForumAccid              int
	Macs                    string
	LastPweMail             *int64
	TempBan                 *int64
	Greason                 byte
	PaypalNX                int
	MPoints                 int
	CardNX                  int
	DonatorPoints           byte
	Guest                   byte
	LastLoginInMilliSeconds int64
	LeaderPoints            byte
	PQPoints                byte
	LastKnownIP             string
	Pin                     string
	NomePessoal             string
	Fb                      string
	Twt                     string
	BetaPoints              int
	SiteLogged              string
	WebAdmin                int
	Nick                    string
	Mute                    int
	IP                      string
	Gender                  byte
}

func SelectAccountForLogin(username string) (*Account, error) {
	row := GetGlobalDB().QueryRow(`SELECT
	id,
	password,
	salt,
	createdat,
	tempban,
	banned,
	GM,
	greason,
	pin,
	gender
	FROM accounts WHERE name = ?`,
		username)

	var account Account
	if err := row.Scan(
		&account.ID,
		&account.Password,
		&account.Salt,
		&account.CreateDat,
		&account.TempBan,
		&account.Banned,
		&account.GM,
		&account.Greason,
		&account.Pin,
		&account.Gender,
	); nil != err {
		return nil, errors.Trace(err)
	}
	return &account, nil
}

func UpdateAccountSetLoggedIn(username string, logged bool) error {
	lflag := 0
	if logged {
		lflag = 1
	}
	_, err := GetGlobalDB().Exec(`UPDATE accounts SET
	loggedin = ? AND
	lastlogin = ?
	WHERE username = ?`,
		lflag, time.Now().Format(consts.TimeFormat), username)
	return errors.Trace(err)
}

func UpdateAccountSetGender(username string, gender int) error {
	_, err := GetGlobalDB().Exec(`UPDATE accounts SET
	gender = ?
	WHERE username = ?`,
		gender, username)
	return errors.Trace(err)
}

/*func SelectAccountByUsername(username string) (*Account, error) {
	row, err := GetGlobalDB().Query(`SELECT id,
	password,
	status,
	isAdmin,
	isBanned,
	banReason,
	gender,
	dob,
	unix_timestamp(createTime),
	unix_timestamp(updateTime)
	FROM accounts WHERE username = ?`,
		username)
	if nil != err {
		return nil, errors.Trace(err)
	}
	var account Account
	if err = row.Scan(
		&account.ID,
		&account.Password,
		&account.Status,
		&account.IsAdmin,
		&account.IsBanned,
		&account.BanReason,
		&account.Gender,
		&account.Dob,
		&account.CreateTime,
		&account.UpdateTime,
	); nil != err {
		return nil, errors.Trace(err)
	}
	account.Username = username
	return &account, nil
}

func SelectAccountByID(id int64) (*Account, error) {
	row, err := GetGlobalDB().Query(`SELECT username,
	password,
	status,
	isAdmin,
	isBanned,
	banReason,
	gender,
	dob,
	unix_timestamp(createTime),
	unix_timestamp(updateTime)
	FROM accounts WHERE id = ?`,
		id)
	if nil != err {
		return nil, errors.Trace(err)
	}
	var account Account
	if err = row.Scan(
		&account.Username,
		&account.Password,
		&account.Status,
		&account.IsAdmin,
		&account.IsBanned,
		&account.BanReason,
		&account.Gender,
		&account.Dob,
		&account.CreateTime,
		&account.UpdateTime,
	); nil != err {
		return nil, errors.Trace(err)
	}
	account.ID = id
	return &account, nil
}
*/
