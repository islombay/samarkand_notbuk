package handlers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/islombay/noutbuk_seller/api/status"
	"github.com/islombay/noutbuk_seller/pkg/auth"
)

type AuthorizationReq struct {
	Auth string `header:"Authorization" binding:"required"`
}

func (v1 *Handler) isAuth(c *gin.Context) (*auth.Token, *status.Status) {
	var m AuthorizationReq
	if err := c.BindHeader(&m); err != nil {
		return nil, &status.StatusUnauthorized
	}

	token, err := auth.ParseToken(m.Auth)
	if err != nil {
		if errors.Is(err, auth.ErrTokenInvalid) || errors.Is(err, auth.ErrTokenExpired) {
			return nil, &status.StatusUnauthorized
		}
		return nil, &status.StatusInternal
	}

	return token, nil
}

func (v1 *Handler) MiddlewareIsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, errStatus := v1.isAuth(c)
		if errStatus != nil {
			v1.response(c, *errStatus)
			return
		}

		if token.Role != auth.RoleAdmin {
			v1.response(c, status.StatusForbidden)
			return
		}

		c.Set("token", token)
		c.Next()
	}
}

func GetContextValue[T any](c *gin.Context, key string) (T, bool) {
	value, exists := c.Get(key)
	if !exists {
		var zero T
		return zero, false
	}

	// Perform type assertion
	result, ok := value.(T)
	if !ok {
		var zero T
		return zero, false
	}
	return result, true
}
