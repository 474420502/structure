package arraylist

import (
	"fmt"
	"log"
)

type ArrayList struct {
	data    []interface{}
	headidx uint // [ nil(hdix) 1 nil(tidx) ]
	tailidx uint
	size    uint

	growthSize uint
	shrinkSize uint
}

const (
	listMaxLimit = uint(1) << 63
	listMinLimit = uint(8)
	initCap      = uint(8)
	//growthFactor = float32(2.0)  // growth by 100%
	//shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

func New() *ArrayList {
	l := &ArrayList{}
	l.data = make([]interface{}, initCap, initCap)
	l.tailidx = initCap / 2
	l.headidx = l.tailidx - 1
	// l.shrinkSize = listMinLimit
	return l
}

func (l *ArrayList) Iterator() *Iterator {
	return &Iterator{al: l, cur: 0, isInit: false}
}

func (l *ArrayList) CircularIterator() *CircularIterator {
	return &CircularIterator{al: l, cur: 0, isInit: false}
}

func (l *ArrayList) Clear() {
	l.data = make([]interface{}, 8, 8)
	l.tailidx = initCap / 2
	l.headidx = l.tailidx - 1
	l.size = 0
}

func (l *ArrayList) Empty() bool {
	return l.size == 0
}

func (l *ArrayList) Size() uint {
	return l.size
}

func (l *ArrayList) shrink() {

	if l.size <= listMinLimit {
		return
	}

	if l.size <= l.shrinkSize {
		lcap := uint(len(l.data))
		nSize := lcap - lcap>>2
		temp := make([]interface{}, nSize, nSize)

		ghidx := l.size >> 2
		gtidx := ghidx + l.size + 1
		copy(temp[ghidx+1:], l.data[l.headidx+1:l.tailidx])
		l.data = temp
		l.headidx = ghidx
		l.tailidx = gtidx

		// l.shrinkSize = l.shrinkSize - lcap>>2
		l.shrinkSize = l.size - l.size>>2

	}
}

// 后续需要优化 growth 策略
func (l *ArrayList) growth() {

	if l.size >= listMaxLimit {
		log.Panic("list size is over listMaxLimit", listMaxLimit)
	}

	lcap := uint(len(l.data))
	nSize := lcap << 1
	temp := make([]interface{}, nSize, nSize)

	ghidx := lcap / 2
	gtidx := ghidx + l.size + 1
	copy(temp[ghidx+1:], l.data[l.headidx+1:l.tailidx])
	l.data = temp
	l.headidx = ghidx
	l.tailidx = gtidx

	l.shrinkSize = l.size - l.size>>2
}

func (l *ArrayList) Push(value interface{}) {
	for l.tailidx+1 > uint(len(l.data)) {
		l.growth()
	}
	l.data[l.tailidx] = value
	l.tailidx++
	l.size += 1
}

func (l *ArrayList) PushFront(values ...interface{}) {
	psize := uint(len(values))
	for l.headidx+1-psize > listMaxLimit {
		l.growth()
		// panic("growth -1")
	}

	for _, v := range values {
		l.data[l.headidx] = v
		l.headidx--
	}
	l.size += psize
}

func (l *ArrayList) PushBack(values ...interface{}) {
	psize := uint(len(values))
	for l.tailidx+psize > uint(len(l.data)) { // [0 1 2 3 4 5 6]
		l.growth()
	}

	for _, v := range values {
		l.data[l.tailidx] = v
		l.tailidx++
	}
	l.size += psize
}

func (l *ArrayList) Front() (result interface{}) {
	if l.size != 0 {
		return l.data[l.headidx+1]
	}
	return nil
}

func (l *ArrayList) Back() (result interface{}) {
	if l.size != 0 {
		return l.data[l.tailidx-1]
	}
	return nil
}

func (l *ArrayList) PopFront() (result interface{}, found bool) {
	if l.size != 0 {
		l.size--
		l.headidx++
		result = l.data[l.headidx]
		l.shrink()
		return result, true
	}
	return nil, false
}

func (l *ArrayList) PopBack() (result interface{}, found bool) {
	if l.size != 0 {
		l.size--
		l.tailidx--
		result = l.data[l.tailidx]
		l.shrink()
		return result, true
	}
	return nil, false
}

func (l *ArrayList) Index(idx int) (interface{}, bool) {
	var uidx uint = (uint)(idx)
	if uidx < l.size {
		return l.data[uidx+l.headidx+1], true
	}
	return nil, false
}

func (l *ArrayList) Remove(idx int) (result interface{}, isfound bool) {

	if idx < 0 {
		return nil, false
	}

	var uidx = (uint)(idx)
	if uidx >= l.size {
		return nil, false
	}

	offset := l.headidx + 1 + uidx

	isfound = true
	result = l.data[offset]
	// l.data[offset] = nil // cleanup reference

	if uint(len(l.data))-l.tailidx > l.headidx {
		copy(l.data[offset:], l.data[offset+1:l.tailidx]) // shift to the left by one (slow operation, need ways to optimize this)
		l.tailidx--
	} else {
		copy(l.data[l.headidx+2:], l.data[l.headidx+1:offset])
		l.headidx++
	}

	l.size--
	l.shrink()

	return
}

func (l *ArrayList) Contains(values ...interface{}) bool {

	for _, searchValue := range values {
		found := false
		for _, element := range l.data[l.headidx+1 : l.tailidx] {
			if element == searchValue {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (l *ArrayList) Values() []interface{} {
	newElements := make([]interface{}, l.size, l.size)
	copy(newElements, l.data[l.headidx+1:l.tailidx])
	return newElements
}

func (l *ArrayList) String() string {
	return fmt.Sprintf("%v", l.Values())
}

func (l *ArrayList) Traverse(every func(interface{}) bool) {
	for i := uint(0); i < l.size; i++ {
		if !every(l.data[i+l.headidx+1]) {
			break
		}
	}
}
