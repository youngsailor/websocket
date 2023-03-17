package config

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/spf13/viper"
	"path"
)

type Conf struct {
	Ip             string `mapstructure:"ip"`
	Port           int    `mapstructure:"port"`
	HeartbeatTime  int    `mapstructure:"heartbeat_time"`
	WorkerPoolSize uint64 `mapstructure:"worker_pool_size"`
	MaxWorkTaskLen uint32 `mapstructure:"max_work_task_len"`
	MaxConn        int    `mapstructure:"max_conn"` //待定
	Timeout        int    `mapstructure:"timeout"`
}

var WsConf Conf

func init() {
	viper.SetConfigName("socket") //配置文件名字
	viper.SetConfigType("toml")   //配置文件类型
	rootPath := gfile.SelfDir()
	cPath := path.Join(rootPath, "manifest/config")

	viper.AddConfigPath(cPath) //配置文件搜索路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&WsConf)
	if err != nil {
		panic(err)
	}

}
