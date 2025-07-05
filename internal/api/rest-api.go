package api

import (
	pb "github.com/Lykeion-org/lexora-controller/internal/grpc/generated"
	"github.com/Lykeion-org/lexora-controller/internal/queues"
	"github.com/gin-gonic/gin"
)

type api struct {
	router *gin.Engine
	port string
	dictionarySvc pb.DictionaryServiceClient
	nlpSvc *queues.MQProducer
}


func NewApi(dictionaryService pb.DictionaryServiceClient, nlpService *queues.MQProducer) *api {
	r := gin.Default()

	return &api {
		router: r,
		port: ":8080",
		dictionarySvc: dictionaryService,
		nlpSvc: nlpService,
		
	}
}

func (a *api) InitializeApi(port string) {
	a.port = port
	v1 := a.router.Group("/api/v1")
	AddLanguageRoutes(v1, a.dictionarySvc, a.nlpSvc)
	AddAdminRoutes(v1)
	

	a.router.Run(port)
}

