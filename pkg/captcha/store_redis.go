package captcha

import (
	"errors"
	"gohub/pkg/app"
	"gohub/pkg/config"
	"gohub/pkg/redis"
	"time"
)

// RedisStore 实现 base64Captcha.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	keyPrefix   string
}

// Set 实现 base64Captcha.Store interface的Set方法
func (s *RedisStore) Set(key string, value string) error {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))

	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("captcha.debug_expire_time"))
	}

	if ok := s.RedisClient.Set(s.keyPrefix+key, value, ExpireTime); !ok {
		return errors.New("无法存储验证码答案")
	}

	return nil
}

// Get
func (s *RedisStore) Get(key string, clear bool) string {
	key = s.keyPrefix + key
	val := s.RedisClient.Get(key)

	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

// Verify
func (s *RedisStore) Verify(key string, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
