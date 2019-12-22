package server

import (
	"testing"
)

func Test_NewServerGRPC(t *testing.T) {
	s, err := NewServerGRPC(ServerGRPCConfig{
		Port: 5000,
	})
	if err != nil {
		t.Fatal(err)
	}

	if err := s.Listen(); err != nil {
		t.Fatal(err)
	} else {
		defer s.Close()
	}
}
