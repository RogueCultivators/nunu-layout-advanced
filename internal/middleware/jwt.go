package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nunu/nunu-layout-advanced/api/v1"
	"github.com/go-nunu/nunu-layout-advanced/pkg/jwt"
	"github.com/go-nunu/nunu-layout-advanced/pkg/log"
	"go.uber.org/zap"
)

type AuthMiddleware struct {
	j      *jwt.JWT
	logger *log.Logger
}

func NewAuthMiddleware(j *jwt.JWT, logger *log.Logger) *AuthMiddleware {
	return &AuthMiddleware{
		j:      j,
		logger: logger,
	}
}

func (am *AuthMiddleware) StrictAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			am.logger.WithContext(ctx).Warn("No token", zap.Any("data", map[string]interface{}{
				"url":    ctx.Request.URL,
				"params": ctx.Params,
			}))
			v1.HandleError(ctx, v1.ErrUnauthorized)
			ctx.Abort()
			return
		}

		claims, err := am.j.ParseToken(tokenString)
		if err != nil {
			am.logger.WithContext(ctx).Error("token error", zap.Any("data", map[string]interface{}{
				"url":    ctx.Request.URL,
				"params": ctx.Params,
			}), zap.Error(err))
			v1.HandleError(ctx, v1.ErrUnauthorized)
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		recoveryLoggerFunc(ctx, am.logger)
		ctx.Next()
	}
}

func (am *AuthMiddleware) NoStrictAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			tokenString, _ = ctx.Cookie("accessToken")
		}
		if tokenString == "" {
			tokenString = ctx.Query("accessToken")
		}
		if tokenString == "" {
			ctx.Next()
			return
		}

		claims, err := am.j.ParseToken(tokenString)
		if err != nil {
			ctx.Next()
			return
		}

		ctx.Set("claims", claims)
		recoveryLoggerFunc(ctx, am.logger)
		ctx.Next()
	}
}

func recoveryLoggerFunc(ctx *gin.Context, logger *log.Logger) {
	if userInfo, ok := ctx.MustGet("claims").(*jwt.MyCustomClaims); ok {
		logger.WithValue(ctx, zap.String("UserId", userInfo.UserId))
	}
}
