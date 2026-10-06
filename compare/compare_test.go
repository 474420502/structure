package compare

import "testing"

func sign(v int) int {
	switch {
	case v < 0:
		return -1
	case v > 0:
		return 1
	default:
		return 0
	}
}

func TestScalarContract(t *testing.T) {
	for _, tc := range []struct {
		a, b int
		want int
	}{
		{1, 2, -1},
		{2, 1, 1},
		{2, 2, 0},
		{-5, 5, -1},
	} {
		if got := sign(Any(tc.a, tc.b)); got != tc.want {
			t.Fatalf("Any(%d,%d) sign = %d, want %d", tc.a, tc.b, got, tc.want)
		}
		if got := sign(AnyEx(tc.a, tc.b)); got != tc.want {
			t.Fatalf("AnyEx(%d,%d) sign = %d, want %d", tc.a, tc.b, got, tc.want)
		}
		if got := sign(AnyDesc(tc.a, tc.b)); got != -tc.want {
			t.Fatalf("AnyDesc(%d,%d) sign = %d, want %d", tc.a, tc.b, got, -tc.want)
		}
	}
}

func TestArrayContract(t *testing.T) {
	for _, tc := range []struct {
		a, b string
		want int
	}{
		{"a", "b", -1},
		{"b", "a", 1},
		{"a", "aa", -1},
		{"aa", "a", 1},
		{"a", "a", 0},
		{"-5", "1", -1},
		{"abc", "abd", -1},
		{"abd", "abcx", 1},
		{"bcd", "axy", 1},
		{"a", "ba", -1},
		{"ba", "a", 1},
	} {
		if got := sign(ArrayAny(tc.a, tc.b)); got != tc.want {
			t.Fatalf("ArrayAny(%q,%q) sign = %d, want %d", tc.a, tc.b, got, tc.want)
		}
		if got := sign(ArrayAnyEx(tc.a, tc.b)); got != tc.want {
			t.Fatalf("ArrayAnyEx(%q,%q) sign = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestArrayLenContract(t *testing.T) {
	for _, tc := range []struct {
		a, b string
		want int
	}{
		{"aa", "b", 1},
		{"b", "aa", -1},
		{"ab", "ba", -1},
		{"ab", "aa", 1},
		{"same", "same", 0},
	} {
		if got := sign(ArrayLenAny(tc.a, tc.b)); got != tc.want {
			t.Fatalf("ArrayLenAny(%q,%q) sign = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func runes(s string) []rune { return []rune(s) }

func TestRunesContract(t *testing.T) {
	// RunesDesc is a descending comparison: it returns a positive value when
	// the first argument is ordered before the second.
	if got := sign(RunesDesc(runes("b"), runes("a"))); got >= 0 {
		t.Fatalf("RunesDesc(b,a) sign = %d, want negative", got)
	}
	if got := sign(RunesDesc(runes("a"), runes("b"))); got <= 0 {
		t.Fatalf("RunesDesc(a,b) sign = %d, want positive", got)
	}
	if got := sign(RunesDesc(runes("hello"), runes("world"))); got <= 0 {
		t.Fatalf("RunesDesc(hello,world) sign = %d, want positive", got)
	}
	if got := sign(RunesDesc(runes("same"), runes("same"))); got != 0 {
		t.Fatalf("RunesDesc equal = %d, want 0", got)
	}
	// Prefix and differing-content branches.
	if got := sign(RunesDesc(runes("a"), runes("ab"))); got <= 0 {
		t.Fatalf("RunesDesc(a,ab) = %d, want positive", got)
	}
	if got := sign(RunesDesc(runes("ab"), runes("a"))); got >= 0 {
		t.Fatalf("RunesDesc(ab,a) = %d, want negative", got)
	}
	if got := sign(RunesDesc(runes("b"), runes("ac"))); got > 0 {
		t.Fatalf("RunesDesc(b,ac) = %d, want negative", got)
	}

	// RunesLenDesc orders longer rune slices first.
	if got := sign(RunesLenDesc(runes("aa"), runes("b"))); got >= 0 {
		t.Fatalf("RunesLenDesc longer-first failed: %d", got)
	}
	if got := sign(RunesLenDesc(runes("b"), runes("aa"))); got <= 0 {
		t.Fatalf("RunesLenDesc shorter-first failed: %d", got)
	}
	if got := sign(RunesLenDesc(runes("ba"), runes("ab"))); got >= 0 {
		t.Fatalf("RunesLenDesc equal-length descending failed: %d", got)
	}
	if got := sign(RunesLenDesc(runes("ab"), runes("ab"))); got != 0 {
		t.Fatalf("RunesLenDesc equal = %d, want 0", got)
	}
}

func TestDeprecatedAliasesMatchStandard(t *testing.T) {
	for i := -3; i <= 3; i++ {
		for j := -3; j <= 3; j++ {
			if sign(AnyEx(i, j)) != sign(Any(i, j)) {
				t.Fatalf("AnyEx != Any for (%d,%d)", i, j)
			}
			if sign(ArrayAnyEx(string(rune('a'+i)), string(rune('a'+j)))) != sign(ArrayAny(string(rune('a'+i)), string(rune('a'+j)))) {
				t.Fatalf("ArrayAnyEx != ArrayAny for (%d,%d)", i, j)
			}
		}
	}
}
