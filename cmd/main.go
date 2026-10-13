package main

import (
	"log/slog"

	"github.com/Lykeion-org/go-shared/pkg/config"
	"github.com/Lykeion-org/lexora-controller/internal/grpc"
	"github.com/Lykeion-org/lexora-controller/internal/handlers"
	"github.com/Lykeion-org/lexora-controller/internal/queues"
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

	slog.Info("Connecting to message queues", "jobs", cfg.TextAnalyzerJobQueueName, "results", cfg.TextAnalyzerFeedbackQueueName)
	mq, err := queues.NewConnection(cfg.MessageQueueInstanceTarget)
	if err != nil {
		panic(err)
	}
	defer mq.Close()

	jobs, err := queues.NewProducer(mq.Channel, cfg.TextAnalyzerJobQueueName)
	if err != nil {
		panic(err)
	}

	results, err := queues.NewConsumer(mq.Channel, cfg.TextAnalyzerFeedbackQueueName)
	if err != nil {
		panic(err)
	}
	feedbackHandler := handlers.NewTextAnalyzerFeedbackHandler(dictionary.Client)
	if err := results.Consume(feedbackHandler.HandleMessage); err != nil {
		panic(err)
	}

	lexora := grpc.NewLexoraGrpcServer(dictionary.Client, jobs)
	if err := lexora.StartServer(cfg.ControllerServicePort); err != nil {
		panic(err)
	}

	slog.Info("Started grpc server", "port", cfg.ControllerServicePort)


	select {}

}
