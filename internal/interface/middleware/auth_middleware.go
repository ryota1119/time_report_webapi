package middleware

import (
	"context"
	"fmt"
	"github.com/ryota1119/time_resport_webapi/internal/infrastructure/auth"
	"github.com/ryota1119/time_resport_webapi/internal/infrastructure/logger"
	"github.com/ryota1119/time_resport_webapi/internal/interface/handler"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ryota1119/time_resport_webapi/internal/domain/entities"
	"github.com/ryota1119/time_resport_webapi/internal/helper/auth_context"
)

type AuthMiddleware interface {
	AuthMiddleware() gin.HandlerFunc
	RequireAdmin() gin.HandlerFunc
}

type authMiddleware struct {
	authService auth.Service
}

var _ AuthMiddleware = (*authMiddleware)(nil)

func NewAuthMiddleware(
	authService auth.Service,
) AuthMiddleware {
	return &authMiddleware{
		authService: authService,
	}
}

// AuthMiddleware は AccessToken を検証するミドルウェア
func (a *authMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		req := c.Request

		tokenStr := ""

		// Authorization ヘッダーを取得
		authHeader := req.Header.Get("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenStr = strings.TrimPrefix(authHeader, "Bearer ")
		}

		if tokenStr == "" {
			if cookie, err := c.Cookie("access_token"); err == nil {
				tokenStr = cookie
			}
		}
		logger.Info(tokenStr)
		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, handler.ErrUnauthorized.Error())
			c.Abort()
			return
		}

		user, org, err := a.authService.AuthenticateFromToken(ctx, tokenStr)
		if err != nil {
			c.JSON(http.StatusUnauthorized, handler.ErrUnauthorized.Error())
			c.Abort()
			return
		}

		ctx = context.WithValue(ctx, "organization_id", org.ID)
		ctx = context.WithValue(ctx, "user_id", user.ID)
		ctx = context.WithValue(ctx, "role", user.Role)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// RequireAdmin はユーザーが管理者かどうかをチェックする
func (a *authMiddleware) RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		role := auth_context.ContextUserRole(ctx)
		fmt.Println("role", role)
		if auth_context.ContextUserRole(ctx) != entities.AdminRole {
			c.JSON(http.StatusForbidden, handler.ErrForbidden.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
