package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/islombay/noutbuk_seller/api"
	"github.com/islombay/noutbuk_seller/config"
	"github.com/islombay/noutbuk_seller/pkg/logs"
	"github.com/islombay/noutbuk_seller/pkg/start"
	"github.com/islombay/noutbuk_seller/service"
	"github.com/islombay/noutbuk_seller/storage/postgresql"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found" + err.Error())
	}

	cfg := config.Load()

	log := logs.NewLogger("app", logs.LevelDebug)
	defer func() {
		if err := logs.Cleanup(log); err != nil {
			return
		}
	}()

	storage := postgresql.NewPostgresStore(cfg.DB, log)
	defer storage.Close()

	services := service.New(storage, log)

	start.Init(cfg.DB, log)

	r := gin.Default()

	api.NewApi(r, services, cfg, log)

	go func() {
		if err := r.Run(cfg.Server.Host + ":" + cfg.Server.Port); err != nil {
			log.Panic("Error listening server", logs.Error(err))
			os.Exit(1)
		}
	}()

	tickerPing := time.NewTicker(2 * time.Minute)

	log.Debug("setting picker for ping")
	go func() {
		for range tickerPing.C {
			sendRequest(
				fmt.Sprintf("http://%s/%s",
					cfg.Server.Public,
					"ping",
				))
		}
	}()

	log.Info("Server running")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	tickerPing.Stop()
}

func sendRequest(url string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		defer res.Body.Close()
	}
}
