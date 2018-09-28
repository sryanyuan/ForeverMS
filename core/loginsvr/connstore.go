package loginsvr

import (
	"github.com/sryanyuan/ForeverMS/core/msconn"
)

const (
	connKeyGender  = "gender"
	connKeyAdmin   = "admin"
	connKeyUserID  = "userID"
	connKeyLogined = "logined"
)

func storeConnGender(conn msconn.IConn, v int) {
	conn.SetInt(connKeyGender, v)
}

func loadConnGender(conn msconn.IConn) int {
	return conn.GetInt(connKeyGender)
}

func storeConnAdmin(conn msconn.IConn, v int) {
	conn.SetInt(connKeyAdmin, v)
}

func loadConnAdmin(conn msconn.IConn) int {
	return conn.GetInt(connKeyAdmin)
}

func storeConnUserID(conn msconn.IConn, v int) {
	conn.SetInt(connKeyUserID, v)
}

func loadConnUserID(conn msconn.IConn) int {
	return conn.GetInt(connKeyUserID)
}

func storeConnLogined(conn msconn.IConn, v int) {
	conn.SetInt(connKeyLogined, v)
}

func loadConnLogined(conn msconn.IConn) int {
	return conn.GetInt(connKeyLogined)
}
