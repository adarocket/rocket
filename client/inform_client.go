package client

import (
	"context"
	"time"

	pb "github.com/adarocket/proto"

	"google.golang.org/grpc"
)

// InformClient is a client to call laptop service RPCs
type InformClient struct {
	service pb.InformerClient
}

// NewInformClient -
func NewInformClient(cc *grpc.ClientConn) *InformClient {
	service := pb.NewInformerClient(cc)
	return &InformClient{service}
}

// GetStatistic -
func (informClient *InformClient) GetNodeList() (resp *pb.GetNodeListResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.GetNodeListRequest{}

	return informClient.service.GetNodeList(ctx, req)
}

// GetStatistic -
func (informClient *InformClient) GetStatistic(nodeUUID string) (resp *pb.SaveStatisticRequest, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.GetStatisticRequest{
		Uuid: nodeUUID,
	}

	return informClient.service.GetStatistic(ctx, req)
}
