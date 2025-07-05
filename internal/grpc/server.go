package grpc

import (
	"context"
	"net"
	"google.golang.org/grpc"
	lyk_grpc "github.com/Lykeion-org/go-shared/pkg/grpc"
	lexora "github.com/Lykeion-org/go-shared/pkg/grpc/generated/lexora"
	shared "github.com/Lykeion-org/go-shared/pkg/grpc/generated/shared"
)

type lexoraServer struct {
	lexora.LexoraServiceServer
	server *grpc.Server
	dictionary lexora.DictionaryServiceClient
}

func NewLexoraGrpcServer(dc lexora.DictionaryServiceClient) lyk_grpc.Server{
	return &lexoraServer{
		dictionary: dc,
	}
}

func (s *lexoraServer) StartServer(target string) error {
	grpcServer := grpc.NewServer()
	lexora.RegisterLexoraServiceServer(grpcServer, s)

	listener, err := net.Listen("tcp", target)
	if err != nil {
		return err
	}

	go func(){
		if err := grpcServer.Serve(listener); err != nil{
			panic(err)
		}
	}()

	s.server = grpcServer
	return nil
}

func (s *lexoraServer) StopServer() error {
		s.server.GracefulStop()
		return nil
}

func (s *lexoraServer) GetLanguageDefinitions(ctx context.Context, in *lexora.GetLanguageDefinitionsRequest) (*lexora.GetLanguageResponse, error){
	var languages []string = EnumNamesOrdered(shared.Language_name) 
	
	return &lexora.GetLanguageResponse{
		Languages: languages,
	}, nil
}
func (s *lexoraServer) GetSymbolDefinitions(ctx context.Context, in *lexora.GetSymbolDefinitionsRequest) (*lexora.GetSymbolDefinitionsResponse, error){
	var symbolTypes []string = EnumNamesOrdered(shared.SymbolType_name)

	return &lexora.GetSymbolDefinitionsResponse{
		SymbolTypes: symbolTypes,
	}, nil
}
func (s *lexoraServer) GetWordAttributeDefinitions(ctx context.Context, in *lexora.GetWordAttributeDefinitionsRequest) (*lexora.GetWordAttributeDefinitionsResponse, error){
	var wordTenses []string = EnumNamesOrdered(shared.WordTense_name)
	var wordGenders []string = EnumNamesOrdered(shared.WordGender_name)
	var wordNumbers []string = EnumNamesOrdered(shared.WordNumber_name)
	var wordCases []string = EnumNamesOrdered(shared.WordCase_name)
	var wordAspects []string = EnumNamesOrdered(shared.WordAspect_name)
	var wordMoods []string = EnumNamesOrdered(shared.WordMood_name)
	var wordPersons []string = EnumNamesOrdered(shared.WordPerson_name)

	return &lexora.GetWordAttributeDefinitionsResponse{
		Tense: wordTenses,
		Gender: wordGenders,
		Number: wordNumbers,
		Case: wordCases,
		Aspect: wordAspects,
		Mood: wordMoods,
		Person: wordPersons,
	}, nil
}


