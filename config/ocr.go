package config

import (
	"time"
)

type TTOCR struct {
	Key      string        `json:"key"      yaml:"key"`
	Interval time.Duration `json:"interval" yaml:"interval"`
	Timeout  time.Duration `json:"timeout"  yaml:"timeout"`
	// https://www.kancloud.cn/ttorc/ttorc/3119237
	ItemId string `json:"item_id"  yaml:"item_id"`
}

func (c Config) TT() TTOCR {
	if c.TTOCR.Interval < time.Second {
		c.TTOCR.Interval = 3 * time.Second
	}
	if c.TTOCR.Timeout < 60*time.Second {
		c.TTOCR.Timeout = 90 * time.Second
	}
	if c.TTOCR.ItemId == "" {
		c.TTOCR.ItemId = "388"
	}
	return c.TTOCR
}

func TT() TTOCR {
	return C().TT()
}
