package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/starudream/go-lib/core/v2/codec/yaml"
	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

type Config struct {
	Accounts []Account `json:"accounts" yaml:"accounts"`
	Cron     Cron      `json:"cron"     yaml:"cron"`

	// 打码接口
	RROCRKey string `json:"rrocr.key" yaml:"rrocr.key"`
	TTOCR    TTOCR  `json:"ttocr"     yaml:"ttocr"`
}

func (c Config) FirstAccount() Account {
	if len(c.Accounts) == 0 {
		osutil.PanicErr(fmt.Errorf("no account found"))
	}
	return c.Accounts[0]
}

type Cron struct {
	Spec    string `json:"spec"    yaml:"spec"`
	Startup bool   `json:"startup" yaml:"startup"`
}

var (
	_c = Config{
		Cron: Cron{
			Spec:    "5 4 8 * * *",
			Startup: false,
		},
	}
	_cMu = sync.Mutex{}
)

func init() {
	_ = config.Unmarshal("", &_c)
	_ = config.Unmarshal("cron", &_c.Cron)
	_ = config.Unmarshal("ttocr", &_c.TTOCR)
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

	filename := config.LoadedFile()
	if filename == "" {
		filename = filepath.Join(osutil.ExeDir(), osutil.ExeName()+".yaml")
		slog.Info("config file not found, save to default file", slog.String("file", filename))
	}

	err = os.WriteFile(config.LoadedFile(), bs, 0644)
	if err != nil {
		return fmt.Errorf("write config file error: %w", err)
	}

	slog.Info("save config success", slog.String("file", filename))

	return nil
}
