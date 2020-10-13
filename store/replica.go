package store

import (
	"context"
	"log"
	"quickkv/helper"
	"quickkv/quickkvpb"
	"sync"

	uf "github.com/ac5tin/usefulgo"
	"google.golang.org/grpc"
)

var replicas = []string{}

// ListReplicas - return list of all replicas
func (s *Store) ListReplicas() []string {
	return replicas
}

// AddReplicaServer - add replica server
func (s *Store) AddReplicaServer(server string) {
	replicas = append(replicas, server)
	go s.replicateSingleServer(server)
}

// ResetReplicasList - resets the replicas list back to empty array
func (s *Store) ResetReplicasList() {
	replicas = []string{}
}

// RmReplicaServer - remove replica server
func (s *Store) RmReplicaServer(server string) {
	uf.ArrRMS(&replicas, server)
}

func (s *Store) replicateSingleServer(server string) error {
	b, err := s.GetBinary()
	if err != nil {
		return err
	}
	cmp, err := helper.Compress(&b)
	if err != nil {
		return err
	}

	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(server, opts)
	if err != nil {
		return err
	}
	client := quickkvpb.NewReplicaServiceClient(conn)
	request := &quickkvpb.Data{Binary: cmp}
	_, err = client.Replicate(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}

// ReplicateBin - exactly the same as Replicate() except it takes in the []byte as an argument instead of looking for it in the store
func (s *Store) ReplicateBin(b *[]byte) error {
	cmp, err := helper.Decompress(b)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, r := range replicas {
		wg.Add(1)
		go func(server string) {
			defer wg.Done()
			opts := grpc.WithInsecure()
			conn, err := grpc.Dial(server, opts)
			if err != nil {
				log.Println(err.Error())
			}
			client := quickkvpb.NewReplicaServiceClient(conn)
			request := &quickkvpb.Data{Binary: cmp}
			_, err = client.Replicate(context.Background(), request)
			if err != nil {
				log.Println(err.Error())
			}
		}(r)

	}
	wg.Wait()
	return nil
}

// Replicate - start replication
func (s *Store) Replicate() error {
	b, err := s.GetBinary()
	if err != nil {
		return err
	}
	cmp, err := helper.Compress(&b)
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	for _, r := range replicas {
		wg.Add(1)
		go func(server string) {
			defer wg.Done()
			opts := grpc.WithInsecure()
			conn, err := grpc.Dial(server, opts)
			if err != nil {
				log.Println(err.Error())
			}
			client := quickkvpb.NewReplicaServiceClient(conn)
			request := &quickkvpb.Data{Binary: cmp}
			_, err = client.Replicate(context.Background(), request)
			if err != nil {
				log.Println(err.Error())
			}
		}(r)

	}
	wg.Wait()
	return nil
}
