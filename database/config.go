package database

import "AkashaPanel/objects"

func GetConfig(key string) string {
	conf := &objects.Config{Key: key}
	db.Find(conf)
	return conf.Value
}

func SetConfig(key, value string) {
	conf := &objects.Config{Key: key, Value: value}
	db.Create(conf)
}

func InitConfig(key, value string) {
	if GetConfig(key) != "" {
		return
	}
	SetConfig(key, value)
}
