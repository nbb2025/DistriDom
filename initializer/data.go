package initializer

import (
	"fmt"
	"time"

	"github.com/nbb2025/distri-domain/app/core/model"
	"github.com/nbb2025/distri-domain/app/static/config"
	"github.com/nbb2025/distri-domain/pkg/tool/orm"
	myLogger "github.com/nbb2025/distri-domain/pkg/util/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func pgsqlInit() {
	_config := config.Conf.PostgresConfig
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai",
		_config.Host, _config.Username, _config.Password, _config.Database, _config.Port)
	var db *gorm.DB
	var err error

	if config.Conf.Env == "dev" {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			QueryFields:                              true,
			Logger: &myLogger.ZapGormLogger{
				LogLevel:      logger.Info,
				SlowThreshold: time.Second,
			},
		})
	} else {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			QueryFields:                              true,
		})
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	// Auto migrate models
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		fmt.Println("Failed to auto migrate:", err)
		return
	}
	orm.SetDB(db)
}
