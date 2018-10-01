package dataloader

import (
	"strconv"

	"github.com/sryanyuan/ForeverMS/core/wz"
)

func ConvertPathIntDefault(path string, data wz.MapleData, def int) int {
	d := data.ChildByPath(path)
	if nil == d {
		return def
	}
	if d.Type() == wz.STRING {
		v := wz.GetString(d)
		if nil == v {
			return def
		}
		iv, err := strconv.Atoi(*v)
		if nil != err {
			return def
		}
		return iv
	}
	v := wz.GetInt(d)
	if nil == v {
		return def
	}
	return int(*v)
}
