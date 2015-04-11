package server

import (
	"testing"
)

func TestVersion(t *testing.T) {
	s := &Server{}

	got := s.Version()
	if version != got {
		t.Errorf("invalid result, want=%v, got=%v", version, got)
	}
}
