package cache

import (
	"github.com/go-redis/redis/v7"
	"time"
)

type RedisCache struct {
	client *redis.Client
}

/**
GetBytes(key string) []byte
GetString(key string) string
GetMap(key string) map[string]string
GetList(key string) []string
GetValue(key string, val *interface{}) error

SetValue(key string, val interface{}) error
SetList(key string, val []string) error
SetString(key, val string) error
SetMap(key string, m map[string]string) error
SetBytes(key string, b []byte) error

Command(args... string) error
*/

func (rc *RedisCache) GetBytes(key string) []byte {
	s, e := rc.client.Get(key).Bytes()
	if e != nil {
		return nil
	}
	return s
}

func (rc *RedisCache) GetString(key string) string {
	s, e := rc.client.Get(key).Result()
	if e != nil {
		return ""
	}
	return s
}

func (rc *RedisCache) GetMap(key string) map[string]string {
	m, e := rc.client.HGetAll(key).Result()
	if e != nil {
		return nil
	}
	return m
}

func (rc *RedisCache) GetList(key string) []string {
	l, e := rc.client.LRange(key, 0, -1).Result()
	if e != nil {
		return nil
	}
	return l
}

func (rc *RedisCache) GetValue(key string) (string, error) {
	return rc.client.Get(key).Result()
}

func (rc *RedisCache) SetValue(key string, val interface{}, ex time.Duration) error {
	return rc.client.Set(key, val, ex).Err()
}

func (rc *RedisCache) SetMap(key string, m map[string]string, ex time.Duration) error {
	for k, v := range m {
		e := rc.client.HSet(key, k, v).Err()
		if e != nil {
			return e
		}
	}
	rc.Expire(key, ex)
	return nil
}

func (rc *RedisCache) SetBytes(key string, val []byte, ex time.Duration) error {
	return rc.client.Set(key, val, ex).Err()
}

func (rc *RedisCache) Command(args ...string) error {
	return rc.client.Do(args).Err()
}

func (rc *RedisCache) SetList(key string, val ...interface{}) error {
	rc.client.LPush(key, val...)
	return nil
}

func (rc *RedisCache) SetString(key string, val string, ex time.Duration) error {
	rc.client.Set(key, val, ex)
	return nil
}

func (rc *RedisCache) Expire(key string, ex time.Duration) {
	if ex > 0 {
		rc.client.Expire(key, ex)
	}
}
