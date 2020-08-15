package grpcserver

import (
	"context"
	"log"
	"quickkv/quickkvpb"
	"quickkv/store"
)

// IOService - run io service
func (*server) SendIO(ctx context.Context, request *quickkvpb.IoRequest) (*quickkvpb.IoResponse, error) {
	reqtype := request.Type

	var errno uint32 = 0
	var err error = nil
	switch reqtype {
	case 0:
		// read
		err = store.STORE.Load()
		if err != nil {
			log.Println(err.Error())
			errno = 1
		}
		break
	case 1:
		// write
		break
	default:
		break
	}
	response := &quickkvpb.IoResponse{
		Error: errno,
	}
	return response, err
}
