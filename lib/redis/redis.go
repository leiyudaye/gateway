/*
 * @Descripttion:
 * @Author: lly
 * @Date: 2021-06-11 22:28:42
 * @LastEditors: lly
 * @LastEditTime: 2021-06-12 03:00:52
 */
package redis

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func Get(conf string, key string, v interface{}) error {
	conn, err := redis.Dial("tcp", conf)
	if err != nil {
		return fmt.Errorf("redis connect failed, err=%v", err)
	}
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return fmt.Errorf("redis set failed, err=%v", err)
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return fmt.Errorf("json marshal failed, err=%v", err)
	}
	return nil
}

func Set(conf string, key string, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("json marshal failed, err=%v", err)
	}

	conn, err := redis.Dial("tcp", conf)
	if err != nil {
		return fmt.Errorf("redis connect failed, err=%v", err)
	}
	_, err = conn.Do("SET", key, string(data))
	if err != nil {
		return fmt.Errorf("redis set failed, err=%v", err)
	}
	return nil
}

func ZAdd(conf string, key string, field string, value int64) (int64, error) {
	conn, err := redis.Dial("tcp", conf)
	if err != nil {
		return 0, fmt.Errorf("redis connect failed, err=%v", err)
	}
	score, err := redis.Int64(conn.Do("ZADD", key, value, field))
	if err != nil {
		return 0, fmt.Errorf("zadd failed, err=%v", err)
	}
	return score, nil
}

func ZIncr(conf string, key string, field string, value int64) (int64, error) {
	conn, err := redis.Dial("tcp", conf)
	if err != nil {
		return 0, fmt.Errorf("redis connect failed, err=%v", err)
	}
	score, err := redis.Int64(conn.Do("ZINCR", key, value, field))
	if err != nil {
		return 0, fmt.Errorf("zadd failed, err=%v", err)
	}
	return score, nil
}

func ZRem(conf string, key string, field string) error {
	conn, err := redis.Dial("tcp", conf)
	if err != nil {
		return fmt.Errorf("redis connect failed, err=%v", err)
	}
	_, err = conn.Do("ZREM", key, field)
	if err != nil {
		return fmt.Errorf("zadd failed, err=%v", err)
	}
	return nil
}
