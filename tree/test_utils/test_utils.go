package testutils

import (
	"strconv"
)

var TestedArray = []int{5, 8, 50, 10, 11, 14, 1, 99}
var TestedBigArray = []int{988, 690, 952, 485, 371, 659, 189, 817, 108, 598, 949, 254, 377, 717, 602, 265, 97, 789, 168, 547, 654, 851, 425, 872, 418, 365, 971, 547, 726, 211, 417, 214, 331, 410, 829, 901, 775, 153, 694, 585, 935, 921, 66, 41, 890, 734, 348, 27, 756, 207, 864, 884, 1, 641, 708, 607, 306, 98, 419, 324, 177, 634, 663, 510, 223, 101, 992, 288, 759, 272, 496, 951, 286, 1, 969, 52, 806, 351, 715, 303, 103, 303, 909, 776, 649, 268, 767, 73, 762, 165, 594, 982, 486, 655, 14, 549, 556, 52, 216, 218}

var TestedBytes = func() [][]byte {
	var result [][]byte
	for _, v := range TestedArray {
		i := strconv.Itoa(v)
		result = append(result, []byte(i))
	}
	return result
}()

var TestedBytesWords = [][]byte{[]byte("world"), []byte("word"), []byte("hello")}
