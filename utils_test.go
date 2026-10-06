package utils

import (
	"errors"
	"testing"
)

func TestTryPanic(t *testing.T) {
	if err := TryPanic(func() {}); err != nil {
		t.Fatalf("no panic should yield nil, got %v", err)
	}

	sentinel := errors.New("boom")
	if err := TryPanic(func() { panic(sentinel) }); err != sentinel {
		t.Fatalf("error panic = %v want sentinel", err)
	}

	if err := TryPanic(func() { panic("text") }); err == nil || err.Error() != "text" {
		t.Fatalf("string panic = %v want text", err)
	}
}

func TestRangdom(t *testing.T) {
	for i := 0; i < 32; i++ {
		result := Rangdom(1, 32)
		// length is Intn(s+e)+s, so it stays within [s, s+e].
		if len(result) < 1 || len(result) > 33 {
			t.Fatalf("Rangdom returned %d bytes", len(result))
		}
	}
}

func TestExpect(t *testing.T) {
	Expect("1 2", 1, 2)
	if err := TryPanic(func() { Expect("mismatch", 1) }); err == nil {
		t.Fatal("Expect should panic on mismatch")
	}
}
