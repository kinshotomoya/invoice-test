package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	customHanlder "invoice-test/internal/handler"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
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

	// web server

	server := http.Server{
		Addr: ":8080",
	}

	handler := &customHanlder.Handler{}

	http.HandleFunc("/api/invoices", handler.InvoiceHandler)

	go func() {
		log.Println("server is running")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server is not running. %s", err)
		}
	}()

	baseCtx := context.Background()
	signalCtx, cancel := signal.NotifyContext(baseCtx, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	defer cancel()

	<-signalCtx.Done()

	serverShutDownCtx, cancel := context.WithTimeout(baseCtx, 10*time.Second)
	defer cancel()
	err := server.Shutdown(serverShutDownCtx)
	if err != nil {
		log.Printf("fatal server shutdown grafefully")
		return
	}

	log.Println("process exit")

}
