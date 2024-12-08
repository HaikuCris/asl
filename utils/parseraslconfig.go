package utils

import (
	"AslGo/config"

	"gopkg.in/yaml.v2"
)

type Configuration struct {
	App    AppConfiguration    `yaml:"app"`
	Linux  LinuxConfiguration  `yaml:"linux"`
	Deploy DeployConfiguration `yaml:"deploy"`
}

// app配置
type AppConfiguration struct {
	ASL_INIT_FOLDER string `yaml:"ASL_INIT_FOLDER"`
	ASL_VERSION     string `yaml:"ASL_VERSION"`
	ASL_MODE        string `yaml:"ASL_MODE"`
	ASL_PARTITIONS  string `yaml:"ASL_PARTITIONS"`
	ASL_LOG_RETURE  bool   `yaml:"ASL_LOG_RETURE"`
}

// linux配置
type LinuxConfiguration struct {
	PATH            string `yaml:"PATH"`
	LD_LIBRARY_PATH string `yaml:"LD_LIBRARY_PATH"`
	PROOT_TMP_DIR   string `yaml:"PROOT_TMP_DIR"`
	PROOT_LOADER    string `yaml:"PROOT_LOADER"`
}

// deploy配置
type DeployConfiguration struct {
	DEPLOYMENT_SEQUENCE_BEFORE string `yaml:"DEPLOYMENT_SEQUENCE_BEFORE"`
	DEPLOYMENT_SEQUENCE_LATER  string `yaml:"DEPLOYMENT_SEQUENCE_LATER"`
	CREATE_USER                bool   `yaml:"CREATE_USER"`
	USER_NAME                  string `yaml:"USER_NAME"`
	USER_PASSWORD              string `yaml:"USER_PASSWORD"`
	PRIVILEGED_USERS           string `yaml:"PRIVILEGED_USERS"`
}

// 返回asl.yaml解析后的参数
func GetAslConfig() *Configuration {
	var aslConfig Configuration
	err := yaml.Unmarshal([]byte(config.AslYAML), &aslConfig)
	if err != nil {
		panic(err)
	}
	return &aslConfig
}
