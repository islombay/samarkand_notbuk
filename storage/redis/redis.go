package redisdb

import (
	"fmt"
	"os"

	"github.com/islombay/noutbuk_seller/config"
	"github.com/islombay/noutbuk_seller/pkg/logs"
	"github.com/redis/go-redis/v9"
)

var (
	ErrKeyNotFound = fmt.Errorf("key_not_found")
)

type RedisStore struct {
	redis *redis.Client
	cfg   *config.RedisConfig
	log   logs.LoggerInterface

	otp *OTP
}

func NewRedisStore(cfg *config.RedisConfig, log logs.LoggerInterface) *RedisStore {
	client := redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			Password: os.Getenv("REDIS_PWD"),
			DB:       0,
		},
	)

	return &RedisStore{
		redis: client,
		cfg:   cfg,
		log:   log,

		otp: NewOTP(client, log),
	}
}

func (rs *RedisStore) OTP() *OTP {
	return rs.otp
}
