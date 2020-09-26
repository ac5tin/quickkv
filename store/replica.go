package store

import (
	"bytes"
	"context"
	"quickkv/quickkvpb"

	uf "github.com/ac5tin/usefulgo"
	"github.com/ulikunitz/xz"
	"google.golang.org/grpc"
)

var replicas = []string{}

// AddReplicaServer - add replica server
func (s *Store) AddReplicaServer(server string) {
	replicas = append(replicas, server)
	go s.replicateSingleServer(server)
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
	var buf bytes.Buffer
	xw, err := xz.NewWriter(&buf)
	if err != nil {
		return err
	}
	if _, err := xw.Write(b); err != nil {
		return err
	}
	if err := xw.Close(); err != nil {
		return err
	}

	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(server, opts)
	if err != nil {
		return err
	}
	client := quickkvpb.NewReplicaServiceClient(conn)
	request := &quickkvpb.Data{Binary: buf.Bytes()}
	_, err = client.Replicate(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}

// Replicate - start replication
func (s *Store) Replicate() error {
	b, err := s.GetBinary()
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	xw, err := xz.NewWriter(&buf)
	if err != nil {
		return err
	}
	if _, err := xw.Write(b); err != nil {
		return err
	}
	if err := xw.Close(); err != nil {
		return err
	}

	for _, r := range replicas {
		opts := grpc.WithInsecure()
		conn, err := grpc.Dial(r, opts)
		if err != nil {
			return err
		}
		client := quickkvpb.NewReplicaServiceClient(conn)
		request := &quickkvpb.Data{Binary: buf.Bytes()}
		_, err = client.Replicate(context.Background(), request)
		if err != nil {
			return err
		}
	}
	return nil
}
