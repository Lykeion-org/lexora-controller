package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	pb "github.com/Lykeion-org/lexora-controller/internal/grpc/generated"
	"github.com/Lykeion-org/lexora-controller/internal/queues"
	"github.com/gin-gonic/gin"
)

type LanguageHandler struct {
	Client pb.DictionaryServiceClient
}

func(lh *LanguageHandler) GetReferent(c *gin.Context){
	resp, err := lh.Client.GetReferent(context.Background(), &pb.GetReferentRequest{Uid: c.Param("uid")})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}	
	
	ref := ProtoToReferent(resp)
		c.JSON(http.StatusOK, ref)
}

func(lh *LanguageHandler) PostReferent(c *gin.Context){
	var req CreateReferentRequest
		
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})			
		return
	}

	grpcReq := &pb.CreateReferentRequest{			
		EnReference: req.EnReference,
		ImageSource: &req.ImageSource,
	}

	res, err := lh.Client.CreateReferent(context.Background(), grpcReq); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})			
		return }

	ref := ProtoToReferent(res)
	c.JSON(http.StatusOK, ref)
}

func(lh *LanguageHandler) PutReferent(c *gin.Context){
	var req Referent
		
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &pb.UpdateReferentRequest{
		Referent: &pb.Referent{
		Uid: req.Uid,
		EnReference: req.EnReference,
		ImageSource: req.ImageSource,
		},
	}

	res, err := lh.Client.UpdateReferent(context.Background(), grpcReq); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return }

	ref := ProtoToReferent(res)
			
	c.JSON(http.StatusOK, ref)
}

func(lh *LanguageHandler) SearchReferent(c *gin.Context){
	query := c.Query("query")

	resp, err := lh.Client.FindReferents(context.Background(), &pb.FindReferentsRequest{Query: query, SearchMode: pb.SearchMode_SEARCH_MODE_WORD})
	if err != nil { 
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	referents := ProtoToReferentList(resp.GetReferents())
	c.JSON(http.StatusOK, referents)

}

func(lh *LanguageHandler) GetSymbol(c *gin.Context){
	res, err := lh.Client.GetSymbol(context.Background(), &pb.GetSymbolRequest{Uid: c.Param("uid")})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}	

	sym := ProtoToSymbol(res)
		c.JSON(http.StatusOK, sym)
}

func(lh *LanguageHandler) PostSymbol(c *gin.Context){
	var req CreateSymbolRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})			
		return
	}

	grpcReq := &pb.CreateSymbolRequest{
		ReferentUid: req.ReferentUid,			
		Language:    pb.Language(req.Language),
		SymbolType:  pb.SymbolType(req.SymbolType),
	}

	res, err := lh.Client.CreateSymbol(context.Background(), grpcReq); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})			
		return 
	}

	sym := ProtoToSymbol(res)
	c.JSON(http.StatusOK, sym)

}

func(lh *LanguageHandler) PutSymbol(c *gin.Context){
	var req Symbol
		
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &pb.UpdateSymbolRequest{
		Symbol: &pb.Symbol{
		Uid: req.Uid,
		Language: pb.Language(req.Language),
		SymbolType: pb.SymbolType(req.SymbolType),
		},
	}

	fmt.Println(grpcReq);

	res, err := lh.Client.UpdateSymbol(context.Background(), grpcReq); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return }

	ref := ProtoToSymbol(res)
			
	c.JSON(http.StatusOK, ref)
}


func(lh *LanguageHandler) GetWord(c *gin.Context){
	res, err := lh.Client.GetWord(context.Background(), &pb.GetWordRequest{Uid: c.Param("uid")})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	
		wrd := ProtoToWord(res)
		c.JSON(http.StatusOK, wrd)
}

func(lh *LanguageHandler) PostWord(c *gin.Context){
	var req CreateWordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})			
		return
	}

	grpcReq := &pb.CreateWordRequest{
		SymbolUid: 		req.SymbolUid,
		Word: req.Word,
		WordType: pb.WordType(req.WordType),
		SoundSource: &req.SoundSource,
		Ipa: &req.Ipa,
	}

	res, err := lh.Client.CreateWord(context.Background(), grpcReq); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})			
		return 
	}

	wrd := ProtoToWord(res)
	c.JSON(http.StatusOK, wrd)
}

func(lh *LanguageHandler) PutWord(c *gin.Context){
	var req Word
		
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &pb.UpdateWordRequest{
		Word: &pb.Word{
		Uid: req.Uid,
		Word: req.Word,
		SoundSource: req.SoundSource,
		Ipa: req.Ipa,
		},
	}

	res, err := lh.Client.UpdateWord(context.Background(), grpcReq); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return }

	ref := ProtoToWord(res)
			
	c.JSON(http.StatusOK, ref)
}


func(lh *LanguageHandler) PostReferentSymbolLink(c *gin.Context){
	var req LinkSymbolToReferentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &pb.LinkSymbolToReferentRequest{
		ReferentUid: req.ReferentUid,
		SymbolUid: req.SymbolUid,
	}

	res, err := lh.Client.LinkSymbolToReferent(context.Background(), grpcReq); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return }
			
	c.JSON(http.StatusOK, res)
}

func(lh *LanguageHandler) PostSymbolWordLink(c *gin.Context){
	var req LinkWordToSymbolRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &pb.LinkWordToSymbolRequest{
		WordUid: req.WordUid,
		SymbolUid: req.SymbolUid,
	}

	res, err := lh.Client.LinkWordToSymbol(context.Background(), grpcReq); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return }
			
	c.JSON(http.StatusOK, res)
}

func(lh *LanguageHandler) PostWordLemmaLink(c *gin.Context){
	var req SetWordAsLemmaRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	grpcReq := &pb.SetSymbolLemmaRequest{
		WordUid: req.WordUid,
		SymbolUid: req.SymbolUid,
	}

	res, err := lh.Client.SetSymbolLemma(context.Background(), grpcReq); if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return }
			
	c.JSON(http.StatusOK, res)
}

