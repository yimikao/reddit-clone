package token

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTMaker struct {
	secretkey string
}

const minSecretKeyLen = 32

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeyLen {
		return nil, errors.New("secet key too short. min 32 character expected")
	}
	return &JWTMaker{secretkey: secretKey}, nil
}

func (m *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	claims, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(m.secretkey)
}

func (m *JWTMaker) VerifyToken(token string) (*Payload, error) {
	//I have three scenarios
	// (1) Invalid method
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(m.secretkey), nil
	}
	//(2) Expired token - found in token claims
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		vErr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(vErr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
	}
	// (3) claims is a valid payload
	p, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return p, nil

}
