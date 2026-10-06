package bloom

import (
	"hash/fnv"
	"testing"
)

// TestFnv164MatchesStdlib guards the persisted bit layout: fnv164 must produce
// exactly the same values as the hash/fnv.New64 that earlier versions used.
func TestFnv164MatchesStdlib(t *testing.T) {
	inputs := [][]byte{
		nil,
		[]byte(""),
		[]byte("a"),
		[]byte("hello world"),
		{0x00, 0x01, 0x02, 0xff, 0xfe},
	}
	for _, in := range inputs {
		h := fnv.New64()
		_, _ = h.Write(in)
		if got, want := fnv164(in), h.Sum64(); got != want {
			t.Fatalf("fnv164(%q)=%d, stdlib=%d", in, got, want)
		}
	}
}
