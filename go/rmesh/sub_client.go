package rmesh

import (
	rpc "github.com/pravahio/cross-lang-mesh-rpc/go/rpc"
)

type SubClient struct {
	rpc.Mesh_SubscribeClient
}
