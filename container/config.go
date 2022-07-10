package container

import (
	"github.com/crackeer/go-api-svc/define"
	"github.com/crackeer/gopkg/config"
)

// AppConfig appconfig
var AppConfig *define.AppConfig

// InitConfig ...
//  @param configPath
//  @return error
func InitConfig(configPath string) error {
	AppConfig = &define.AppConfig{}
	return config.LoadYamlConfig(configPath, AppConfig)
}
