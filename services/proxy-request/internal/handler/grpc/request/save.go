package grpcrequest

import (
	"context"
	"fmt"

	pb_request "github.com/sazonovItas/proxy-manager/proxy-request/api/proto/pb"
)

func (rh *RequestHandler) SaveProxyRequest(
	ctx context.Context,
	r *pb_request.SaveRequest,
) (*pb_request.SaveResponse, error) {
	return &pb_request.SaveResponse{Id: fmt.Sprintf("hi, %s", r.Request.Host)}, nil
}
