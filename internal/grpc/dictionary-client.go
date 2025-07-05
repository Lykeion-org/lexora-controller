package grpc

import (
	"log/slog"
	lyk_grpc "github.com/Lykeion-org/go-shared/pkg/grpc"
	"github.com/Lykeion-org/go-shared/pkg/grpc/generated/lexora"
	"google.golang.org/grpc"
)

type DictionaryClient struct {
	Client lexora.DictionaryServiceClient
	connection *grpc.ClientConn
}

func NewDictionaryClient(target string) *DictionaryClient {
	conn, err := lyk_grpc.NewClientConnection(target)
	if err != nil {
		slog.Warn(
			"Failed to connect to Lexora dictionary service", 
			"target", target, 
			"error",err,
		)
		return nil
	}

	client := lexora.NewDictionaryServiceClient(conn)

	return &DictionaryClient{
		Client: client,
		connection: conn,
	}
}

func (c *DictionaryClient) StopClient() {
	c.connection.Close()
	c.connection = nil
}
