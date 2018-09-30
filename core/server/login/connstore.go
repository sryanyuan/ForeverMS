package login

import (
	"github.com/sryanyuan/ForeverMS/core/netio"
)

const (
	connKeyGender  = "gender"
	connKeyAdmin   = "admin"
	connKeyUserID  = "userID"
	connKeyLogined = "logined"
)

func storeConnGender(conn netio.IConn, v int) {
	conn.SetInt(connKeyGender, v)
}

func loadConnGender(conn netio.IConn) int {
	return conn.GetInt(connKeyGender)
}

func storeConnAdmin(conn netio.IConn, v int) {
	conn.SetInt(connKeyAdmin, v)
}

func loadConnAdmin(conn netio.IConn) int {
	return conn.GetInt(connKeyAdmin)
}

func storeConnUserID(conn netio.IConn, v int) {
	conn.SetInt(connKeyUserID, v)
}

func loadConnUserID(conn netio.IConn) int {
	return conn.GetInt(connKeyUserID)
}

func storeConnLogined(conn netio.IConn, v int) {
	conn.SetInt(connKeyLogined, v)
}

func loadConnLogined(conn netio.IConn) int {
	return conn.GetInt(connKeyLogined)
}
