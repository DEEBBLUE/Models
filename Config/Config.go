package serviceConfig

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type GlobalConfig struct {
	DBService  		ServiceConfig 	`yaml:"dbService"`
	ExProxy  			ServiceConfig 	`yaml:"exProxy"`
	RCService 		ServiceConfig 	`yaml:"rcService"`
	CRService 		ServiceConfig 	`yaml:"crService"`
	SsoService 		ServiceConfig 	`yaml:"ssoService"`
	RateService		ServiceConfig 	`yaml:"rateService"`
	StockAdapter 	ServiceConfig		`yaml:"stockAdapter"`
	ApiGetAway 		ServiceConfig 	`yaml:"ApiGetAway"`
	ConfService 	ServiceConfig		`yaml:"confService"`
}

type ServiceConfig struct {
	Host string 	`yaml: "host"`
	Port string   `yaml: "port"`
}

func InitConfig(path string) GlobalConfig {
	var conf GlobalConfig
	data,err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}	

	if err := yaml.Unmarshal(data,&conf);err != nil {
		panic(err)
	}
	return conf
}
