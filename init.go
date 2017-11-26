package main

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/train-cat/notifier/api"
	"github.com/train-cat/notifier/helper"
)

func init() {
	initLogging()
	initConfig()
	helper.InitHelper()
	api.Init()
}

func initLogging() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)
}

func initConfig() {
	cfgFile := flag.String("config", "config.json", "config file")
	flag.Parse()

	viper.SetConfigFile(*cfgFile)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	} else {
		log.Error(err.Error())
		os.Exit(helper.ExitCodeErrorInitConfig)
	}
}
