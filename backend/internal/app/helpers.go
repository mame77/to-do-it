package app

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// NewUUID ... 新しいUUIDを生成
func NewUUID() string {
	return uuid.New().String()
}

// GetUserIDFromContext ... 本来はJWTトークンから取得する
func GetUserIDFromContext(c echo.Context) string {
	// TODO: 担当DがJWT認証を実装したら、
	// claims := c.Get("user").(*jwt.Token).Claims.(*MyCustomClaims)
	// return claims.UserID
	return "user_123" // (実行確認用の固定値)
}
