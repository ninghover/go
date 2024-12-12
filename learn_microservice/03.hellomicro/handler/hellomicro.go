package handler

import (
	"context"

	log "go-micro.dev/v4/logger"

	pb "hellomicro/proto"
)

type Hellomicro struct{}

func (e *Hellomicro) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received Hellomicro.Call request: %v", req)
	rsp.Msg = "Hello 你好 " + req.Name
	return nil
}
