package config

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/log"
	"github.com/pelletier/go-toml"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger log.Logger
var Sconfig ServerConfig

func init() {
	multiwriter := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "wangsu_cdn_operate.log",
		MaxSize:    200,
		MaxBackups: 3,
		MaxAge:     7,
		Compress:   false,
	})
	Logger = log.NewLogfmtLogger(multiwriter)
	Logger = log.With(Logger, "ts", log.DefaultTimestamp)
	Logger = log.With(Logger, "caller", log.DefaultCaller)
	serverConfigFile, err := AbsPath("config.toml")
	if err != nil {
		level.Error(Logger).Log("get config file path error", err)
		return
	}
	b, err := ioutil.ReadFile(serverConfigFile)
	if err != nil {
		level.Error(Logger).Log("read config file error", err)
		return
	}
	err = toml.Unmarshal(b, &Sconfig)
	if err != nil {
		level.Error(Logger).Log("unmarshal config error", err)
		return
	}
}

// ServerConfig 配置文件结构体
type ServerConfig struct {
	Global Global `toml:"global"`
}
type Global struct {
	Ak        string `toml:"ak"`
	Sk        string `toml:"sk"`
	Endpoint  string `toml:"endpoint"`
	Region    string `toml:"region"`
	Bucket    string `toml:"bucket"`
	Urlaction string `toml:"urlaction"`
	Diraction string `toml:"riraction"`
}

// AbsPath 获取文件绝对路径
func AbsPath(filename string) (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return ex, err
	}
	exPath := filepath.Dir(ex)
	_, err = os.Stat(exPath + "/" + filename)
	if err != nil {
		return exPath, err
	}
	absPath := exPath + "/" + filename
	return absPath, nil
}
