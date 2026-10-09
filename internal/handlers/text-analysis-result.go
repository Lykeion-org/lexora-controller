package handlers

import (
	"context"
	"encoding/json"

	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/lexora"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/shared"
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
	//TODO: create word suggestions for objects that have no matching referent yet
	for _, obj := range result.Analysis.Objects {
		grpcReq := &lexora.FindReferentsRequest{
			Query: obj.Lemma,
			SearchMode: lexora.SearchMode_SEARCH_MODE_WORD,
			Language: shared.Language_LANGUAGE_UNSPECIFIED,
		}
		_, err := h.dictionarySvc.FindReferents(context.Background(), grpcReq); if err != nil {
			return err
		}
	}

	return nil
}