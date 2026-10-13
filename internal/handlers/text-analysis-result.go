package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"strings"
	"time"

	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/lexora"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/shared"
	"github.com/Lykeion-org/lexora-controller/internal/models"
)

// TextAnalyzerFeedbackHandler turns analyser results into word suggestions
// for every word that is not in the dictionary yet.
type TextAnalyzerFeedbackHandler struct {
	dictionarySvc lexora.DictionaryServiceClient
}

func NewTextAnalyzerFeedbackHandler(dictionaryService lexora.DictionaryServiceClient) *TextAnalyzerFeedbackHandler {
	return &TextAnalyzerFeedbackHandler{
		dictionarySvc: dictionaryService,
	}
}

func (h *TextAnalyzerFeedbackHandler) HandleMessage(message []byte) error {
	var result models.AnalyzeTextResponseInternal
	if err := json.Unmarshal(message, &result); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	groups := []struct {
		words      []models.WordAnalysis
		symbolType shared.SymbolType
	}{
		{result.Analysis.Objects, shared.SymbolType_SYMBOL_TYPE_OBJECT},
		{result.Analysis.Actions, shared.SymbolType_SYMBOL_TYPE_ACTION},
		{result.Analysis.Modifiers, shared.SymbolType_SYMBOL_TYPE_MODIFIER},
	}

	created, known := 0, 0
	for _, group := range groups {
		for _, word := range group.words {
			isKnown, err := h.isKnownWord(ctx, word.Text, shared.Language_LANGUAGE_EN)
			if err != nil {
				return err
			}
			if isKnown {
				known++
				continue
			}

			_, err = h.dictionarySvc.CreateWordSuggestion(ctx, &lexora.CreateWordSuggestionRequest{
				Suggestion: &lexora.WordSuggestion{
					Word:         word.Text,
					Lemma:        word.Lemma,
					Language:     shared.Language_LANGUAGE_EN,
					SymbolType:   group.symbolType,
					WordType:     wordTypeFromPos(word.Pos),
					Attributes:   convertAttributes(word.Attributes),
					JobUid:       result.JobId,
					RequesterUid: result.RequesterId,
				},
			})
			if err != nil {
				return err
			}
			created++
		}
	}

	slog.Info("Processed text analysis result", "job", result.JobId, "suggestions", created, "already_known", known)
	return nil
}

// isKnownWord reports whether some symbol in the language already has this exact word form.
func (h *TextAnalyzerFeedbackHandler) isKnownWord(ctx context.Context, text string, language shared.Language) (bool, error) {
	resp, err := h.dictionarySvc.FindReferents(ctx, &lexora.FindReferentsRequest{
		Query:      text,
		SearchMode: lexora.SearchMode_SEARCH_MODE_WORD,
		Language:   language,
	})
	if err != nil {
		return false, err
	}

	for _, referent := range resp.GetReferents() {
		for _, symbol := range referent.GetSymbols() {
			if symbol.GetLanguage() != language {
				continue
			}
			for _, w := range symbol.GetWords() {
				if strings.EqualFold(w.GetWord(), text) {
					return true, nil
				}
			}
		}
	}
	return false, nil
}

func wordTypeFromPos(pos string) shared.WordType {
	switch pos {
	case "NOUN", "PROPN":
		return shared.WordType_WORD_TYPE_NOUN
	case "VERB":
		return shared.WordType_WORD_TYPE_VERB
	case "ADJ":
		return shared.WordType_WORD_TYPE_ADJECTIVE
	case "ADV":
		return shared.WordType_WORD_TYPE_ADVERB
	default:
		return shared.WordType_WORD_TYPE_UNSPECIFIED
	}
}

func convertAttributes(a models.WordAnalysisAttributes) *shared.WordAttributes {
	return &shared.WordAttributes{
		Gender:         shared.WordGender(a.Gender),
		Number:         shared.WordNumber(a.Number),
		Unique:         a.Unique,
		Diminutive:     a.Diminutive,
		Case:           shared.WordCase(a.Case),
		Mood:           shared.WordMood(a.Mood),
		Tense:          shared.WordTense(a.Tense),
		Aspect:         shared.WordAspect(a.Aspect),
		Person:         shared.WordPerson(a.Person),
		DirectObject:   a.DirectObject,
		IndirectObject: a.IndirectObject,
		Valency:        int32(a.Valency),
		Reflexive:      a.Reflexive,
	}
}
