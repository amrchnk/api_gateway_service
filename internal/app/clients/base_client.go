package clients

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"log"
	"time"
)

func GRPCClientConnection(ctx context.Context, dsn string) *grpc.ClientConn{
	connectionCtx,cancel:=context.WithTimeout(ctx,time.Second*5)
	defer cancel()

	conn,err:=grpc.DialContext(
		connectionCtx,
		dsn,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			PermitWithoutStream: true,
			Timeout:             30 * time.Second,
		}),
	)
	if err != nil {
		log.Fatalf("failed to create GRPC connection: %s", err)
	}
	return conn
}

