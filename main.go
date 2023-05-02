package main

import (
	"log"
	"os"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

var logger *zap.Logger

type RouteCfg struct {
	Route string `yaml:"route"`
	Path  string `yaml:"path"`
}

type Config struct {
	Port         string     `yaml:"port"`
	AccessLog    bool       `ỳaml:"accessLog"`
	StaticRoutes []RouteCfg `yaml:"staticRoutes"`
}

func main() {

	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalln("main:: could not setup zap logger")
	}

	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		logger.Fatal("could not open the config.yaml file")
	}

	config := Config{}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		logger.Fatal("could not convert the configuration", zap.Error(err))
	}

	app := iris.New()

	if config.AccessLog {
		ac := accesslog.File("./access.log")
		ac.AddOutput(os.Stdout)
		app.UseRouter(ac.Handler)
	}

	for _, value := range config.StaticRoutes {
		logger.Info("add new route -->" + value.Route + " || " + value.Path)
		app.HandleDir(value.Route, value.Path)
	}

	logger.Info("going to start the server on port " + ":" + config.Port)
	app.Run(iris.Addr(":" + config.Port))
}
