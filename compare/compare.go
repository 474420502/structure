// Package compare provides reusable key comparators for the ordered
// containers in this repository.
//
// # Standard contract
//
// Every ordered container expects the same comparator contract:
//
//   - a negative result means k1 < k2
//   - a positive result means k1 > k2
//   - zero means k1 == k2
//
// Any, AnyDesc, ArrayAny, ArrayLenAny, RunesDesc and RunesLenDesc follow this
// contract directly. They are accepted by every tree, set, map, priority
// queue, heap, and list in the repository.
//
// AnyEx and ArrayAnyEx are deprecated compatibility aliases that now delegate
// to Any and ArrayAny. They previously used a legacy index-style encoding
// (k1 < k2 returned 1, k1 > k2 returned 0, k1 == k2 returned -1) that the
// ordered containers no longer consume.
package compare

// Compare is the signature shared by every comparator.
type Compare[T any] func(k1, k2 T) int

// ArrayType is the ordered byte-slice and string key family.
type ArrayType interface {
	[]byte | string
}

// DefaultAny is the scalar key family supported by the generic comparators.
type DefaultAny interface {
	int | int64 | int32 | int8 | float32 | float64 | uint8 | uint | uint32 | uint64
}

// Any compares scalar keys in ascending order using the standard contract.
func Any[T DefaultAny](k1, k2 T) int {
	switch {
	case k1 > k2:
		return 1
	case k1 < k2:
		return -1
	default:
		return 0
	}
}

// AnyDesc compares scalar keys in descending order using the standard
// contract.
func AnyDesc[T DefaultAny](k1, k2 T) int {
	switch {
	case k1 > k2:
		return -1
	case k1 < k2:
		return 1
	default:
		return 0
	}
}

// AnyEx is a deprecated compatibility alias for Any.
//
// Deprecated: use Any. AnyEx previously used the legacy index-style encoding,
// which is no longer consumed by any container.
func AnyEx[T DefaultAny](k1, k2 T) int {
	return Any(k1, k2)
}

// ArrayAny compares byte slices and strings lexicographically with shorter
// values first. It follows the standard contract.
func ArrayAny[T ArrayType](k1, k2 T) int {
	switch {
	case len(k1) > len(k2):
		for i := 0; i < len(k2); i++ {
			if k1[i] != k2[i] {
				if k1[i] > k2[i] {
					return 1
				}
				return -1
			}
		}
		return 1
	case len(k1) < len(k2):
		for i := 0; i < len(k1); i++ {
			if k1[i] != k2[i] {
				if k1[i] > k2[i] {
					return 1
				}
				return -1
			}
		}
		return -1
	default:
		for i := 0; i < len(k1); i++ {
			if k1[i] != k2[i] {
				if k1[i] > k2[i] {
					return 1
				}
				return -1
			}
		}
		return 0
	}
}

// ArrayLenAny compares byte slices and strings by length first and then
// lexicographically. It follows the standard contract.
func ArrayLenAny[T ArrayType](k1, k2 T) int {
	switch {
	case len(k1) > len(k2):
		return 1
	case len(k1) < len(k2):
		return -1
	default:
		for i := 0; i < len(k1); i++ {
			if k1[i] != k2[i] {
				if k1[i] > k2[i] {
					return 1
				}
				return -1
			}
		}
		return 0
	}
}

// ArrayAnyEx is a deprecated compatibility alias for ArrayAny.
//
// Deprecated: use ArrayAny. ArrayAnyEx previously used the legacy index-style
// encoding, which is no longer consumed by any container.
func ArrayAnyEx[T ArrayType](k1, k2 T) int {
	return ArrayAny(k1, k2)
}

// RunesDesc compares rune slices in descending order using the standard
// contract.
func RunesDesc(k2, k1 interface{}) int {
	s1 := k1.([]rune)
	s2 := k2.([]rune)

	switch {
	case len(s1) > len(s2):
		for i := 0; i < len(s2); i++ {
			if s1[i] != s2[i] {
				if s1[i] > s2[i] {
					return 1
				}
				return -1
			}
		}
		return 1
	case len(s1) < len(s2):
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				if s1[i] > s2[i] {
					return 1
				}
				return -1
			}
		}
		return -1
	default:
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				if s1[i] > s2[i] {
					return 1
				}
				return -1
			}
		}
		return 0
	}
}

// RunesLenDesc compares rune slices by length first and then lexicographically,
// in descending order. It follows the standard contract.
func RunesLenDesc(k2, k1 interface{}) int {
	s1 := k1.([]rune)
	s2 := k2.([]rune)

	switch {
	case len(s1) > len(s2):
		return 1
	case len(s1) < len(s2):
		return -1
	default:
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				if s1[i] > s2[i] {
					return 1
				}
				return -1
			}
		}
		return 0
	}
}
