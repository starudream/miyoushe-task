package config

import (
	"fmt"

	"github.com/starudream/miyoushe-task/util"
)

type Device struct {
	Id string `json:"id,omitempty" yaml:"id,omitempty"`
	// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/other/authentication.md#x-rpc-client_type
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
	// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/other/authentication.md#x-rpc-device_name
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/other/authentication.md#x-rpc-device_model
	Model string `json:"model,omitempty" yaml:"model,omitempty"`
	// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/other/authentication.md#x-rpc-sys_version
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
	// https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/other/authentication.md#x-rpc-channel
	Channel string `json:"channel,omitempty" yaml:"channel,omitempty"`
}

func (d Device) TableCellString() string {
	return fmt.Sprintf("%s (%s)", d.Id, d.Name)
}

var _d = Device{
	Id:      "",
	Type:    "2",
	Name:    "Xiaomi 22011211C",
	Model:   "22011211C",
	Version: "13",
	Channel: "miyousheluodi",
}

func NewDevice() Device {
	d := _d
	d.Id = util.UUID()
	return d
}

// Headers https://github.com/UIGF-org/mihoyo-api-collect/blob/3a9116ea538941cfead749572df1f364cb9f9c8d/other/authentication.md#%E8%AF%B7%E6%B1%82%E5%A4%B4
func (d Device) Headers() map[string]string {
	return map[string]string{
		"x-rpc-device_id":    def(d.Id, util.UUID()),
		"x-rpc-client_type":  def(d.Type, _d.Type),
		"x-rpc-device_name":  def(d.Name, _d.Name),
		"x-rpc-device_model": def(d.Model, _d.Model),
		"x-rpc-sys_version":  def(d.Version, _d.Version),
		"x-rpc-channel":      def(d.Channel, _d.Channel),
	}
}

func def(v string, def ...string) string {
	if v != "" {
		return v
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}
