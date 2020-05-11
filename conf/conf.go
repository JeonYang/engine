package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

type engineConf struct {
	RpcPort string `yaml:"rpcPort"`
	DataDir string `yaml:"dataDir"`
	Logger  struct {
		// 保存文件
		Path         string `yaml:"path"`
		FileName     string `yaml:"fileName"`
		Level        string `yaml:"level"`
		MaxAge       int    `yaml:"maxAge"`       // 文件最大保存时间
		RotationTime int    `yaml:"rotationTime"` // 日志切割时间间隔
	}
}

var EngineConf = &engineConf{}

func InitConf(confFile string) error {
	dir, err := filepath.Abs(os.Args[0])
	if err != nil {
		return fmt.Errorf("get: %s abs fail,err: %v", os.Args[0], err)
	}
	AppAbsDir, _ = filepath.EvalSymlinks(dir)
	dir, err = os.Getwd()
	if err != nil {
		return fmt.Errorf("get pwd fail,err: %v", err)
	}
	if filepath.Base(dir) != "bin" {
		AppRootDir = dir
	} else {
		AppRootDir = filepath.Dir(dir)
	}

	if confFile == "" {
		confFile = filepath.Join(AppRootDir, "conf", defaultConfName)
	}

	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		return fmt.Errorf("read conf: %s fail,err: %v", confFile, err)
	}
	err = yaml.Unmarshal(yamlFile, EngineConf)
	if err != nil {
		return fmt.Errorf("yaml unmarshal fail,err: %v", err)
	}

	defaultConf()
	return nil
}

func defaultConf() {
	if EngineConf.RpcPort == "" {
		EngineConf.RpcPort = defaultRpcPort
	}
	if EngineConf.DataDir == "" {
		EngineConf.DataDir = filepath.Join(AppRootDir, "data")
	}
	if EngineConf.Logger.Path == "" {
		EngineConf.Logger.Path = filepath.Join(AppRootDir, "log")
	}
	if EngineConf.Logger.FileName == "" {
		EngineConf.Logger.FileName = defaultLogFileName
	}
	if EngineConf.Logger.Level == "" {
		EngineConf.Logger.Level = defaultLogLevel
	}
	if EngineConf.Logger.MaxAge == 0 {
		EngineConf.Logger.MaxAge = defaultLogMaxAge
	}
	if EngineConf.Logger.RotationTime == 0 {
		EngineConf.Logger.RotationTime = defaultLogRotationTime
	}
}

func (conf *engineConf) PluginPath(pluginName string) string {
	return filepath.Join(conf.DataDir, "plugin", pluginName)
}

func (conf *engineConf) EnginePath() string {
	return filepath.Join(conf.DataDir, "engine",)
}

func (conf *engineConf) PanicFile() string {
	return filepath.Join(conf.Logger.Path, defaultLogPanicFileName)
}
