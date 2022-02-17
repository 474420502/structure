package compare

import "time"

type Compare[T any] func(k1, k2 T) int

type BytesType interface {
	[]byte | string
}

// BytesLenAny Bytes []byte compare
func BytesLenAny[T BytesType](k1, k2 T) int {

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

// BytesAny compare bytes
func BytesAny[T BytesType](k1, k2 T) int {
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

type DefaultAny interface {
	int | int64 | int32 | int8 | float32 | float64 | uint8 | uint | uint32 | uint64
}

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

// Bytes []byte compare
func Bytes(k1, k2 interface{}) int {
	s1 := k1.([]byte)
	s2 := k2.([]byte)

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

// Bytes []byte compare
func BytesLen(k1, k2 interface{}) int {
	s1 := k1.([]byte)
	s2 := k2.([]byte)

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

// String compare
func StringLen(k1, k2 interface{}) int {
	s1 := k1.(string)
	s2 := k2.(string)

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

// String comp
func String(k1, k2 interface{}) int {
	s1 := k1.(string)
	s2 := k2.(string)

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

// Runes []rune compare
func Runes(k1, k2 interface{}) int {
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

// RunesLen []rune compare
func RunesLen(k1, k2 interface{}) int {
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

func Int(k1, k2 interface{}) int {
	c1 := k1.(int)
	c2 := k2.(int)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Int8(k1, k2 interface{}) int {
	c1 := k1.(int8)
	c2 := k2.(int8)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Int16(k1, k2 interface{}) int {
	c1 := k1.(int16)
	c2 := k2.(int16)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Int32(k1, k2 interface{}) int {
	c1 := k1.(int32)
	c2 := k2.(int32)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Int64(k1, k2 interface{}) int {
	c1 := k1.(int64)
	c2 := k2.(int64)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func UInt(k1, k2 interface{}) int {
	c1 := k1.(uint)
	c2 := k2.(uint)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func UInt8(k1, k2 interface{}) int {
	c1 := k1.(uint8)
	c2 := k2.(uint8)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func UInt16(k1, k2 interface{}) int {
	c1 := k1.(uint16)
	c2 := k2.(uint16)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func UInt32(k1, k2 interface{}) int {
	c1 := k1.(uint32)
	c2 := k2.(uint32)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func UInt64(k1, k2 interface{}) int {
	c1 := k1.(uint64)
	c2 := k2.(uint64)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Float32(k1, k2 interface{}) int {
	c1 := k1.(float32)
	c2 := k2.(float32)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Float64(k1, k2 interface{}) int {
	c1 := k1.(float64)
	c2 := k2.(float64)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Byte(k1, k2 interface{}) int {
	c1 := k1.(byte)
	c2 := k2.(byte)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Rune(k1, k2 interface{}) int {
	c1 := k1.(rune)
	c2 := k2.(rune)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Time(k1, k2 interface{}) int {
	c1 := k1.(time.Time)
	c2 := k2.(time.Time)
	switch {
	case c1.After(c2):
		return 1
	case c1.Before(c2):
		return -1
	default:
		return 0
	}
}

// Bytes []byte compare
func BytesDesc(k2, k1 interface{}) int {
	s1 := k1.([]byte)
	s2 := k2.([]byte)

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

// Bytes []byte compare
func BytesLenDesc(k2, k1 interface{}) int {
	s1 := k1.([]byte)
	s2 := k2.([]byte)

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

// String compare
func StringLenDesc(k2, k1 interface{}) int {
	s1 := k1.(string)
	s2 := k2.(string)

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

// String comp
func StringDesc(k2, k1 interface{}) int {
	s1 := k1.(string)
	s2 := k2.(string)

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

// Runes []rune compare
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

// RunesLen []rune compare
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

func IntDesc(k2, k1 interface{}) int {
	c1 := k1.(int)
	c2 := k2.(int)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Int8Desc(k2, k1 interface{}) int {
	c1 := k1.(int8)
	c2 := k2.(int8)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Int16Desc(k2, k1 interface{}) int {
	c1 := k1.(int16)
	c2 := k2.(int16)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Int32Desc(k2, k1 interface{}) int {
	c1 := k1.(int32)
	c2 := k2.(int32)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Int64Desc(k2, k1 interface{}) int {
	c1 := k1.(int64)
	c2 := k2.(int64)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func UIntDesc(k2, k1 interface{}) int {
	c1 := k1.(uint)
	c2 := k2.(uint)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func UInt8Desc(k2, k1 interface{}) int {
	c1 := k1.(uint8)
	c2 := k2.(uint8)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func UInt16Desc(k2, k1 interface{}) int {
	c1 := k1.(uint16)
	c2 := k2.(uint16)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func UInt32Desc(k2, k1 interface{}) int {
	c1 := k1.(uint32)
	c2 := k2.(uint32)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func UInt64Desc(k2, k1 interface{}) int {
	c1 := k1.(uint64)
	c2 := k2.(uint64)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Float32Desc(k2, k1 interface{}) int {
	c1 := k1.(float32)
	c2 := k2.(float32)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func Float64Desc(k2, k1 interface{}) int {
	c1 := k1.(float64)
	c2 := k2.(float64)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func ByteDesc(k2, k1 interface{}) int {
	c1 := k1.(byte)
	c2 := k2.(byte)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func RuneDesc(k2, k1 interface{}) int {
	c1 := k1.(rune)
	c2 := k2.(rune)
	switch {
	case c1 > c2:
		return 1
	case c1 < c2:
		return -1
	default:
		return 0
	}
}

func TimeDesc(k2, k1 interface{}) int {
	c1 := k1.(time.Time)
	c2 := k2.(time.Time)
	switch {
	case c1.After(c2):
		return 1
	case c1.Before(c2):
		return -1
	default:
		return 0
	}
}
