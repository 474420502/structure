package treeset

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/474420502/random"
	"github.com/474420502/structure/compare"
)

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
