package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

type MetaToken struct {
	CredentialId string `json:"credential_id"`
	UserId       string `json:"user_id"`
	UserType     string `json:"user_type"`
	issuedAt     int64
	expiredAt    int64
}

func (m *MetaToken) Valid() error {
	now := time.Now().Unix()

	if delta := time.Unix(now, 0).Sub(time.Unix(m.expiredAt, 0)); delta >= 0 {
		return fmt.Errorf("token is expired by %v", delta)
	}

	if delta := time.Unix(now, 0).Sub(time.Unix(m.issuedAt, 0)); delta < 0 {
		return fmt.Errorf("token used before issued")
	}

	return nil
}

func NewJWTObj(jwtSecretKey string) *jwtAuth {
	return &jwtAuth{jwtSecretKey}
}

type jwtAuth struct {
	jwtSecretKey string
}

func (ja *jwtAuth) Sign(Data map[string]interface{}, ExpiredAtSec time.Duration) (string, error) {
	if ja.jwtSecretKey == "" {
		return "", fmt.Errorf("jwt secret key is missing")
	}

	expiredAt := time.Now()
	issuedAt := expiredAt.Add(time.Second * ExpiredAtSec)

	// metadata for your jwt
	claims := jwt.MapClaims{}
	claims["iat"] = issuedAt.Unix()
	claims["exp"] = expiredAt.Unix()

	for i, v := range Data {
		claims[i] = v
	}

	to := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	accessToken, err := to.SignedString([]byte(ja.jwtSecretKey))

	if err != nil {
		logrus.Error(err.Error())
		return accessToken, err
	}

	return accessToken, nil
}

func (ja *jwtAuth) VerifyToken(accessToken string) (token *jwt.Token, err error) {
	if ja.jwtSecretKey == "" {
		return nil, fmt.Errorf("jwt secret key is missing")
	}

	// verify issued date and expired date
	verifyCallbackFn := func(t *jwt.Token) (interface{}, error) {
		if err := t.Claims.Valid(); err != nil {
			return nil, err
		}
		return []byte(ja.jwtSecretKey), nil
	}

	// parse token
	if token, err = jwt.Parse(accessToken, verifyCallbackFn); err != nil {
		return
	}

	return
}
