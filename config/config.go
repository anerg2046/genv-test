package config

import "github.com/sakirsensoy/genv"

type appConfig struct {
	Host  string
	Port  int
	Debug bool
}

var App = &appConfig{
	Host:  genv.Key("APP_HOST").String(),
	Port:  genv.Key("APP_PORT").Default(8080).Int(),
	Debug: genv.Key("APP_DEBUG").Default(false).Bool(),
}
