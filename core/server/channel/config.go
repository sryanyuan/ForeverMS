package channel

import "github.com/sryanyuan/ForeverMS/core/models"

type Config struct {
	LogLevel      string                  `toml:"log-level"`
	ListenClients string                  `toml:"listen-clients"`
	DataSource    models.DataSourceConfig `toml:"data-source"`
	SendOps       string                  `toml:"send-ops"`
	RecvOps       string                  `toml:"recv-ops"`
	TestServer    bool                    `toml:"test-server"`
}
