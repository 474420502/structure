package compare

type Compare[T any] func(k1, k2 T) int

type ArrayType interface {
	[]byte | string
}

// ArrayLenAny Bytes []byte compare
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

// ArrayAny compare bytes
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

type DefaultAny interface {
	int | int64 | int32 | int8 | float32 | float64 | uint8 | uint | uint32 | uint64
}

func AnyEx[T DefaultAny](k1, k2 T) int {

	if k1 > k2 {
		return 0
	} else if k1 < k2 {
		return 1
	} else {
		return -1
	}

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

// // Bytes []byte compare
// func Bytes(k1, k2 interface{}) int {
// 	s1 := k1.([]byte)
// 	s2 := k2.([]byte)

// 	switch {
// 	case len(s1) > len(s2):
// 		for i := 0; i < len(s2); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return 1
// 	case len(s1) < len(s2):
// 		for i := 0; i < len(s1); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return -1
// 	default:
// 		for i := 0; i < len(s1); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return 0
// 	}
// }

// // Bytes []byte compare
// func BytesLen(k1, k2 interface{}) int {
// 	s1 := k1.([]byte)
// 	s2 := k2.([]byte)

// 	switch {
// 	case len(s1) > len(s2):
// 		return 1
// 	case len(s1) < len(s2):
// 		return -1
// 	default:
// 		for i := 0; i < len(s1); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return 0
// 	}
// }

// // String compare
// func StringLen(k1, k2 interface{}) int {
// 	s1 := k1.(string)
// 	s2 := k2.(string)

// 	switch {
// 	case len(s1) > len(s2):

// 		return 1
// 	case len(s1) < len(s2):

// 		return -1
// 	default:
// 		for i := 0; i < len(s1); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return 0
// 	}

// }

// // String comp
// func String(k1, k2 interface{}) int {
// 	s1 := k1.(string)
// 	s2 := k2.(string)

// 	switch {
// 	case len(s1) > len(s2):
// 		for i := 0; i < len(s2); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return 1
// 	case len(s1) < len(s2):
// 		for i := 0; i < len(s1); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return -1
// 	default:
// 		for i := 0; i < len(s1); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return 0
// 	}
// }

// // Runes []rune compare
// func Runes(k1, k2 interface{}) int {
// 	s1 := k1.([]rune)
// 	s2 := k2.([]rune)

// 	switch {
// 	case len(s1) > len(s2):
// 		for i := 0; i < len(s2); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return 1
// 	case len(s1) < len(s2):
// 		for i := 0; i < len(s1); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return -1
// 	default:
// 		for i := 0; i < len(s1); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return 0
// 	}
// }

// // RunesLen []rune compare
// func RunesLen(k1, k2 interface{}) int {
// 	s1 := k1.([]rune)
// 	s2 := k2.([]rune)

// 	switch {
// 	case len(s1) > len(s2):

// 		return 1
// 	case len(s1) < len(s2):

// 		return -1
// 	default:
// 		for i := 0; i < len(s1); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return 0
// 	}
// }

// type TimeType[T any] interface {
// 	After(T) bool
// 	Before(T) bool
// 	*T
// }

// func Time[P any, T TimeType[P]](k1, k2 T) int {
// 	switch {
// 	case k1.After(*k2):
// 		return 1
// 	case k1.Before(*k2):
// 		return -1
// 	default:
// 		return 0
// 	}
// }

// func TimeDesc[P any, T TimeType[P]](k2, k1 T) int {

// 	if k1.After(*k2) {
// 		return 1
// 	} else if k1.Before(*k2) {
// 		return -1
// 	}
// 	return 0

// }

// // Bytes []byte compare
// func BytesDesc(k2, k1 interface{}) int {
// 	s1 := k1.([]byte)
// 	s2 := k2.([]byte)

// 	switch {
// 	case len(s1) > len(s2):
// 		for i := 0; i < len(s2); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return 1
// 	case len(s1) < len(s2):
// 		for i := 0; i < len(s1); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return -1
// 	default:
// 		for i := 0; i < len(s1); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return 0
// 	}
// }

// // Bytes []byte compare
// func BytesLenDesc(k2, k1 interface{}) int {
// 	s1 := k1.([]byte)
// 	s2 := k2.([]byte)

// 	switch {
// 	case len(s1) > len(s2):
// 		return 1
// 	case len(s1) < len(s2):
// 		return -1
// 	default:
// 		for i := 0; i < len(s1); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return 0
// 	}
// }

// // String compare
// func StringLenDesc(k2, k1 interface{}) int {
// 	s1 := k1.(string)
// 	s2 := k2.(string)

// 	switch {
// 	case len(s1) > len(s2):

// 		return 1
// 	case len(s1) < len(s2):

// 		return -1
// 	default:
// 		for i := 0; i < len(s1); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return 0
// 	}

// }

// // String comp
// func StringDesc(k2, k1 interface{}) int {
// 	s1 := k1.(string)
// 	s2 := k2.(string)

// 	switch {
// 	case len(s1) > len(s2):
// 		for i := 0; i < len(s2); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return 1
// 	case len(s1) < len(s2):
// 		for i := 0; i < len(s1); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return -1
// 	default:
// 		for i := 0; i < len(s1); i++ {
// 			if s1[i] != s2[i] {
// 				if s1[i] > s2[i] {
// 					return 1
// 				}
// 				return -1
// 			}
// 		}
// 		return 0
// 	}
// }

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
