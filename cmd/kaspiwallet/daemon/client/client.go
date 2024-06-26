package client

import (
	"context"
	"github.com/kaspikr/kaspid/cmd/kaspiwallet/daemon/server"
	"time"

	"github.com/pkg/errors"

	"github.com/kaspikr/kaspid/cmd/kaspiwallet/daemon/pb"
	"google.golang.org/grpc"
)

// Connect connects to the kaspiwalletd server, and returns the client instance
func Connect(address string) (pb.KaspiwalletdClient, func(), error) {
	// Connection is local, so 1 second timeout is sufficient
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(server.MaxDaemonSendMsgSize)))
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, nil, errors.New("kaspiwallet daemon is not running, start it with `kaspiwallet start-daemon`")
		}
		return nil, nil, err
	}

	return pb.NewKaspiwalletdClient(conn), func() {
		conn.Close()
	}, nil
}
