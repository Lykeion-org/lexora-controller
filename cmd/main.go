package main

import (
	"log/slog"
	"github.com/Lykeion-org/go-shared/pkg/config"
	"github.com/Lykeion-org/lexora-controller/internal/grpc"
)

func main(){

	cfg, err := config.LoadConfigFile[Config]("../config.yaml")
	if err != nil {
		panic(err)
	}

	initializeLogger(cfg.LogLevel)

	slog.Info("Starting lexora application...")
	slog.Info("Opening grpc connections")

	dictionary := grpc.NewDictionaryClient(cfg.DictionaryServiceTarget); defer dictionary.StopClient()

	lexora := grpc.NewLexoraGrpcServer(dictionary.Client)
	lexora.StartServer(cfg.ControllerServicePort)

	slog.Info("Started grpc server", "port", cfg.ControllerServicePort)


	select {}

}