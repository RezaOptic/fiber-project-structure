package main

import (
	"flag"
	"github.com/RezaOptic/fiber-project-structure/config"
	"github.com/RezaOptic/fiber-project-structure/logic"
	"github.com/RezaOptic/fiber-project-structure/repository"
	grpcEngine "github.com/RezaOptic/fiber-project-structure/router/grpc"
	httpEngine "github.com/RezaOptic/fiber-project-structure/router/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)


func main() {
	configFlag := flag.String("config", "dev", "config flag can be dev for develop or prod for production")
	prodConfigPath := flag.String("config-path", "", "config-path production config file path")

	// init service configs
	config.Init(configFlag, prodConfigPath)

	// init repositories
	repository.Init()

	UserRepo := repository.NewUserRepo(repository.DBS.SqlConnection)
	UserLogic := logic.NewUserLogic(&UserRepo)
	UserController := httpEngine.NewUserController(UserLogic)
	// run http and grpc servers
	wg := sync.WaitGroup{}
	wg.Add(2)
	go httpEngine.Run(config.Configs.Service.HttpPort, UserController)
	go grpcEngine.Run(config.Configs.Service.GrpcPort)

	// handle os signals
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		select {
		case <-sigc:
			// wg.Done()
			// TODO...
			os.Exit(1)
		}
	}()

	wg.Wait()
}
