package handlers

import (
	"context"
	"encoding/json"

	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/lexora"
	"github.com/Lykeion-org/lexora-controller/internal/models"
	"github.com/Lykeion-org/lexora-controller/internal/queues"
)

type TextAnalyzerFeedbackHandler struct {
	feedbackQueue *queues.Consumer
	dictionarySvc  lexora.DictionaryServiceClient
}

func NewTextAnalyzerFeedbackHandler(feedbackQueue *queues.Consumer, dictionaryService lexora.DictionaryServiceClient) *TextAnalyzerFeedbackHandler {
	return &TextAnalyzerFeedbackHandler{
		feedbackQueue: feedbackQueue,
		dictionarySvc: dictionaryService,
	}
}

func (h *TextAnalyzerFeedbackHandler) handleMessage(message []byte)  error {
	var result models.AnalyzeTextResponseInternal
	
	err := json.Unmarshal(message, &result); if err != nil {
		return err
	}

	//lookup referents for objects
	for _, obj := range result.Analysis.Objects {
		grpcReq := &lexora.FindReferentsRequest{
			Query: "",
			SearchMode: lexora.SearchMode_SEARCH_MODE_ANY,
			Language: lexora.Language_LANGUAGE_UNSPECIFIED,
		}
		referents, err := h.dictionarySvc.FindReferents(context.Background(), grpcReq); if err != nil {
			return err
		}

		
	}		
}