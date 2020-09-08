package lib

import (
	"google.golang.org/grpc"
)

// NewClient - return new QuickKV client
func NewClient(address string) (QkvClient, error) {
	opts := grpc.WithInsecure()
	cc, err := grpc.Dial(address, opts)
	if err != nil {
		return QkvClient{}, err
	}
	return QkvClient{conn: cc}, nil
}
