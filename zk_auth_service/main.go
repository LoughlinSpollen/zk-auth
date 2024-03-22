package main

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"

	"zk_auth_service/debug"
	"zk_auth_service/pkg/infra/network/rpc"
	"zk_auth_service/pkg/usecase"

	cache "zk_auth_service/pkg/infra/cache"
	database "zk_auth_service/pkg/infra/db"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	_ = debug.Tags()
}

func main() {
	log.Trace("starting auth service")

	dbService := database.NewAuthDB()
	if err := dbService.Connect(); err != nil {
		log.Fatalf("could not connect to database : %v", err)
	}

	cache := cache.NewSessionCache()
	authUsecase := usecase.NewAuthUsecase(dbService, cache)

	service := rpc.NewAuthService(authUsecase)
	service.Connect()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		for {
			s := <-quit
			if s == syscall.SIGTERM || s == syscall.SIGINT {
				service.Close()
				dbService.Close()
				log.Trace("stopped auth service")
				return
			}
		}
	}()
}
