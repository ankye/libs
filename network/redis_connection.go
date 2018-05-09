package network

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/gonethopper/libs/logs"
)

//RedisConnection redis connection wrapper
type RedisConnection struct {
	Address  string
	Password string
	Redis    redis.Conn
}

//NewRedisConnection 创建一个redis connection
func NewRedisConnection(address, password string) (*RedisConnection, error) {
	conn := new(RedisConnection)
	conn.Address = address
	conn.Password = password
	var c redis.Conn
	var err error
	//gets the Redis connection.
	if c, err = redis.Dial("tcp", address); err != nil {
		return nil, err
	}

	if len(password) != 0 {
		if _, err = c.Do("AUTH", password); err != nil {
			return nil, err
		}
	}
	conn.Redis = c

	return conn, nil
}

//Ping 检测连接
func (s *RedisConnection) Ping() error {
	pongStr := ""
	var err error
	if pongStr, err = redis.String(s.Redis.Do("PING")); err != nil {
		return err
	}

	if pongStr != "PONG" {
		err = fmt.Errorf("Redis PING != PONG(%v)", pongStr)
		return err
	}
	return nil
}

//Get get value from redis, key is string
func (s *RedisConnection) Get(key string) (interface{}, error) {
	v, err := s.Redis.Do("GET", key)
	if err == nil {
		return v, nil
	}
	return nil, err

}

//Set set value to redis,key is string, if timeout is setted, than key will have Expire,单位秒
func (s *RedisConnection) Set(key string, val interface{}, timeout time.Duration) error {
	_, err := s.Redis.Do("SET", key, val)
	if err != nil {
		return err
	}
	if timeout != 0 {
		return s.SetExpire(key, timeout)
	}
	return nil
}

//SetExpire set expire time for key
func (s *RedisConnection) SetExpire(key string, timeout time.Duration) error {
	_, err := s.Redis.Do("Expire", key, int64(timeout/time.Second))
	if err != nil {
		logs.Error("redis set expire key err [%v]", err)
		return err
	}
	return err
}

//Incr 指定自增key
func (s *RedisConnection) Incr(key string) error {
	_, err := redis.Bool(s.Redis.Do("INCRBY", key, 1))
	return err
}

//Decr 指定自减key
func (s *RedisConnection) Decr(key string) error {
	_, err := redis.Bool(s.Redis.Do("INCRBY", key, -1))
	return err
}

//Del 根据key 删除value
func (s *RedisConnection) Del(key string) error {
	if _, err := s.Redis.Do("DEL", key); err != nil {
		return err
	}
	return nil
}

//IsExists 判断key是否存在
func (s *RedisConnection) IsExists(key string) bool {
	v, err := redis.Bool(s.Redis.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return v
}
