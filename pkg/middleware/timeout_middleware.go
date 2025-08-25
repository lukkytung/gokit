package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 创建一个带超时的 context
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// 将带超时的 context 替换到 request 中
		c.Request = c.Request.WithContext(ctx)

		// 创建一个通道来捕获处理完成信号
		done := make(chan struct{})

		go func() {
			c.Next() // 执行后续 handler
			close(done)
		}()

		select {
		case <-ctx.Done(): // 超时
			c.JSON(http.StatusGatewayTimeout, gin.H{
				"error":  "Request timed out",
				"detail": ctx.Err().Error(),
			})
			c.Abort() // 终止后续处理
		case <-done: // 正常完成
		}
	}
}
