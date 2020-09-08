package lib

import (
	"context"
	"errors"
	"quickkv/quickkvpb"
)

// Get - get value from store
func (qc *QkvClient) Get(key string) ([]byte, error) {
	client := quickkvpb.NewStoreServiceClient(qc.conn)
	request := &quickkvpb.GetRequest{Key: key}
	resp, err := client.Get(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if resp.Status.Error == 0 {
		return resp.Data.Data, nil
	}
	return nil, errors.New("Unexpected error")
}

// Set - set value to store
func (qc *QkvClient) Set(key string, value []byte) error {
	client := quickkvpb.NewStoreServiceClient(qc.conn)
	request := &quickkvpb.SetRequest{Key: key, Data: &quickkvpb.Record{Data: value}}
	resp, err := client.Set(context.Background(), request)
	if err != nil {
		return err
	}
	if resp.Error == 0 {
		return nil
	}
	return errors.New("Unexpected error")
}
