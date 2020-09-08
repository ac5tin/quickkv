package lib

import "google.golang.org/grpc"

// QkvClient - QuickKV client object
type QkvClient struct {
	conn *grpc.ClientConn
}
