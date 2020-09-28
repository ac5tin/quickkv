package grpcserver

import (
	"bytes"
	"context"
	"io/ioutil"
	"quickkv/quickkvpb"
	"quickkv/store"

	"github.com/ulikunitz/xz"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (*server) Replicate(ctx context.Context, request *quickkvpb.Data) (*emptypb.Empty, error) {
	empty := emptypb.Empty{}
	bin := request.Binary
	xr, err := xz.NewReader(bytes.NewBuffer(bin))
	if err != nil {
		return &empty, err
	}
	data, err := ioutil.ReadAll(xr)
	if err != nil {
		return &empty, err
	}
	store.STORE.Reset()
	store.STORE.LoadBinary(data)

	return &empty, nil
}
