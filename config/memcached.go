package config

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

var Memcache *memcache.Client

func ConnectToMemcache() (err error) {
	servers := make([]string, AppEnv.MemcacheServersCount)
	for i := 0; i < AppEnv.MemcacheServersCount; i++ {
		servers[i] = fmt.Sprintf("%s%d:%d", AppEnv.MemcacheHostRange, i+1, AppEnv.MemcacheInitialPort+i)
	}

	Memcache = memcache.New(servers...)
	err = Memcache.Ping()
	return
}

func GetMemcacheVal(key string) (val string, err error) {
	var item *memcache.Item
	item, err = Memcache.Get(key)
	val = string(item.Value)
	return
}

func SetMemcacheVal(key string, val string) (err error) {
	err = Memcache.Set(&memcache.Item{Key: key, Value: []byte(val)})
	return
}

func DeleteMemcacheVal(key string) (err error) {
	err = Memcache.Delete(key)
	return
}
