package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/nbb2025/distri-domain/app/static/config"
	"github.com/nbb2025/distri-domain/pkg/tool/jwt"
	"github.com/nbb2025/distri-domain/pkg/tool/req-resp/resp"
	"strings"
)

// noAuthRouters 不需要认证的路由
var noAuthRouters = []string{
	//config.Conf.App.ApiPrefix + "/base",
	"/base",
}

// JwtAuth jwt认证
// 解析请求头中的token
// 解析成功则通过，失败返回错误信息
func JwtAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 前端页面无需鉴权
		if !strings.HasPrefix(context.Request.URL.Path, config.Conf.App.ApiPrefix) {
			context.Next()
			return
		}

		for _, v := range noAuthRouters {
			if strings.HasPrefix(context.Request.URL.Path, v) {
				context.Next()
				return
			}
		}
		token := context.Request.Header.Get("AccessToken")
		tokenObj, err := jwt.ValidateToken(token, config.Conf.JwtConfig.AccessTokenSecret)
		if err != nil || !tokenObj.Valid {
			resp.Error(context, resp.CODE_TOKEN_EXPIRED)
			context.Abort()
			return
		}
		userAuthInfo, err := jwt.GetUserInfoFromJwt(tokenObj)
		if err != nil {
			resp.Error(context, resp.CODE_TOKEN_EXPIRED)
			context.Abort()
			return
		}
		// 将用户信息存储在Gin的上下文中，以便后续可以直接使用
		context.Set("userInfo", *userAuthInfo)
		context.Next()
	}
}
