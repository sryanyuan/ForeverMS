package loginsvr

import (
	"github.com/BurntSushi/toml"
	"github.com/juju/errors"
	"github.com/sryanyuan/ForeverMS/core/models"
)

type Config struct {
	LogLevel      string                  `toml:"log-level"`
	ListenClients string                  `toml:"listen-clients"`
	DataSource    models.DataSourceConfig `toml:"data-source"`
}

func (c *Config) LoadFromFile(f string) error {
	_, err := toml.DecodeFile(f, c)
	return errors.Trace(err)
}
