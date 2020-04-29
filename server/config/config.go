package config

import (
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func init() {
	release := flag.String("env", "dev", "dev or prod")
	flag.Parse()

	fmt.Println("env is :", *release)

	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(*release)

	config.AddConfigPath("profile/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}

}

//GetConfig ...
func GetConfig() *viper.Viper {
	return config
}
