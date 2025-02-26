package configs

import (
	_ "embed"
	"fmt"
	"log/slog"
	"os"

	"github.com/BurntSushi/toml"
)

//go:embed default.toml
var defaultCfg string

type Dbcfg struct {
	Username string
	Password string
	Host     string
	Port     uint16
	Dbname   string
}

type http struct {
	Port   uint16
	Expose bool
}

type security struct {
	JwtSecret string `toml:"jwt-secret"`
}

type Config struct {
	Db       Dbcfg `toml:"database"`
	Http     http
	Security security
}

func LoadConfig() (Config, error) {
	var self Config
	var dat string

	f := "/etc/local_run/config.toml"

	if _, err := os.Stat(f); err != nil {
		slog.Info(fmt.Sprintf("The file %s doesn't exist, using builtin config instead", f))
		dat = defaultCfg
	} else {
		content, err := os.ReadFile(f)
		if err != nil {
			return Config{}, err
		}

		dat = string(content)
	}

	_, err := toml.Decode(dat, &self)
	if err != nil {
		return Config{}, err
	}

	return self, nil
}
