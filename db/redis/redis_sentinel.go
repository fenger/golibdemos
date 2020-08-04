package redis

import (
	"fmt"
	"github.com/FZambia/sentinel"
	"github.com/gomodule/redigo/redis"
	"strings"
	"time"
)

var SentinelRedisConnPool *redis.Pool

func InitRedisSentinelConnPool() {
	redisAddr := "10.128.5.40:26379,10.128.7.19:26379,10.128.7.20:26379"
	redisAddrs := strings.Split(redisAddr, ",")
	masterName := "mymaster"

	sntl := &sentinel.Sentinel{
		Addrs:      redisAddrs,
		MasterName: masterName,
		Dial: func(addr string) (redis.Conn, error) {
			timeout := 500 * time.Millisecond
			c, err := redis.DialTimeout("tcp", addr, timeout, timeout, timeout)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}

	SentinelRedisConnPool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			masterAddr, err := sntl.MasterAddr()
			if err != nil {
				return nil, err
			}
			c, err := redis.Dial("tcp", masterAddr)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: CheckRedisRole,
	}
}

func CheckRedisRole(c redis.Conn, t time.Time) error {
	if !sentinel.TestRole(c, "master") {
		return fmt.Errorf("Role check failed")
	} else {
		return nil
	}
}

func SentinelSet() {
	InitRedisSentinelConnPool()
	rc := SentinelRedisConnPool.Get()
	defer rc.Close()

	_, err := rc.Do("SET", "hello", "world")
	if err != nil {
		fmt.Println("set failed", err)
		return
	}

}

func SentinelGet() {
	InitRedisSentinelConnPool()
	rc := SentinelRedisConnPool.Get()
	defer rc.Close()

	value, err := redis.String(rc.Do("GET", "hello"))
	if err != nil {
		fmt.Println("get failed", err)
		return
	}
	fmt.Printf("get the value: %v\n", value)

}
