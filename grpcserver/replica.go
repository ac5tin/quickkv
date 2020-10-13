package grpcserver

import (
	"context"
	"quickkv/helper"
	"quickkv/quickkvpb"
	"quickkv/store"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (*server) Replicate(ctx context.Context, request *quickkvpb.Data) (*emptypb.Empty, error) {
	empty := emptypb.Empty{}
	bin := request.Binary
	data, err := helper.Decompress(&bin)
	if err != nil {
		return nil, err
	}
	store.STORE.Reset()
	store.STORE.LoadBinary(data)

	return &empty, nil
}
