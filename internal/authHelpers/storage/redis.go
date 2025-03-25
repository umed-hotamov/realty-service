package storage

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/umed-hotamov/realty-service/internal/domain"
)

type AuthSession struct {
	RefreshToken string
	RefreshExp   int64
	Fingerprint  string
	Payload      domain.AuthPayload
}

type SessionStorage struct {
	redisClient *redis.Client
}

func NewSessionStorage(redisClient *redis.Client) *SessionStorage {
	return &SessionStorage{redisClient: redisClient}
}

func (s *SessionStorage) Get(refreshToken string) (AuthSession, error) {
	sessionJson, err := s.redisClient.Get(refreshToken).Bytes()
	if err != nil {
		return AuthSession{}, err
	}

	var session AuthSession
	err = json.Unmarshal(sessionJson, &session)
	if err != nil {
		return AuthSession{}, err
	}

	return session, nil
}

func (s *SessionStorage) Put(refreshToken string, session AuthSession,
	expireTime time.Duration) error {
	sessionJson, err := json.Marshal(session)
	if err != nil {
		return err
	}

	return s.redisClient.Set(refreshToken, sessionJson, expireTime).Err()
}

func (s *SessionStorage) Delete(refreshToken string) error {
	return s.redisClient.Del(refreshToken).Err()
}
