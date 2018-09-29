package opcode

import (
	"github.com/juju/errors"
	"github.com/magiconair/properties"
)

func LoadRecvOpsFromFile(f string) error {
	return errors.Trace(loadProperties(f, &RecvOps))
}

func LoadSendOpsFromFile(f string) error {
	return errors.Trace(loadProperties(f, &SendOps))
}

func loadProperties(f string, v interface{}) error {
	p, err := properties.LoadFile(f, properties.UTF8)
	if nil != err {
		return errors.Trace(err)
	}
	if err = p.Decode(v); nil != err {
		return errors.Trace(err)
	}
	return nil
}
