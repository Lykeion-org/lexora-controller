package main

import (
	"os"
	"log"
	"log/slog"
)

type Config struct {
	ControllerServicePort string `yaml:"controller_service_port"`
	DictionaryServiceTarget string `yaml:"dictionary_service_target"`
	MessageQueueInstanceTarget string `yaml:"message_queue_instance_target"`
	TextAnalyzerJobQueueName string `yaml:"text_analyzer_job_queue_name"`
	TextAnalyzerFeedbackQueueName string `yaml:"text_analyzer_feedback_queue_name"`
	ServiceFeedbackQueueName string `yaml:"service_feedback_queue_name"`
	LogLevel string `yaml:"log_level"`
}

func initializeLogger(logLevel string){
	loggerConfig := &slog.HandlerOptions{
        Level: slog.LevelDebug,
	}
	switch logLevel {
	case "debug":
		loggerConfig.Level = slog.LevelDebug
	case "info":
		loggerConfig.Level = slog.LevelInfo
	case "warn":
		loggerConfig.Level = slog.LevelWarn
	case "error":
		loggerConfig.Level = slog.LevelError
	default:
		log.Println("Configuration log level is invalid, defaulting to debug")
		loggerConfig.Level = slog.LevelDebug
	}

	handler := slog.NewTextHandler(os.Stdout, loggerConfig)
	logger := slog.New(handler)
	
	slog.SetDefault(logger)
	log.SetOutput(slog.NewLogLogger(handler, slog.LevelInfo).Writer())
}