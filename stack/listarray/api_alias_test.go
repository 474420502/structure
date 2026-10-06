package lastack

import "testing"

func TestLenAlias(t *testing.T) {
	st := New[int]()
	if st.Len() != 0 {
		t.Fatalf("empty stack Len should be 0, got %d", st.Len())
	}
	st.Push(1)
	st.Push(2)
	if st.Len() != 2 {
		t.Fatalf("Len should report 2, got %d", st.Len())
	}
}
