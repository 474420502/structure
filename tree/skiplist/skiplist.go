package skiplist

import (
	"math/rand"
	"sync"

	"github.com/474420502/structure/compare"
)

type Node[KEY any, VALUE any] struct {
	Key     KEY
	Value   VALUE
	Forward []*Node[KEY, VALUE]
}

type Slice[KEY any, VALUE any] struct {
	Key   KEY
	Value VALUE
}

type Iterator[KEY any, VALUE any] struct {
	list  *SkipList[KEY, VALUE]
	node  *Node[KEY, VALUE]
	index int64
}

type SkipList[KEY any, VALUE any] struct {
	header   *Node[KEY, VALUE]
	level    int
	size     int64
	maxLevel int
	compare  compare.Compare[KEY]
	lock     sync.RWMutex
}

const (
	MaxLevel = 16
)

func New[KEY any, VALUE any](comp compare.Compare[KEY]) *SkipList[KEY, VALUE] {
	header := &Node[KEY, VALUE]{
		Forward: make([]*Node[KEY, VALUE], MaxLevel),
	}
	return &SkipList[KEY, VALUE]{
		header:   header,
		level:    1,
		size:     0,
		maxLevel: MaxLevel,
		compare:  comp,
	}
}

func (s *SkipList[KEY, VALUE]) randomLevel() int {
	level := 1
	for level < s.maxLevel && rand.Intn(2) == 0 {
		level++
	}
	return level
}

func (s *SkipList[KEY, VALUE]) Put(key KEY, value VALUE) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	update := make([]*Node[KEY, VALUE], MaxLevel)
	for i := 0; i < MaxLevel; i++ {
		update[i] = s.header
	}
	current := s.header

	for i := s.level - 1; i >= 0; i-- {
		for current.Forward[i] != nil && s.compare(current.Forward[i].Key, key) < 0 {
			current = current.Forward[i]
		}
		update[i] = current
	}

	if current.Forward[0] != nil && s.compare(current.Forward[0].Key, key) == 0 {
		current.Forward[0].Value = value
		return false
	}

	newLevel := s.randomLevel()
	if newLevel >= s.level {
		for i := s.level; i < newLevel; i++ {
			update[i] = s.header
		}
		s.level = newLevel + 1
	}

	if s.level > s.maxLevel {
		s.level = s.maxLevel
	}

	newNode := &Node[KEY, VALUE]{
		Key:     key,
		Value:   value,
		Forward: make([]*Node[KEY, VALUE], s.level),
	}

	for i := 0; i < newLevel; i++ {
		newNode.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = newNode
	}

	s.size++
	return true
}

func (s *SkipList[KEY, VALUE]) Get(key KEY) (VALUE, bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	current := s.header
	for i := s.level - 1; i >= 0; i-- {
		for current.Forward[i] != nil && s.compare(current.Forward[i].Key, key) < 0 {
			current = current.Forward[i]
		}
	}

	current = current.Forward[0]
	if current != nil && s.compare(current.Key, key) == 0 {
		return current.Value, true
	}
	var zero VALUE
	return zero, false
}

func (s *SkipList[KEY, VALUE]) Remove(key KEY) *Slice[KEY, VALUE] {
	s.lock.Lock()
	defer s.lock.Unlock()

	update := make([]*Node[KEY, VALUE], MaxLevel)
	for i := 0; i < MaxLevel; i++ {
		update[i] = s.header
	}
	current := s.header

	for i := s.level - 1; i >= 0; i-- {
		for current.Forward[i] != nil && s.compare(current.Forward[i].Key, key) < 0 {
			current = current.Forward[i]
		}
		update[i] = current
	}

	current = current.Forward[0]
	if current == nil || s.compare(current.Key, key) != 0 {
		return nil
	}

	for i := 0; i < s.level; i++ {
		if update[i].Forward[i] == current {
			update[i].Forward[i] = current.Forward[i]
		}
	}

	for s.level > 1 && s.header.Forward[s.level-1] == nil {
		s.level--
	}

	s.size--
	return &Slice[KEY, VALUE]{Key: current.Key, Value: current.Value}
}

func (s *SkipList[KEY, VALUE]) Size() int64 {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.size
}

func (s *SkipList[KEY, VALUE]) Iterator() *Iterator[KEY, VALUE] {
	return &Iterator[KEY, VALUE]{
		list:  s,
		node:  s.header.Forward[0],
		index: 0,
	}
}

func (s *SkipList[KEY, VALUE]) Head() *Slice[KEY, VALUE] {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.header.Forward[0] == nil {
		return nil
	}
	return &Slice[KEY, VALUE]{
		Key:   s.header.Forward[0].Key,
		Value: s.header.Forward[0].Value,
	}
}

func (s *SkipList[KEY, VALUE]) Tail() *Slice[KEY, VALUE] {
	s.lock.Lock()
	defer s.lock.Unlock()

	current := s.header.Forward[0]
	if current == nil {
		return nil
	}
	for current.Forward[0] != nil {
		current = current.Forward[0]
	}
	return &Slice[KEY, VALUE]{
		Key:   current.Key,
		Value: current.Value,
	}
}

func (s *SkipList[KEY, VALUE]) Index(i int64) *Slice[KEY, VALUE] {
	s.lock.Lock()
	defer s.lock.Unlock()

	if i < 0 || i >= s.size {
		return nil
	}

	current := s.header.Forward[0]
	for i > 0 {
		current = current.Forward[0]
		i--
	}
	return &Slice[KEY, VALUE]{
		Key:   current.Key,
		Value: current.Value,
	}
}

func (s *SkipList[KEY, VALUE]) IndexOf(key KEY) int64 {
	s.lock.Lock()
	defer s.lock.Unlock()

	var index int64 = 0
	current := s.header.Forward[0]
	for current != nil {
		if s.compare(current.Key, key) == 0 {
			return index
		}
		current = current.Forward[0]
		index++
	}
	return -1
}

func (s *SkipList[KEY, VALUE]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.header = &Node[KEY, VALUE]{
		Forward: make([]*Node[KEY, VALUE], MaxLevel),
	}
	s.level = 1
	s.size = 0
}

func (s *SkipList[KEY, VALUE]) Height() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.level
}

func (s *SkipList[KEY, VALUE]) Traverse(every func(s *Slice[KEY, VALUE]) bool) {
	s.lock.Lock()
	defer s.lock.Unlock()

	current := s.header.Forward[0]
	for current != nil {
		if !every(&Slice[KEY, VALUE]{Key: current.Key, Value: current.Value}) {
			break
		}
		current = current.Forward[0]
	}
}

func (iter *Iterator[KEY, VALUE]) Valid() bool {
	return iter.node != nil
}

func (iter *Iterator[KEY, VALUE]) Next() {
	if iter.node != nil {
		iter.node = iter.node.Forward[0]
		iter.index++
	}
}

func (iter *Iterator[KEY, VALUE]) Prev() {
	iter.list.lock.Lock()
	defer iter.list.lock.Unlock()

	var prev *Node[KEY, VALUE] = iter.list.header
	var current *Node[KEY, VALUE] = iter.list.header.Forward[0]

	for current != nil && current != iter.node {
		prev = current
		current = current.Forward[0]
	}

	if current == iter.node {
		if prev != iter.list.header {
			iter.node = prev
			iter.index--
		} else {
			iter.node = nil
		}
	} else if current == nil {
		iter.node = nil
	}
}

func (iter *Iterator[KEY, VALUE]) SeekToFirst() {
	iter.node = iter.list.header.Forward[0]
	iter.index = 0
}

func (iter *Iterator[KEY, VALUE]) SeekToLast() {
	iter.list.lock.Lock()
	defer iter.list.lock.Unlock()

	var last *Node[KEY, VALUE]
	current := iter.list.header.Forward[0]
	for current != nil {
		last = current
		current = current.Forward[0]
	}
	iter.node = last
	if last != nil {
		iter.index = iter.list.size - 1
	}
}

func (iter *Iterator[KEY, VALUE]) SeekGE(key KEY) bool {
	iter.list.lock.Lock()
	defer iter.list.lock.Unlock()

	current := iter.list.header
	level := iter.list.level - 1

	for i := level; i >= 0; i-- {
		for current.Forward[i] != nil && iter.list.compare(current.Forward[i].Key, key) < 0 {
			current = current.Forward[i]
		}
	}

	iter.node = current.Forward[0]
	iter.index = 0
	return iter.node != nil
}

func (iter *Iterator[KEY, VALUE]) SeekGT(key KEY) bool {
	iter.list.lock.Lock()
	defer iter.list.lock.Unlock()

	current := iter.list.header
	level := iter.list.level - 1

	for i := level; i >= 0; i-- {
		for current.Forward[i] != nil && iter.list.compare(current.Forward[i].Key, key) <= 0 {
			current = current.Forward[i]
		}
	}

	iter.node = current.Forward[0]
	iter.index = 0
	return iter.node != nil
}

func (iter *Iterator[KEY, VALUE]) SeekLE(key KEY) bool {
	iter.list.lock.Lock()
	defer iter.list.lock.Unlock()

	var result *Node[KEY, VALUE]
	current := iter.list.header
	level := iter.list.level - 1

	for i := level; i >= 0; i-- {
		for current.Forward[i] != nil && iter.list.compare(current.Forward[i].Key, key) <= 0 {
			result = current.Forward[i]
			current = current.Forward[i]
		}
	}

	iter.node = result
	return iter.node != nil
}

func (iter *Iterator[KEY, VALUE]) SeekLT(key KEY) bool {
	iter.list.lock.Lock()
	defer iter.list.lock.Unlock()

	var result *Node[KEY, VALUE]
	current := iter.list.header
	level := iter.list.level - 1

	for i := level; i >= 0; i-- {
		for current.Forward[i] != nil && iter.list.compare(current.Forward[i].Key, key) < 0 {
			result = current.Forward[i]
			current = current.Forward[i]
		}
	}

	iter.node = result
	return iter.node != nil
}

func (iter *Iterator[KEY, VALUE]) Key() KEY {
	if iter.node == nil {
		var zero KEY
		return zero
	}
	return iter.node.Key
}

func (iter *Iterator[KEY, VALUE]) Value() VALUE {
	if iter.node == nil {
		var zero VALUE
		return zero
	}
	return iter.node.Value
}

func (iter *Iterator[KEY, VALUE]) Index() int64 {
	return iter.index
}

func (iter *Iterator[KEY, VALUE]) Clone() *Iterator[KEY, VALUE] {
	return &Iterator[KEY, VALUE]{
		list:  iter.list,
		node:  iter.node,
		index: iter.index,
	}
}

func (iter *Iterator[KEY, VALUE]) SeekByIndex(index int64) {
	iter.list.lock.Lock()
	defer iter.list.lock.Unlock()

	if index < 0 || index >= iter.list.size {
		iter.node = nil
		return
	}

	current := iter.list.header.Forward[0]
	for index > 0 {
		current = current.Forward[0]
		index--
	}
	iter.node = current
	iter.index = index
}

func (s *SkipList[KEY, VALUE]) RemoveHead() *Slice[KEY, VALUE] {
	if s.size == 0 {
		return nil
	}
	return s.Remove(s.header.Forward[0].Key)
}

func (s *SkipList[KEY, VALUE]) RemoveTail() *Slice[KEY, VALUE] {
	s.lock.Lock()
	defer s.lock.Unlock()

	if s.size == 0 {
		return nil
	}

	current := s.header.Forward[0]
	for current.Forward[0] != nil {
		current = current.Forward[0]
	}
	tailKey := current.Key

	update := make([]*Node[KEY, VALUE], MaxLevel)
	for i := 0; i < MaxLevel; i++ {
		update[i] = s.header
	}

	for i := s.level - 1; i >= 0; i-- {
		prev := s.header
		for prev.Forward[i] != nil && prev.Forward[i] != current && s.compare(prev.Forward[i].Key, tailKey) < 0 {
			prev = prev.Forward[i]
		}
		update[i] = prev
	}

	if update[0].Forward[0] != current {
		return nil
	}

	for i := 0; i < s.level; i++ {
		if update[i].Forward[i] == current {
			update[i].Forward[i] = current.Forward[i]
		}
	}

	for s.level > 1 && s.header.Forward[s.level-1] == nil {
		s.level--
	}

	s.size--
	return &Slice[KEY, VALUE]{Key: current.Key, Value: current.Value}
}

func (s *SkipList[KEY, VALUE]) Set(key KEY, value VALUE) bool {
	return s.Put(key, value)
}

func (s *SkipList[KEY, VALUE]) PutDuplicate(key KEY, value VALUE, do func(*Slice[KEY, VALUE])) bool {
	s.lock.Lock()
	defer s.lock.Unlock()

	update := make([]*Node[KEY, VALUE], MaxLevel)
	for i := 0; i < MaxLevel; i++ {
		update[i] = s.header
	}
	current := s.header

	for i := s.level - 1; i >= 0; i-- {
		for current.Forward[i] != nil && s.compare(current.Forward[i].Key, key) < 0 {
			current = current.Forward[i]
		}
		update[i] = current
	}

	if current.Forward[0] != nil && s.compare(current.Forward[0].Key, key) == 0 {
		if do != nil {
			do(&Slice[KEY, VALUE]{Key: current.Forward[0].Key, Value: current.Forward[0].Value})
		}
		current.Forward[0].Value = value
		return false
	}

	newLevel := s.randomLevel()
	if newLevel >= s.level {
		for i := s.level; i < newLevel; i++ {
			update[i] = s.header
		}
		s.level = newLevel + 1
	}

	if s.level > s.maxLevel {
		s.level = s.maxLevel
	}

	newNode := &Node[KEY, VALUE]{
		Key:     key,
		Value:   value,
		Forward: make([]*Node[KEY, VALUE], s.level),
	}

	for i := 0; i < newLevel; i++ {
		newNode.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = newNode
	}

	s.size++
	return true
}

func (s *SkipList[KEY, VALUE]) String() string {
	return "skiplist"
}

func (s *SkipList[KEY, VALUE]) RemoveIndex(index int64) *Slice[KEY, VALUE] {
	s.lock.Lock()
	defer s.lock.Unlock()

	if index < 0 || index >= s.size {
		return nil
	}

	current := s.header.Forward[0]
	for index > 0 {
		current = current.Forward[0]
		index--
	}

	if current == nil {
		return nil
	}

	key := current.Key
	update := make([]*Node[KEY, VALUE], MaxLevel)
	for i := 0; i < MaxLevel; i++ {
		update[i] = s.header
	}

	for i := s.level - 1; i >= 0; i-- {
		prev := s.header
		for prev.Forward[i] != nil && prev.Forward[i] != current && s.compare(prev.Forward[i].Key, key) < 0 {
			prev = prev.Forward[i]
		}
		update[i] = prev
	}

	if update[0].Forward[0] != current {
		return nil
	}

	for i := 0; i < s.level; i++ {
		if update[i].Forward[i] == current {
			update[i].Forward[i] = current.Forward[i]
		}
	}

	for s.level > 1 && s.header.Forward[s.level-1] == nil {
		s.level--
	}

	s.size--
	return &Slice[KEY, VALUE]{Key: current.Key, Value: current.Value}
}

func (s *SkipList[KEY, VALUE]) RemoveRangeByIndex(low, high int64) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if low < 0 {
		low = 0
	}
	if high >= s.size {
		high = s.size - 1
	}
	if low > high {
		return
	}

	count := high - low + 1

	current := s.header.Forward[0]
	for low > 0 {
		current = current.Forward[0]
		low--
	}

	for i := int64(0); i < count && current != nil; i++ {
		key := current.Key
		next := current.Forward[0]

		update := make([]*Node[KEY, VALUE], MaxLevel)
		for j := 0; j < MaxLevel; j++ {
			update[j] = s.header
		}

		marker := s.header
		for j := s.level - 1; j >= 0; j-- {
			for marker.Forward[j] != nil && marker.Forward[j] != current && s.compare(marker.Forward[j].Key, key) < 0 {
				marker = marker.Forward[j]
			}
			update[j] = marker
		}

		for j := 0; j < s.level; j++ {
			if update[j].Forward[j] == current {
				update[j].Forward[j] = current.Forward[j]
			}
		}

		for s.level > 1 && s.header.Forward[s.level-1] == nil {
			s.level--
		}

		s.size--
		current = next
	}
}

func (s *SkipList[KEY, VALUE]) Slice() []Slice[KEY, VALUE] {
	s.lock.Lock()
	defer s.lock.Unlock()

	result := make([]Slice[KEY, VALUE], 0, s.size)
	current := s.header.Forward[0]
	for current != nil {
		result = append(result, Slice[KEY, VALUE]{Key: current.Key, Value: current.Value})
		current = current.Forward[0]
	}
	return result
}

func (s *SkipList[KEY, VALUE]) Intersection(other *SkipList[KEY, VALUE]) *SkipList[KEY, VALUE] {
	result := New[KEY, VALUE](s.compare)

	s.lock.Lock()
	other.lock.Lock()
	defer s.lock.Unlock()
	defer other.lock.Unlock()

	current := s.header.Forward[0]
	for current != nil {
		found := other.searchInternal(current.Key)
		if found != nil {
			result.putInternal(current.Key, found.Value)
		}
		current = current.Forward[0]
	}
	return result
}

func (s *SkipList[KEY, VALUE]) searchInternal(key KEY) *Node[KEY, VALUE] {
	current := s.header
	for i := s.level - 1; i >= 0; i-- {
		for current.Forward[i] != nil && s.compare(current.Forward[i].Key, key) < 0 {
			current = current.Forward[i]
		}
	}
	current = current.Forward[0]
	if current != nil && s.compare(current.Key, key) == 0 {
		return current
	}
	return nil
}

func (s *SkipList[KEY, VALUE]) putInternal(key KEY, value VALUE) {
	update := make([]*Node[KEY, VALUE], MaxLevel)
	for i := 0; i < MaxLevel; i++ {
		update[i] = s.header
	}
	current := s.header

	for i := s.level - 1; i >= 0; i-- {
		for current.Forward[i] != nil && s.compare(current.Forward[i].Key, key) < 0 {
			current = current.Forward[i]
		}
		update[i] = current
	}

	if current.Forward[0] != nil && s.compare(current.Forward[0].Key, key) == 0 {
		current.Forward[0].Value = value
		return
	}

	newLevel := s.randomLevel()
	if newLevel >= s.level {
		for i := s.level; i < newLevel; i++ {
			update[i] = s.header
		}
		s.level = newLevel + 1
	}

	if s.level > s.maxLevel {
		s.level = s.maxLevel
	}

	newNode := &Node[KEY, VALUE]{
		Key:     key,
		Value:   value,
		Forward: make([]*Node[KEY, VALUE], s.level),
	}

	for i := 0; i < newLevel; i++ {
		newNode.Forward[i] = update[i].Forward[i]
		update[i].Forward[i] = newNode
	}

	s.size++
}

func (s *SkipList[KEY, VALUE]) UnionSets(other *SkipList[KEY, VALUE]) *SkipList[KEY, VALUE] {
	result := New[KEY, VALUE](s.compare)

	s.lock.Lock()
	defer s.lock.Unlock()

	current := s.header.Forward[0]
	for current != nil {
		result.Put(current.Key, current.Value)
		current = current.Forward[0]
	}

	other.lock.Lock()
	defer other.lock.Unlock()

	current = other.header.Forward[0]
	for current != nil {
		result.Put(current.Key, current.Value)
		current = current.Forward[0]
	}

	return result
}

func (s *SkipList[KEY, VALUE]) DifferenceSets(other *SkipList[KEY, VALUE]) *SkipList[KEY, VALUE] {
	result := New[KEY, VALUE](s.compare)

	s.lock.Lock()
	defer s.lock.Unlock()

	current := s.header.Forward[0]
	for current != nil {
		if other.searchInternal(current.Key) == nil {
			result.putInternal(current.Key, current.Value)
		}
		current = current.Forward[0]
	}
	return result
}

func (s *SkipList[KEY, VALUE]) Trim(low, high KEY) {
	s.lock.Lock()
	defer s.lock.Unlock()

	var keys []KEY
	current := s.header.Forward[0]
	for current != nil {
		if s.compare(current.Key, low) < 0 || s.compare(current.Key, high) > 0 {
			keys = append(keys, current.Key)
		}
		current = current.Forward[0]
	}

	for _, k := range keys {
		s.removeInternal(k)
	}
}

func (s *SkipList[KEY, VALUE]) TrimByIndex(low, high int64) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if low < 0 {
		low = 0
	}
	if high >= s.size {
		high = s.size - 1
	}
	if low > high {
		return
	}

	var keysToRemove []KEY

	current := s.header.Forward[0]
	var idx int64 = 0
	for current != nil {
		if idx < low || idx > high {
			keysToRemove = append(keysToRemove, current.Key)
		}
		current = current.Forward[0]
		idx++
	}

	for _, k := range keysToRemove {
		s.removeInternal(k)
	}
}

func (s *SkipList[KEY, VALUE]) removeInternal(key KEY) *Slice[KEY, VALUE] {
	update := make([]*Node[KEY, VALUE], MaxLevel)
	for i := 0; i < MaxLevel; i++ {
		update[i] = s.header
	}
	current := s.header

	for i := s.level - 1; i >= 0; i-- {
		for current.Forward[i] != nil && s.compare(current.Forward[i].Key, key) < 0 {
			current = current.Forward[i]
		}
		update[i] = current
	}

	current = current.Forward[0]
	if current == nil || s.compare(current.Key, key) != 0 {
		return nil
	}

	for i := 0; i < s.level; i++ {
		if update[i].Forward[i] == current {
			update[i].Forward[i] = current.Forward[i]
		}
	}

	for s.level > 1 && s.header.Forward[s.level-1] == nil {
		s.level--
	}

	s.size--
	return &Slice[KEY, VALUE]{Key: current.Key, Value: current.Value}
}

func (s *SkipList[KEY, VALUE]) Slices() []Slice[KEY, VALUE] {
	return s.Slice()
}

func NewWithMaxLevel[KEY any, VALUE any](level int, comp compare.Compare[KEY]) *SkipList[KEY, VALUE] {
	if level < 1 || level > MaxLevel {
		level = MaxLevel
	}
	header := &Node[KEY, VALUE]{
		Forward: make([]*Node[KEY, VALUE], MaxLevel),
	}
	return &SkipList[KEY, VALUE]{
		header:   header,
		level:    1,
		size:     0,
		maxLevel: level,
		compare:  comp,
	}
}
