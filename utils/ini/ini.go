package ini

import (
	"gopkg.in/ini.v1"
)

var Config *ini.File
var Section *ini.Section

func Load(config string) error {
	var err error
	Config, err = ini.Load(config)
	if err != nil {
		return err
	}
	Section = Config.Section("")
	return nil
}

func String(key string) string {
	return Section.Key(key).String()
}

func Int(key string) (int, error) {
	return Section.Key(key).Int()
}

func Uint64(key string) (uint64, error) {
	return Section.Key(key).Uint64()
}

func Float64(key string) (float64, error) {
	return Section.Key(key).Float64()
}

func Int64(key string) (int64, error) {
	return Section.Key(key).Int64()
}
