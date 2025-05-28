package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"onestep/internal/infrastructure/config"
	"onestep/internal/infrastructure/database"
	"onestep/internal/interface/http/routes"
)

//TIP Run this program with `go run main.go`

func main() {
	//load the app config by env
	loadAppConfig()
	//
	database.InitDB()
	//init gin
	r := gin.Default()
	routes.RegisterRoutes(r)
	err := r.Run(":8080")
	if err != nil {
		return
	}
}

// loadConfig reads application_dev.yml and sets the config in AppConfig
func loadAppConfig() {
	env := flag.String("app_env", "", "the app env")
	flag.Parse()
	// Set up viper to read the yaml file
	viper.SetConfigName(fmt.Sprintf("application_%s.yml", *env))
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	// Read the yaml file
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(config.AppConfig)
	if err != nil {
		panic(err)
	}
}
