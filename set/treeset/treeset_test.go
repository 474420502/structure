package treeset

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/474420502/random"
	"github.com/474420502/structure/compare"
)

func valuesSnapshot(set *Tree[int, int]) []int {
	values := set.Values()
	return append([]int(nil), values...)
}

func TestTreeSet_Add(t *testing.T) {
	type fields struct {
	}
	type args struct {
		items []int
	}
	tests := []struct {
		name   string
		result string
		fields fields
		args   args
	}{
		{name: "add int", result: "[1 3 5]", args: args{items: []int{1, 5, 3, 3, 5}}},
		{name: "add -int", result: "[-5 1 5 3132]", args: args{items: []int{-5, -5, 3132, 3132, 5, 1, 1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := New[int, int](compare.AnyEx[int])
			for _, item := range tt.args.items {
				set.Set(item, item)
			}
			// log.Println(set.String(), tt.result)
			if set.String() != tt.result {
				t.Error(set.String(), " != ", tt.result)
			}
		})
	}

	type argsstr struct {
		items []string
	}

	tests2 := []struct {
		name   string
		result string
		fields fields
		args   argsstr
	}{
		{name: "add String 1", result: "[1 3 5]", args: argsstr{items: []string{"1", "5", "3", "3", "5"}}},
		// 字符串的 - 字符 在其他序的后面
		{name: "add String 2", result: "[5 3132 1 -5]", args: argsstr{items: []string{"-5", "-5", "3132", "3132", "5", "1", "1", "1"}}},
		{name: "add String 3", result: "[aa bc a b]", args: argsstr{items: []string{"a", "b", "aa", "aa", "bc"}}},
		{name: "add String 4", result: "[我我 他 你 我]", args: argsstr{items: []string{"我", "你", "他", "我", "我我"}}},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			set := New[string, string](compare.ArrayAnyEx[string])
			for _, item := range tt.args.items {
				set.Set(item, item)
			}

			if set.String() != tt.result {
				log.Println(set.String(), tt.result)
				t.Error(set.String(), " != ", tt.result)
			}

			vstr := fmt.Sprint(set.Values())

			if vstr[1:len(vstr)-1] != strings.ReplaceAll(tt.result[1:len(tt.result)-1], ",", "") {
				t.Error(vstr[1:len(vstr)-1], tt.result[1:len(tt.result)-1])
			}
		})
	}
}

func TestTreeSet_Remove(t *testing.T) {
	type fields struct {
		tree *Tree[int, int]
	}
	type args struct {
		addItems    []int
		removeItems []int
	}
	tests := []struct {
		name   string
		result string
		fields fields
		args   args
	}{
		{name: "remove 1", result: "[]",
			args: args{
				addItems:    []int{5, 7, 5, 3, 2},
				removeItems: []int{5, 7, 3, 2}},
		},

		{name: "remove 2", result: "[5]",
			args: args{
				addItems:    []int{5, 7, 5, 3, 2},
				removeItems: []int{7, 3, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := New[int, int](compare.AnyEx[int])
			for _, item := range tt.args.addItems {
				set.Set(item, item)
			}

			for _, item := range tt.args.removeItems {
				set.Remove(item)
			}

			if set.String() != tt.result {
				t.Error(set.String(), " != ", tt.result)
			}
		})
	}
}

func TestTreeSet_Iterator(t *testing.T) {
	set := New[int, int](compare.AnyEx[int])

	set.Set(5, 5)
	set.Set(4, 4)
	set.Set(3, 3)
	set.Set(5, 5)

	iter := set.Iterator()
	iter.SeekToFirst()

	// if not call Next Prev will error
	// 5 4 3
	// if iter.Value() != nil {
	// 	t.Error(iter.Value())
	// }
	if !iter.Vaild() {
		panic("") // 5
	}

	if iter.Value() != 3 {
		t.Error(iter.Value())
	}

	iter.Next()
	if iter.Value() != 4 {
		t.Error(iter.Value())
	}

	iter.Next()
	if iter.Value() != 5 {
		t.Error(iter.Value())
	}

	iter.SeekToLast()
	if !iter.Vaild() {
		panic("")
	}

	if iter.Value() != 5 {
		t.Error(iter.Value())
	}

	iter.Prev()
	if iter.Value() != 4 {
		t.Error(iter.Value())
	}

	iter.SeekToFirst()
	iter.Prev()

	if iter.Vaild() {
		panic("")
	}

	iter.SeekToLast()
	iter.Next()
	if iter.Vaild() {
		panic("")
	}

}

func TestForce(t *testing.T) {
	rand := random.New(t.Name())

	for n := 0; n < 2000; n++ {
		set := New[int, int](compare.AnyEx[int])
		var hashset map[int]bool = make(map[int]bool)
		for i := 0; i < 200; i++ {
			v := rand.Intn(100)
			set.Set(v, v)
			hashset[v] = true
		}

		for k := range hashset {
			if !set.Contains(k) {
				panic("")
			}
		}

		for _, v := range set.Values() {
			if ok := hashset[v]; !ok {
				panic("")
			}
		}

		set.Traverse(func(k, v int) bool {
			if ok := hashset[v]; !ok {
				panic("")
			}
			return true
		})

		set.Clear()
		if !set.Empty() {
			panic("")
		}
	}
}

func TestTreeSet_Union(t *testing.T) {
	tests := []struct {
		name   string
		a      []int
		b      []int
		expect []int
	}{
		{"empty union empty", []int{}, []int{}, []int{}},
		{"empty union non-empty", []int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"non-empty union empty", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{"disjoint", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"overlapping", []int{1, 2, 3}, []int{3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"identical", []int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"subset left", []int{1, 2}, []int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{"subset right", []int{1, 2, 3, 4}, []int{2, 3}, []int{1, 2, 3, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setA := New[int, int](compare.AnyEx[int])
			setB := New[int, int](compare.AnyEx[int])
			for _, v := range tt.a {
				setA.Set(v, v)
			}
			for _, v := range tt.b {
				setB.Set(v, v)
			}
				beforeB := valuesSnapshot(setB)
			result := setA.Union(setB)
				if result != setA {
					t.Fatalf("union should return receiver")
				}
			if result.Size() != uint(len(tt.expect)) {
				t.Errorf("size mismatch: got %d, expect %d", result.Size(), len(tt.expect))
			}
			for _, v := range tt.expect {
				if !result.Contains(v) {
					t.Errorf("missing element %d", v)
				}
			}
				if !reflect.DeepEqual(valuesSnapshot(setB), beforeB) {
					t.Fatalf("union should not mutate other: got %v want %v", setB.Values(), beforeB)
				}
		})
	}
}

func TestTreeSet_Intersection(t *testing.T) {
	tests := []struct {
		name   string
		a      []int
		b      []int
		expect []int
	}{
		{"empty intersect empty", []int{}, []int{}, []int{}},
		{"empty intersect non-empty", []int{}, []int{1, 2, 3}, []int{}},
		{"non-empty intersect empty", []int{1, 2, 3}, []int{}, []int{}},
		{"disjoint", []int{1, 2, 3}, []int{4, 5, 6}, []int{}},
		{"overlapping", []int{1, 2, 3}, []int{3, 4, 5}, []int{3}},
		{"identical", []int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
		{"subset left", []int{1, 2}, []int{1, 2, 3, 4}, []int{1, 2}},
		{"subset right", []int{1, 2, 3, 4}, []int{2, 3}, []int{2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setA := New[int, int](compare.AnyEx[int])
			setB := New[int, int](compare.AnyEx[int])
			for _, v := range tt.a {
				setA.Set(v, v)
			}
			for _, v := range tt.b {
				setB.Set(v, v)
			}
				beforeB := valuesSnapshot(setB)
			result := setA.Intersection(setB)
				if result != setA {
					t.Fatalf("intersection should return receiver")
				}
			if result.Size() != uint(len(tt.expect)) {
				t.Errorf("size mismatch: got %d, expect %d, got %v", result.Size(), len(tt.expect), result.Values())
			}
			for _, v := range tt.expect {
				if !result.Contains(v) {
					t.Errorf("missing element %d", v)
				}
			}
				if !reflect.DeepEqual(valuesSnapshot(setB), beforeB) {
					t.Fatalf("intersection should not mutate other: got %v want %v", setB.Values(), beforeB)
				}
		})
	}
}

func TestTreeSet_Difference(t *testing.T) {
	tests := []struct {
		name   string
		a      []int
		b      []int
		expect []int
	}{
		{"empty diff empty", []int{}, []int{}, []int{}},
		{"empty diff non-empty", []int{}, []int{1, 2, 3}, []int{}},
		{"non-empty diff empty", []int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{"disjoint", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3}},
		{"overlapping", []int{1, 2, 3}, []int{3, 4, 5}, []int{1, 2}},
		{"identical", []int{1, 2, 3}, []int{1, 2, 3}, []int{}},
		{"subset left", []int{1, 2}, []int{1, 2, 3, 4}, []int{}},
		{"subset right", []int{1, 2, 3, 4}, []int{2, 3}, []int{1, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setA := New[int, int](compare.AnyEx[int])
			setB := New[int, int](compare.AnyEx[int])
			for _, v := range tt.a {
				setA.Set(v, v)
			}
			for _, v := range tt.b {
				setB.Set(v, v)
			}
				beforeB := valuesSnapshot(setB)
			result := setA.Difference(setB)
				if result != setA {
					t.Fatalf("difference should return receiver")
				}
			if result.Size() != uint(len(tt.expect)) {
				t.Errorf("size mismatch: got %d, expect %d, got %v", result.Size(), len(tt.expect), result.Values())
			}
			for _, v := range tt.expect {
				if !result.Contains(v) {
					t.Errorf("missing element %d", v)
				}
			}
				if !reflect.DeepEqual(valuesSnapshot(setB), beforeB) {
					t.Fatalf("difference should not mutate other: got %v want %v", setB.Values(), beforeB)
				}
		})
	}
}

func TestTreeSet_UnionRandom(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 1000; n++ {
		setA := New[int, int](compare.AnyEx[int])
		setB := New[int, int](compare.AnyEx[int])
		var mapA, mapB, mapUnion map[int]bool = make(map[int]bool), make(map[int]bool), make(map[int]bool)

		for i := 0; i < 100; i++ {
			v := rand.Intn(200)
			setA.Set(v, v)
			mapA[v] = true
			mapUnion[v] = true
		}
		for i := 0; i < 100; i++ {
			v := rand.Intn(200)
			setB.Set(v, v)
			mapB[v] = true
			mapUnion[v] = true
		}

		result := setA.Union(setB)
		if result.Size() != uint(len(mapUnion)) {
			t.Errorf("union size mismatch: got %d, expect %d", result.Size(), len(mapUnion))
		}

		result.Traverse(func(k, v int) bool {
			if !mapUnion[v] {
				t.Errorf("unexpected element %d in union result", v)
			}
			return true
		})
	}
}

func TestTreeSet_IntersectionRandom(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 1000; n++ {
		setA := New[int, int](compare.AnyEx[int])
		setB := New[int, int](compare.AnyEx[int])
		var mapA, mapB, mapIntersection map[int]bool = make(map[int]bool), make(map[int]bool), make(map[int]bool)

		for i := 0; i < 100; i++ {
			v := rand.Intn(200)
			setA.Set(v, v)
			mapA[v] = true
		}
		for i := 0; i < 100; i++ {
			v := rand.Intn(200)
			setB.Set(v, v)
			mapB[v] = true
		}
		for v := range mapA {
			if mapB[v] {
				mapIntersection[v] = true
			}
		}

		result := setA.Intersection(setB)
		if result.Size() != uint(len(mapIntersection)) {
			t.Errorf("intersection size mismatch: got %d, expect %d", result.Size(), len(mapIntersection))
		}

		result.Traverse(func(k, v int) bool {
			if !mapIntersection[v] {
				t.Errorf("unexpected element %d in intersection result", v)
			}
			return true
		})
	}
}

func TestTreeSet_DifferenceRandom(t *testing.T) {
	rand := random.New(t.Name())
	for n := 0; n < 1000; n++ {
		setA := New[int, int](compare.AnyEx[int])
		setB := New[int, int](compare.AnyEx[int])
		var mapA, mapB, mapDiff map[int]bool = make(map[int]bool), make(map[int]bool), make(map[int]bool)

		for i := 0; i < 100; i++ {
			v := rand.Intn(200)
			setA.Set(v, v)
			mapA[v] = true
		}
		for i := 0; i < 100; i++ {
			v := rand.Intn(200)
			setB.Set(v, v)
			mapB[v] = true
		}
		for v := range mapA {
			if !mapB[v] {
				mapDiff[v] = true
			}
		}

		result := setA.Difference(setB)
		if result.Size() != uint(len(mapDiff)) {
			t.Errorf("difference size mismatch: got %d, expect %d", result.Size(), len(mapDiff))
		}

		result.Traverse(func(k, v int) bool {
			if !mapDiff[v] {
				t.Errorf("unexpected element %d in difference result", v)
			}
			return true
		})
	}
}

func TestIterator_Clone(t *testing.T) {
	set := New[int, int](compare.AnyEx[int])
	for i := 0; i < 10; i++ {
		set.Set(i, i)
	}

	iter1 := set.Iterator()
	iter1.SeekToFirst()
	iter1.Next()
	iter1.Next()

	iter2 := iter1.Clone()
	if iter2.Value() != iter1.Value() {
		t.Errorf("clone value mismatch: got %d, expect %d", iter2.Value(), iter1.Value())
	}

	iter1.Next()
	if iter2.Value() == iter1.Value() {
		t.Errorf("clone should be independent")
	}
}
