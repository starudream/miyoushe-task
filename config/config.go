package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/starudream/go-lib/core/v2/codec/yaml"
	"github.com/starudream/go-lib/core/v2/config"
)

type Config struct {
	Accounts []Account `json:"accounts" yaml:"accounts"`
	Cron     Cron      `json:"cron"     yaml:"cron"`
}

type Cron struct {
	Spec    string `json:"spec"    yaml:"spec"`
	Startup bool   `json:"startup" yaml:"startup"`
}

var (
	_c = Config{
		Cron: Cron{
			Spec:    "0 0 8 * * *",
			Startup: false,
		},
	}
	_cMu = sync.Mutex{}
)

func init() {
	_ = config.Unmarshal("", &_c)
	_ = config.Unmarshal("cron", &_c.Cron)
	config.LoadStruct(_c)
}

func C() Config {
	_cMu.Lock()
	defer _cMu.Unlock()
	return _c
}

func Save() error {
	config.LoadStruct(_c)

	bs, err := yaml.Marshal(config.Raw())
	if err != nil {
		return fmt.Errorf("marshal config error: %w", err)
	}

	err = os.WriteFile(config.LoadedFile(), bs, 0644)
	if err != nil {
		return fmt.Errorf("write config file error: %w", err)
	}

	return nil
}
