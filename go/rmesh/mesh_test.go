package rmesh

import (
	"context"
	"testing"
)

func TestClient(t *testing.T) {
	_, err := NewMeshRPC()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSubscribe(t *testing.T) {
	r, err := NewMeshRPC()
	if err != nil {
		t.Fatal(err)
	}

	_, err = r.Subscribe(context.Background(), "/channel", []string{"/in/delhi"})
	if err != nil {
		t.Fatal(err)
	}
}
