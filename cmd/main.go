package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"os/signal"
	"syscall"
)

func initSettingFile() {
	viper.AutomaticEnv()
	viper.SetConfigName("setting")
	viper.SetConfigType("json")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal load config file %s", err.Error()))
	}
}

func main() {

	// init setting file
	initSettingFile()

	f := viper.GetString("MYSQL_DB_URL")
	fmt.Println(f)

	baseCtx := context.Background()
	signalCtx, cancel := signal.NotifyContext(baseCtx, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	defer cancel()

	<-signalCtx.Done()
	fmt.Println("done")

}
