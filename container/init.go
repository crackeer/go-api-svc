package container

import (
	"fmt"

	"github.com/crackeer/go-api-svc/define"
)

// Init ...
//  @param configPath
//  @return error
func Init(configPath string) error {

	if err := InitConfig(configPath); err != nil {
		panic(fmt.Sprintf("init app config error: %s", err.Error()))
	}

	if err := InitDB(AppConfig.DatabaseConfDir, AppConfig.Env); err != nil {
		panic(fmt.Sprintf("init database config error: %s", err.Error()))
	}
	return nil
}

// GetAppConfig ...
//  @return *define.AppConfig
func GetAppConfig() *define.AppConfig {
	return AppConfig
}
