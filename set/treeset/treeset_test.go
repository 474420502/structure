package treeset

import (
	"fmt"
	"strings"
	"testing"

	"github.com/474420502/random"
	"github.com/474420502/structure/compare"
)

func TestTreeSet_Add(t *testing.T) {
	type fields struct {
		tree *Tree
	}
	type args struct {
		items []interface{}
	}
	tests := []struct {
		name   string
		result string
		fields fields
		args   args
	}{
		{name: "add int", result: "(1, 3, 5)", args: args{items: []interface{}{1, 5, 3, 3, 5}}},
		{name: "add -int", result: "(-5, 1, 5, 3132)", args: args{items: []interface{}{-5, -5, 3132, 3132, 5, 1, 1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := New(compare.Int)
			set.Covers(tt.args.items...)
			if set.String() != tt.result {
				t.Error(set.String(), " != ", tt.result)
			}
		})
	}

	tests2 := []struct {
		name   string
		result string
		fields fields
		args   args
	}{
		{name: "add String 1", result: "(1, 3, 5)", args: args{items: []interface{}{"1", "5", "3", "3", "5"}}},
		{name: "add String 2", result: "(-5, 1, 3132, 5)", args: args{items: []interface{}{"-5", "-5", "3132", "3132", "5", "1", "1", "1"}}},
		{name: "add String 3", result: "(a, aa, b, bc)", args: args{items: []interface{}{"a", "b", "aa", "aa", "bc"}}},
		{name: "add String 4", result: "(他, 你, 我, 我我)", args: args{items: []interface{}{"我", "你", "他", "我", "我我"}}},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			set := New(compare.String)
			set.Covers(tt.args.items...)
			if set.String() != tt.result {
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
		tree *Tree
	}
	type args struct {
		addItems    []interface{}
		removeItems []interface{}
	}
	tests := []struct {
		name   string
		result string
		fields fields
		args   args
	}{
		{name: "remove 1", result: "()",
			args: args{
				addItems:    []interface{}{5, 7, 5, 3, 2},
				removeItems: []interface{}{5, 7, 3, 2}},
		},

		{name: "remove 2", result: "(5)",
			args: args{
				addItems:    []interface{}{5, 7, 5, 3, 2},
				removeItems: []interface{}{7, 3, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := New(compare.Int)
			set.Covers(tt.args.addItems...)
			set.Remove(tt.args.removeItems...)

			if set.String() != tt.result {
				t.Error(set.String(), " != ", tt.result)
			}
		})
	}
}

func TestTreeSet_Iterator(t *testing.T) {
	set := New(compare.Int)
	set.Covers(5, 4, 3, 5)

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
		set := New(compare.Int)
		var hashset map[int]bool = make(map[int]bool)
		for i := 0; i < 200; i++ {
			v := rand.Intn(100)
			set.Cover(v)
			hashset[v] = true
		}

		for k := range hashset {
			if !set.Contains(k) {
				panic("")
			}
		}

		for _, v := range set.Values() {
			if ok := hashset[v.(int)]; !ok {
				panic("")
			}
		}

		set.Traverse(func(v interface{}) bool {
			if ok := hashset[v.(int)]; !ok {
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

type specialItem struct {
	Key   int
	Value int
}

var specialCompare = func(a1, a2 interface{}) int {
	key1 := a1.(*specialItem).Key
	key2 := a2.(*specialItem).Key
	if key1 > key2 {
		return -1
	} else if key1 < key2 {
		return 1
	} else {
		return 0
	}
}

func TestForce2(t *testing.T) {
	rand := random.New(t.Name())

	for n := 0; n < 2000; n++ {
		set := New(specialCompare)
		var hashset map[int]int = make(map[int]int)
		for i := 0; i < 200; i++ {
			k := rand.Intn(100)
			v := rand.Intn(100)
			if set.Add(&specialItem{
				Key:   k,
				Value: v,
			}) {
				hashset[k] = v
			}
		}

		set.Traverse(func(v interface{}) bool {
			s := v.(*specialItem)
			if hashset[s.Key] != s.Value {
				panic("")
			}
			return true
		})

		var ss []interface{}
		for i := 0; i < 200; i++ {
			k := rand.Intn(100)

			ss = append(ss, &specialItem{
				Key:   k,
				Value: i,
			})
			hashset[k] = i
		}

		set.Adds(ss...)

		var is bool = true
		set.Traverse(func(v interface{}) bool {
			s := v.(*specialItem)
			if hashset[s.Key] != s.Value {
				is = false
				return false
			}
			return true
		})

		if is {
			panic("")
		}

		iter := set.Iterator()
		iter2 := set.Iterator()
		iter.SeekForNext(&specialItem{
			Key: 50,
		})
		for iter.Vaild() {
			s := iter.Value().(*specialItem)
			if s.Key > 50 {
				panic("")
			}
			iter2.Seek(s)
			if !iter2.Vaild() {
				panic("")
			}
			iter.Next()
		}

	}
}
