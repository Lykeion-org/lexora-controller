package grpc

import (
	"context"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/lexora"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/shared"
)

func (s *lexoraServer) GetReferent(ctx context.Context, req *lexora.GetReferentRequest) (*shared.Referent, error){
	return s.dictionary.GetReferent(ctx, req)
}
func (s *lexoraServer) GetSymbol(ctx context.Context, req *lexora.GetSymbolRequest) (*shared.Symbol, error){
	return s.dictionary.GetSymbol(ctx, req)
}
func (s *lexoraServer) GetWord(ctx context.Context, req *lexora.GetWordRequest) (*shared.Word, error){
	return s.dictionary.GetWord(ctx, req)
}
func (s *lexoraServer) ListReferents(ctx context.Context, req *lexora.ListReferentsRequest) (*lexora.ListReferentsResponse, error){
	return s.dictionary.ListReferents(ctx, req)
}
func (s *lexoraServer) FindReferents(ctx context.Context, req *lexora.FindReferentsRequest) (*lexora.FindReferentsResponse, error){
	return s.dictionary.FindReferents(ctx, req)
}
func (s *lexoraServer) CreateReferent(ctx context.Context, req *lexora.CreateReferentRequest) (*shared.Referent, error){
	return s.dictionary.CreateReferent(ctx, req)
}
func (s *lexoraServer) CreateSymbol(ctx context.Context, req *lexora.CreateSymbolRequest) (*shared.Symbol, error){
	return s.dictionary.CreateSymbol(ctx, req)
}
func (s *lexoraServer) CreateWord(ctx context.Context, req *lexora.CreateWordRequest) (*shared.Word, error){
	return s.dictionary.CreateWord(ctx, req)
}
func (s *lexoraServer) CreateWordAttributes(ctx context.Context, req *lexora.CreateWordAttributesRequest) (*shared.WordAttributes, error){
	return s.dictionary.CreateWordAttributes(ctx, req)
}
func (s *lexoraServer) UpdateReferent(ctx context.Context, req *lexora.UpdateReferentRequest) (*shared.Referent, error){
	return s.dictionary.UpdateReferent(ctx, req)
}
func (s *lexoraServer) UpdateSymbol(ctx context.Context, req *lexora.UpdateSymbolRequest) (*shared.Symbol, error){
	return s.dictionary.UpdateSymbol(ctx, req)
}
func (s *lexoraServer) UpdateWord(ctx context.Context, req *lexora.UpdateWordRequest) (*shared.Word, error){
	return s.dictionary.UpdateWord(ctx, req)
}
func (s *lexoraServer) UpdateWordAttributes(ctx context.Context, req *lexora.UpdateWordAttributesRequest) (*shared.WordAttributes, error){
	return s.dictionary.UpdateWordAttributes(ctx, req)
}
func (s *lexoraServer) DeleteReferent(ctx context.Context, req *lexora.DeleteReferentRequest) (*shared.Referent, error){
	return s.dictionary.DeleteReferent(ctx, req)
}
func (s *lexoraServer) DeleteSymbol(ctx context.Context, req *lexora.DeleteSymbolRequest) (*shared.Symbol, error){
	return s.dictionary.DeleteSymbol(ctx, req)
}
func (s *lexoraServer) DeleteWord(ctx context.Context, req *lexora.DeleteWordRequest) (*shared.Word, error){
	return s.dictionary.DeleteWord(ctx, req)
}
func (s *lexoraServer) DeleteWordAttributes(ctx context.Context, req *lexora.DeleteWordAttributesRequest) (*shared.WordAttributes, error){
	return s.dictionary.DeleteWordAttributes(ctx, req)
}
func (s *lexoraServer) LinkSymbolToReferent(ctx context.Context, req *lexora.LinkSymbolToReferentRequest) (*lexora.LinkSymbolToReferentResponse, error){
	return s.dictionary.LinkSymbolToReferent(ctx, req)
}
func (s *lexoraServer) LinkWordToSymbol(ctx context.Context, req *lexora.LinkWordToSymbolRequest) (*lexora.LinkWordToSymbolResponse, error){
	return s.dictionary.LinkWordToSymbol(ctx, req)
}
func (s *lexoraServer) SetSymbolLemma(ctx context.Context, req *lexora.SetSymbolLemmaRequest) (*lexora.SetSymbolLemmaResponse, error){
	return s.dictionary.SetSymbolLemma(ctx, req)
}
func (s *lexoraServer) UnlinkSymbolFromReferent(ctx context.Context, req *lexora.UnlinkSymbolFromReferentRequest) (*lexora.UnlinkSymbolFromReferentResponse, error){
	return s.dictionary.UnlinkSymbolFromReferent(ctx, req)
}
func (s *lexoraServer) UnlinkWordFromSymbol(ctx context.Context, req *lexora.UnlinkWordFromSymbolRequest) (*lexora.UnlinkWordFromSymbolResponse, error){
	return s.dictionary.UnlinkWordFromSymbol(ctx, req)
}
func (s *lexoraServer) ListAllLinkedReferents(ctx context.Context, req *lexora.ListAllLinkedReferentsRequest) (*lexora.ListAllLinkedReferentsResponse, error){
	return s.dictionary.ListAllLinkedReferents(ctx, req)
}
func (s *lexoraServer) ListAllLinkedSymbols(ctx context.Context, req *lexora.ListAllLinkedSymbolsRequest) (*lexora.ListAllLinkedSymbolsResponse, error){
	return s.dictionary.ListAllLinkedSymbols(ctx, req)
}