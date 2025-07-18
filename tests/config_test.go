package tests 

import (
	"os"
	"testing"

	serviceConfig "github.com/DEEBBLUE/Models/Config"
)

func Test1(t *testing.T) {
	config_path := os.Getenv("CONFIG_PATH")
	conf := serviceConfig.InitConfig(config_path + "/config.yaml")

	if conf.DBService.Port != 4444{
		t.Errorf("Config not valid")
	}
}
