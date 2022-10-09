package status

import (
	"context"
)

type StatusService struct {
	UnimplementedStatusServer
}

func NewStatusService() *StatusService {
	return &StatusService{}
}

func (s *StatusService) Status(context.Context, *StatusRequest) (*StatusReply, error) {
	return &StatusReply{}, nil
}
