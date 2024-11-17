package config

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var path = "app/static/config/config.yml"
var Conf = new(ProfileInfo)

type ProfileInfo struct {
	*App            `mapstructure:"app"`
	*RedisConfig    `mapstructure:"redis"`
	*PostgresConfig `mapstructure:"postgres"`
	*MongoConfig    `mapstructure:"mongo"`
	*JwtConfig      `mapstructure:"jwt"`
	*DllConfig      `mapstructure:"dll"`
	*MailConfig     `mapstructure:"mail"`
}

// 系统配置
type App struct {
	Env         string `mapstructure:"env"`
	Cache       bool   `mapstructure:"cache"`
	ServiceName string `mapstructure:"service-name"`
	MachineID   int64  `mapstructure:"machine-id"`
	ServerPort  int    `mapstructure:"server-port"`
	ApiPrefix   string `mapstructure:"api-prefix"`
}

// Redis配置
type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type PostgresConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

// MongoDB配置
type MongoConfig struct {
	DBName string `mapstructure:"dbname"`
	URI    string `mapstructure:"uri"`
}

// jwt配置
type JwtConfig struct {
	AccessExpire       int64  `mapstructure:"access-expire"`
	RefreshExpire      int64  `mapstructure:"refresh-expire"`
	Issuer             string `mapstructure:"issuer"`
	AccessTokenSecret  string `mapstructure:"asecret"`
	RefreshTokenSecret string `mapstructure:"rsecret"`
}

// dll加载配置
type DllConfig struct {
	DllPath string `mapstructure:"dll-path"`
	DllName string `mapstructure:"dll-name"`
}

// dll加载配置
type MailConfig struct {
	SmtpServer string `mapstructure:"smtp-server"`
	SmtpPort   int    `mapstructure:"smtp-port"`
	Writer     string `mapstructure:"email-writer"`
	Account    string `mapstructure:"account"`
	Password   string `mapstructure:"password"`
	BCC        string `mapstructure:"email-bcc"`
	CC         string `mapstructure:"email-cc"`
	URLPrefix  string `mapstructure:"url-prefix"`
}

// ConfigInit 初始化配置
// 将配置文件的信息反序列化到结构体中
func ConfigInit() {
	// 获取当前可执行文件的路径
	_exePath, e := os.Executable()
	if e != nil {
		fmt.Println("Error:", e)
		return
	}
	exePath := filepath.Dir(_exePath)
	path = filepath.Join(exePath, path)

	configFile := path
	s := flag.String("f", configFile, "choose conf file.")
	flag.Parse()
	//viper.AddConfigPath(configPath)
	//viper.SetConfigName("conf")     // 读取配置文件
	viper.SetConfigFile(*s)     // 读取配置文件
	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() faild error:%v\n", err)
		return
	}
	// 把读取到的信息反序列化到Conf变量中
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed,err:%v\n", err)
	}
	viper.WatchConfig()                            // （热加载时读取配置）监控配置文件
	viper.OnConfigChange(func(in fsnotify.Event) { // 配置文件修改时触发回调
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Printf("viper.Unmarshal failed,err:%v\n", err)
		}
	})
}

// SaveConfig 将配置写入文件，并在失败时回滚
func SaveConfig(configFile string, config *ProfileInfo) error {
	// 备份当前配置文件
	backupFile := configFile + ".bak"
	if err := backupConfigFile(configFile, backupFile); err != nil {
		return fmt.Errorf("failed to backup conf file: %v", err)
	}

	// 将 ProfileInfo 转换为 map(识别mapstructure标签)
	configMap := make(map[string]interface{})
	err := mapstructure.Decode(config, &configMap)
	if err != nil {
		rollbackConfigFile(configFile, backupFile)
		return fmt.Errorf("failed to decode conf to map: %v", err)
	}

	// 生成 YAML 数据
	data, err := yaml.Marshal(configMap)
	if err != nil {
		rollbackConfigFile(configFile, backupFile)
		return fmt.Errorf("failed to marshal conf data: %v", err)
	}

	// 写入配置文件
	if err := os.WriteFile(configFile, data, 0644); err != nil {
		rollbackConfigFile(configFile, backupFile)
		return fmt.Errorf("failed to write conf file: %v", err)
	}

	return nil
}

// 备份配置文件
func backupConfigFile(configFile, backupFile string) error {
	// 读取原始配置文件内容
	data, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			// 如果配置文件不存在，不需要备份，直接返回 nil
			return nil
		}
		return err
	}

	// 写入备份文件
	if err := os.WriteFile(backupFile, data, 0644); err != nil {
		return err
	}

	return nil
}

// 回滚配置文件
func rollbackConfigFile(configFile, backupFile string) {
	// 恢复备份文件内容到原始配置文件
	data, err := os.ReadFile(backupFile)
	if err != nil {
		fmt.Printf("error: failed to read backup file %s: %v\n", backupFile, err)
		return
	}

	if err := os.WriteFile(configFile, data, 0644); err != nil {
		fmt.Printf("error: failed to restore conf file from backup %s: %v\n", backupFile, err)
	}

	// 删除备份文件
	if err := os.Remove(backupFile); err != nil {
		fmt.Printf("warning: failed to remove backup file %s: %v\n", backupFile, err)
	}
}
