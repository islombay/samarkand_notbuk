package redisdb

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/islombay/noutbuk_seller/pkg/logs"
	"github.com/redis/go-redis/v9"
)

type OTP struct {
	redis *redis.Client
	log   logs.LoggerInterface
}

func NewOTP(r *redis.Client, log logs.LoggerInterface) *OTP {
	return &OTP{
		redis: r,
		log:   log,
	}
}

type SetCodeRequest struct {
	UserID      string
	PhoneNumber string
	Code        string
}

type Code struct {
	RequestID string `json:"request_id"`
	UserID    string `json:"user_id"`
	NeedPhone bool   `json:"need_phone"`
	Code      string `json:"code"`
}

const CodeExpireTime = time.Hour / 4

func (otp *OTP) validateKey(s string) string {
	return fmt.Sprintf("code-samarkand-notbuk:%s", s)
}

func (otp *OTP) SetCode(ctx context.Context, req SetCodeRequest) (*Code, error) {
	code := Code{
		RequestID: uuid.NewString(),
		UserID:    req.UserID,
		NeedPhone: true,
		Code:      req.Code,
	}
	jsonData, err := json.Marshal(&code)
	if err != nil {
		otp.log.Error("could not marshal redis", logs.Error(err))
		return nil, err
	}

	err = otp.redis.Set(
		ctx,
		otp.validateKey(req.PhoneNumber),
		jsonData,
		CodeExpireTime,
	).Err()

	if err != nil {
		otp.log.Error("could not set code to redis", logs.Error(err))
		return nil, err
	}

	return &code, nil
}

func (otp *OTP) GetCode(ctx context.Context, key string) (*Code, error) {
	val, err := otp.redis.Get(ctx, otp.validateKey(key)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrKeyNotFound
		}
		otp.log.Error("could not get val from redis", logs.Error(err))
		return nil, err
	}

	var code Code
	err = json.Unmarshal([]byte(val), &code)
	if err != nil {
		otp.log.Error("could not unmarshal redis to obj", logs.Error(err))
		return nil, err
	}

	return &code, nil
}
