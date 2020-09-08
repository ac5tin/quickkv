package grpcserver

import (
	"context"
	"log"
	"quickkv/quickkvpb"
	"quickkv/store"
)

// Set - set value to a key
func (*server) Set(ctx context.Context, request *quickkvpb.SetRequest) (*quickkvpb.Status, error) {
	var errno uint32 = 0
	var err error = nil
	if err := store.STORE.Set(request.Key, request.Data); err != nil {
		log.Println(err.Error())
		errno = 1
	}

	response := &quickkvpb.Status{
		Error: errno,
	}
	return response, err
}

// Get - get value of a key
func (*server) Get(ctx context.Context, request *quickkvpb.GetRequest) (*quickkvpb.GetResponse, error) {
	var errno uint32 = 0
	var err error = nil
	v, err := store.STORE.Get(request.Key)
	if err != nil {
		log.Println(err.Error())
		errno = 1
	}

	response := &quickkvpb.GetResponse{
		Status: &quickkvpb.Status{Error: errno},
		Data:   &quickkvpb.Record{Data: v.([]byte)},
	}
	return response, err
}
