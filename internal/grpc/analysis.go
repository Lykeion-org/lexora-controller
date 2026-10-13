package grpc

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/lexora"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/shared"
	"github.com/Lykeion-org/lexora-controller/internal/models"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// JobPublisher puts a serialized job on the analyser's job queue.
type JobPublisher interface {
	Publish(ctx context.Context, body []byte) error
}

const maxAnalysisTextLength = 100_000

func (s *lexoraServer) AnalyzeTextAndAddToSuggestions(ctx context.Context, in *lexora.AnalyzeTextAndAddToSuggestionsRequest) (*lexora.AnalyzeTextAndAddToSuggestionsResponse, error) {
	if s.jobs == nil {
		return nil, status.Error(codes.Unavailable, "text analysis queue is not connected")
	}
	text := strings.TrimSpace(in.GetText())
	if text == "" {
		return nil, status.Error(codes.InvalidArgument, "text is required")
	}
	if len(text) > maxAnalysisTextLength {
		return nil, status.Errorf(codes.InvalidArgument, "text is longer than %d bytes", maxAnalysisTextLength)
	}
	switch in.GetLanguage() {
	case shared.Language_LANGUAGE_UNSPECIFIED, shared.Language_LANGUAGE_EN:
	default:
		return nil, status.Error(codes.InvalidArgument, "only English text can be analysed for now")
	}

	jobUid := uuid.NewString()
	body, err := json.Marshal(models.AnalyzeTextRequestInternal{
		JobId:       jobUid,
		RequesterId: in.GetRequesterUid(),
		Language:    "en",
		Text:        text,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "encode job: %v", err)
	}
	if err := s.jobs.Publish(ctx, body); err != nil {
		return nil, status.Errorf(codes.Unavailable, "queue job: %v", err)
	}

	return &lexora.AnalyzeTextAndAddToSuggestionsResponse{
		JobUid:  jobUid,
		UserUid: in.GetRequesterUid(),
	}, nil
}

func (s *lexoraServer) ListAllWordSuggestion(ctx context.Context, in *lexora.ListAllWordSuggestionRequest) (*lexora.ListAllWordSuggestionResponse, error) {
	return s.dictionary.ListAllWordSuggestion(ctx, in)
}
