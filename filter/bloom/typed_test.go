package bloom

import "testing"

func TestBloomTypedKeys(t *testing.T) {
	bl := New(1 << 16)

	// Every branch of withKeyBytes must encode, insert and be found again.
	cases := []interface{}{
		[]byte("raw-bytes"),
		"hello",
		int(42),
		int64(43),
		uint(44),
		uint64(45),
		int32(46),
		uint32(47),
		float64(3.5),
		float32(4.5),
		[4]byte{1, 2, 3, 4}, // default: fixed-size composite via binary.Write
	}
	for _, key := range cases {
		if bl.Add(key) {
			t.Fatalf("first Add(%T) should report a new key", key)
		}
		if !bl.Contains(key) {
			t.Fatalf("Contains(%T) should find the inserted key", key)
		}
		if !bl.Add(key) {
			t.Fatalf("second Add(%T) should report an existing key", key)
		}
	}

	if bl.Contains("not-inserted") {
		t.Fatal("Contains should be false for a missing key")
	}
	if bl.Contains(int(999)) {
		t.Fatal("Contains(int) should be false for a missing key")
	}
}
