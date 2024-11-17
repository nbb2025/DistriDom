package main

// 使用 _引入依赖项在main函数执行会直接调用init函数
import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nbb2025/distri-domain/app/api/router"
	"github.com/nbb2025/distri-domain/app/static/config"
	"github.com/nbb2025/distri-domain/app/static/embeded"
	"github.com/nbb2025/distri-domain/initializer"
	"github.com/nbb2025/distri-domain/pkg/middleware"
	"github.com/nbb2025/distri-domain/pkg/util/ginstatic"
	"github.com/nbb2025/distri-domain/pkg/util/logger"
)

func main() {
	initializer.InitAll()

	if config.Conf.App.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	// 启动服务(使用goroutine解决服务启动时程序阻塞问题)
	go RunServer()

	// 监听信号
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	select {
	case <-signals:
		// 释放资源
		logger.Sync()
		fmt.Println("[GIN-QuickStart] 程序关闭，释放资源")
		return
	}
}

// RunServer 启动服务
func RunServer() {
	// 初始化引擎
	r := gin.New()

	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))

	// 局域网项目跨域
	r.Use(cors.Default())

	// 注册JWT认证中间件
	r.Use(middleware.JwtAuth())

	// 注册防抖中间件
	r.Use(middleware.ThrottleMiddleware())

	// 注册路由
	apiGroup := r.Group(config.Conf.App.ApiPrefix)
	{
		for _, f := range router.Routers {
			f(apiGroup)
		}
	}
	//设置静态内容
	setFrontStaticEmbed(r)

	//fmt.Printf("[GIN-QuickStart] 接口文档地址：http://127.0.0.1:%v/swagger/index.html\n", conf.Conf.ServerPort)
	fmt.Printf("[GIN-QuickStart] 前端页面：http://0.0.0.0:%v/\n", config.Conf.ServerPort)
	fmt.Printf("启动时间:%v\n", time.Now().Format(time.DateTime))
	r.Run(fmt.Sprintf("0.0.0.0:%v", config.Conf.ServerPort))

}

func setFrontStaticFileSystem(r *gin.Engine) {
	// 设置静态文件夹
	staticDir := "app/static/embeded/web"

	// 检查文件夹是否存在
	if _, err := os.Stat(staticDir); os.IsNotExist(err) {
		log.Fatalf("Static directory does not exist: %s", staticDir)
	}

	// 使用 Gin 提供静态文件服务
	r.Use(ginstatic.Serve("/", ginstatic.LocalFile(staticDir, true)))
	r.NoRoute(func(context *gin.Context) {
		accept := context.GetHeader("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := os.ReadFile(staticDir + "/index.html")
			if (err) != nil {
				context.Writer.WriteHeader(404)
				context.Writer.WriteString("Not Found")
				return
			}
			context.Writer.WriteHeader(200)
			context.Writer.Header().Add("Accept", "text/html")
			context.Writer.Write(content)
			context.Writer.Flush()
		}
	})
}

func setFrontStaticEmbed(r *gin.Engine) {
	distFile := embeded.FsWeb
	// 使用 Gin 提供静态文件服务
	r.Use(ginstatic.ServeEmbed("/", distFile))
	r.NoRoute(func(context *gin.Context) {
		accept := context.GetHeader("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := distFile.ReadFile("index.html")
			if err != nil {
				context.Writer.WriteHeader(404)
				context.Writer.WriteString("Not Found")
				return
			}
			context.Writer.WriteHeader(200)
			context.Writer.Header().Add("Accept", "text/html")
			context.Writer.Write(content)
			context.Writer.Flush()
		}
	})
}
