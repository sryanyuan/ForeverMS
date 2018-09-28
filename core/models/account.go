package models

import "github.com/juju/errors"

// Account -
type Account struct {
	ID         int64  `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Status     byte   `json:"status"`
	IsAdmin    byte   `json:"isAdmin"`
	IsBanned   int    `json:"isBanned"`
	BanReason  int    `json:"banReason"`
	Gender     byte   `json:"gender"`
	Dob        int64  `json:"dob"`
	CreateTime *int64 `json:"createTime"`
	UpdateTime *int64 `json:"updateTime"`
}

func SelectAccountByUsername(username string) (*Account, error) {
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
