package server

import (
	"context"
	"github.com/kaspikr/kaspid/cmd/kaspiwallet/daemon/pb"
	"github.com/kaspikr/kaspid/version"
)

func (s *server) GetVersion(_ context.Context, _ *pb.GetVersionRequest) (*pb.GetVersionResponse, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return &pb.GetVersionResponse{
		Version: version.Version(),
	}, nil
}
