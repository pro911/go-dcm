package dal

import (
	"apipost-dcm/internal/pkg/conf"
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
)

var (
	rdb *redis.Client
)

// MustInitRedis 必要加载redis
func MustInitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Conf.Redis.Host, conf.Conf.Redis.Port),
		Username: conf.Conf.Redis.UserName,
		Password: conf.Conf.Redis.Password,
		DB:       conf.Conf.Redis.DataBase,
	})

	defer rdb.Close()

	rdb.Ping(context.Background())

	fmt.Println("redis initialized")
}

func GetRDB() *redis.Client {
	return rdb
}
