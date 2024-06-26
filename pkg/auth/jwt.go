package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Token struct {
	UserID string `json:"user_id"`
	Role   string `json:"role,omitempty"`
	Type   string `json:"type,omitempty"`
	jwt.RegisteredClaims
}

var (
	ErrUnexpectedSigningMethod = fmt.Errorf("unexpected signing method")
	ErrTokenInvalid            = fmt.Errorf("invalid token")
	ErrTokenExpired            = fmt.Errorf("token expired")
)

var (
	RoleClient = "client"
	RoleAdmin  = "admin"

	TokenAccess  = "access"
	TokenRefresh = "refresh"
)

const (
	TokenRefreshExpireLife = 200
	TokenAccessExpireLife  = 2
)

func GenerateToken(t Token) (string, error) {
	duration := time.Hour * TokenAccessExpireLife

	claims := Token{
		UserID: t.UserID,
		// Type:   TokenAccess,
		Role: t.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateTokenRefresh(t Token) (string, error) {
	duration := time.Hour * TokenRefreshExpireLife
	claims := Token{
		UserID: t.UserID,
		Type:   TokenRefresh,
		Role:   t.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokens string) (*Token, error) {
	tkn, err := jwt.ParseWithClaims(tokens, &Token{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrUnexpectedSigningMethod
		}
		return []byte(os.Getenv("secret")), nil
	})
	if err != nil {
		if errors.Is(err, ErrUnexpectedSigningMethod) {
			return nil, ErrTokenInvalid
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		} else if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, ErrTokenInvalid
		} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return nil, ErrTokenInvalid
		} else if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, ErrTokenInvalid
		}
		return nil, err
	}

	if claims, ok := tkn.Claims.(*Token); ok && tkn.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}
