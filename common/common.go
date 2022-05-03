package common

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid := uuid.New().String()
		c.Set("xRequestId", uuid)
		c.Header("X-Request-Id", uuid) // before call c.json
		c.Next()
	}
}
