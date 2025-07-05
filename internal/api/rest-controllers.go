package api

import (
	pb "github.com/Lykeion-org/lexora-controller/internal/grpc/generated"
	"github.com/Lykeion-org/lexora-controller/internal/queues"
	"github.com/gin-gonic/gin"
)

func AddLanguageRoutes(rg *gin.RouterGroup, grpcConn pb.DictionaryServiceClient, nlpQueue *queues.MQProducer) {
	
	handler := &LanguageHandler{
		Client: grpcConn,
	}

	ref := rg.Group("/referent")
	ref.GET("/:uid", handler.GetReferent)
	ref.POST("/", handler.PostReferent)
	ref.PUT("/", handler.PutReferent)
	ref.GET("/search", handler.SearchReferent)

	sym := rg.Group("/symbol")
	sym.GET("/:uid", handler.GetSymbol)
	sym.POST("/", handler.PostSymbol)
	sym.PUT("/", handler.PutSymbol)

	wrd := rg.Group("/word")
	wrd.GET("/:uid", handler.GetWord)
	wrd.POST("/", handler.PostWord)
	wrd.PUT("/", handler.PutWord)


	rg.POST("/link-referent-symbol", handler.PostReferentSymbolLink)
	rg.POST("/link-symbol-word", handler.PostSymbolWordLink)
	rg.POST("/set-word-as-lemma", handler.PostWordLemmaLink)
}

func AddAdminRoutes(rg *gin.RouterGroup){

	handler := &AdminHandler{}
	rg.GET("/ping", handler.Ping)
}