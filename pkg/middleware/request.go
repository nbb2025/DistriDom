package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/nbb2025/distri-domain/pkg/tool/req-resp/resp"
	"github.com/nbb2025/distri-domain/pkg/util/logger"
	"go.uber.org/zap"
	"io"
)

func JsonDataMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		data, err := context.GetRawData()
		if err != nil {
			logger.Error("Error reading raw data:", zap.Error(err))
			resp.Error(context, resp.CODE_NO_PERMISSIONS, resp.CODE_INVALID_PARAMETER.Msg())
			context.Abort()
			return
		}
		// 恢复请求体流数据
		context.Request.Body = io.NopCloser(bytes.NewBuffer(data))

		// 将解析后的JSON数据存储到Context中
		context.Set("rawRequestData", data)

		context.Next()

	}
}
