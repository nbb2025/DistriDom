package initializer

import (
	"fmt"
	"github.com/nbb2025/distri-domain/app/static/config"
	"github.com/nbb2025/distri-domain/pkg/tool/cache"
	"github.com/nbb2025/distri-domain/pkg/tool/objstore"
	"github.com/nbb2025/distri-domain/pkg/tool/orm"
	myLogger "github.com/nbb2025/distri-domain/pkg/util/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func initGormProd(dsn string) (*gorm.DB, error) {
	defer myLogger.Sync() // 确保在程序退出前将日志刷入到输出中

	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,  //禁用外键
		QueryFields:                              false, //打印sql
		Logger: &myLogger.ZapGormLogger{
			LogLevel:      logger.Info,
			SlowThreshold: time.Second,
		},
	})
}

func initGormDev(dsn string) (*gorm.DB, error) {
	defer myLogger.Sync() // 确保在程序退出前将日志刷入到输出中

	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		QueryFields:                              true,
		Logger: &myLogger.ZapGormLogger{
			LogLevel:      logger.Info,
			SlowThreshold: time.Second,
		},
	})
}

// mysqlInit mysql连接初始化
func mysqlInit() {
	_config := config.Conf.MysqlConfig
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		_config.User, _config.Password, _config.Host, _config.Port, _config.DBName)
	var db *gorm.DB
	var err error

	if config.Conf.Env == "dev" {
		db, err = initGormDev(dsn)
	} else {
		db, err = initGormProd(dsn)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(_config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(_config.MaxOpenConns)

	if err != nil {
		fmt.Println(err)
		return
	}

	//// 自动迁移
	//db.AutoMigrate(
	//	// 表
	//	//model.File{},
	//	model.User{},
	//	model.Customer{},
	//	model.UserModule{},
	//)

	orm.SetDB(db)
}

// freeCacheInit FreeCache初始化
func cacheInit() {
	//cache.MyFreeCache = cache.NewFreeCache()
	//cache.MyFileCache = cache.NewFileCache("./cache")
	cache.MyRedis = cache.NewRedisCache()
}

func ossInit(product string) {
	var err error
	//objstore.AliOSSClient, err = objstore.NewAliyunOSSClient(
	//	conf.Conf.AliOSSConfig.Region, //服务器上传下载可用内网
	//	conf.Conf.AliOSSConfig.BucketName,
	//	conf.Conf.AliOSSConfig.AccessID,
	//	conf.Conf.AliOSSConfig.AccessSecret,
	//	conf.Conf.AliOSSConfig.RoleArn,
	//)
	objstore.S3c = objstore.NewS3Client(
		product,
		config.Conf.ObjStoreConfig.Region, //服务器上传下载可用内网
		config.Conf.ObjStoreConfig.BucketName,
		config.Conf.ObjStoreConfig.AccessID,
		config.Conf.ObjStoreConfig.AccessSecret,
	)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
