package rmesh

import (
	"context"
	"fmt"

	rpc "github.com/pravahio/cross-lang-mesh-rpc/go/rpc"
	"google.golang.org/grpc"
)

type MeshRPC struct {
	stub rpc.MeshClient
}

func NewMeshRPC() (*MeshRPC, error) {
	c, err := grpc.Dial("127.0.0.1:5555", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	fmt.Println("no err")

	return &MeshRPC{
		stub: rpc.NewMeshClient(c),
	}, nil
}

func (m *MeshRPC) Subscribe(ctx context.Context, channel string, geospace []string) (rpc.Mesh_SubscribeClient, error) {
	pi := &rpc.PeerTopicInfo{
		Topics: m.getTopicFromGeospace(channel, geospace),
	}

	sub, err := m.stub.Subscribe(ctx, pi)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (m *MeshRPC) getTopicFromGeospace(channel string, geospace []string) []string {
	out := make([]string, len(geospace))

	for _, g := range geospace {
		out = append(out, channel+g)
	}

	return out
}
