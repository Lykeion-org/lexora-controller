package api

import (
	pb "github.com/Lykeion-org/lexora-controller/internal/grpc/generated"
)

func ProtoToReferent(r *pb.Referent) *Referent {
	if r == nil {
		return nil
	}

	referent := &Referent{
		Uid:         r.GetUid(),
		EnReference: r.GetEnReference(),
		ImageSource: derefString(stringPtr(r.GetImageSource())),
		Symbols:     ProtoToSymbolList(r.GetSymbols()),
	}
	return referent
}

func ProtoToReferentList(r []*pb.Referent) []Referent {
	if r == nil {
		return []Referent{}
	}

	result := make([]Referent, len(r))
	for i, referent := range r {
		if referent != nil {
			result[i] = *ProtoToReferent(referent)
		}
	}
	return result
}

func ProtoToSymbol(r *pb.Symbol) *Symbol {
	if r == nil {
		return nil
	}

	symbol := &Symbol{
		Uid:        r.GetUid(),
		Language:   int(r.GetLanguage()),
		SymbolType: int(r.GetSymbolType()),
	}

	if r.Lemma != nil {
		symbol.Lemma = *ProtoToWord(r.GetLemma())
	}

	if r.Words != nil {
		symbol.Words = ProtoToWordList(r.GetWords())
	} else {
		symbol.Words = []Word{}
	}

	return symbol
}

func ProtoToSymbolList(r []*pb.Symbol) []Symbol {
	if r == nil {
		return []Symbol{}
	}

	symbols := make([]Symbol, len(r))
	for i, symbol := range r {
		if symbol != nil {
			symbols[i] = *ProtoToSymbol(symbol)
		}
	}
	return symbols
}

func ProtoToWord(req *pb.Word) *Word {
	if req == nil {
		return nil
	}

	word := &Word{
		Uid:         req.GetUid(),
		Word:        req.GetWord(),
		Ipa:         req.GetIpa(),
		SoundSource: req.GetSoundSource(),
	}

	if req.Attributes != nil {
		word.Attributes = WordAttributes{
			Gender:         int(req.GetAttributes().GetGender()),
			Number:         int(req.GetAttributes().GetNumber()),
			Unique:         req.GetAttributes().GetUnique(),
			Diminutive:     req.GetAttributes().GetDiminutive(),
			Case:           int(req.GetAttributes().GetCase()),
			Mood:           int(req.GetAttributes().GetMood()),
			Tense:          int(req.GetAttributes().GetTense()),
			Aspect:         int(req.GetAttributes().GetAspect()),
			Person:         int(req.GetAttributes().GetPerson()),
			DirectObject:   req.GetAttributes().GetDirectObject(),
			IndirectObject: req.GetAttributes().GetIndirectObject(),
			Valency:        int(req.GetAttributes().GetValency()),
			Reflexive:      req.GetAttributes().GetReflexive(),
		}
	}

	return word
}

func ProtoToWordList(words []*pb.Word) []Word {
	if words == nil {
		return []Word{}
	}

	result := make([]Word, len(words))
	for i, word := range words {
		if word != nil {
			result[i] = *ProtoToWord(word)
		}
	}
	return result
}

func derefString(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}

func stringPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
