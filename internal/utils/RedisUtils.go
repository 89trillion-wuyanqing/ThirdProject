package utils

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

// 声明一个全局的rdb变量
var Rdb *redis.Client

// 初始化连接
func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("连接")
	_, err := Rdb.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

}

//往redis  set操作
func StringPush(key string, value string, outTime time.Duration) error {
	err := Rdb.Set(key, value, outTime).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return errors.New("redis存储失败")
	}
	return nil

}

//redis  get操作
func StringPull(key string) (string, error) {
	val, err := Rdb.Get(key).Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
		return "", redis.Nil
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return "", errors.New("redis获取失败")
	} else {
		return val, nil
	}
}
