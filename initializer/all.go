package initializer

import (
	"github.com/nbb2025/distri-domain/app/static/config"
	"github.com/nbb2025/distri-domain/pkg/util/str"
)

// InitAll 不直接用init函数，初始化顺序不确定，采用函数初始化更稳妥
func InitAll() {
	//根据具体项目需求调整初始化顺序
	config.ConfigInit()

	machineID := config.Conf.App.MachineID

	str.SnowflakeInit(machineID)

	loggerInit()

	pgsqlInit()
}
