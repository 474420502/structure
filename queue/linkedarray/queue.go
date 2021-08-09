package arrayqueue

import (
	"fmt"
)

// ArrayQueue
type ArrayQueue struct {
	size  int64
	data  []interface{}
	cap   int64
	start int64
	end   int64
}

func New() *ArrayQueue {
	return &ArrayQueue{
		size:  0,
		start: 0,
		end:   0,
		cap:   8,
		data:  make([]interface{}, 8),
	}
}

func (queue *ArrayQueue) grow() {
	if queue.size >= queue.cap {
		// 扩容
		cap := queue.cap << 1
		growData := make([]interface{}, cap)
		copy(growData, queue.data[queue.start:])
		copy(growData[queue.cap-queue.start:], queue.data[0:queue.start])
		queue.data = growData
		queue.start = 0
		queue.end = queue.cap
		queue.cap = cap
	}
}

func (queue *ArrayQueue) slimming() {
	if queue.size <= (queue.cap >> 1) {
		cap := (queue.cap - queue.cap>>2)
		growData := make([]interface{}, cap)

		if queue.start > queue.end {
			copy(growData, queue.data[queue.start:])
			copy(growData[queue.cap-queue.start:], queue.data[0:queue.end])
		} else {
			copy(growData, queue.data[queue.start:queue.end])
		}

		queue.data = growData
		queue.start = 0
		queue.end = queue.size
		queue.cap = cap
	}
}

func (queue *ArrayQueue) PushBack(value interface{}) {
	queue.grow()

	queue.data[queue.end] = value
	queue.size++
	queue.end++
	if queue.end >= queue.cap {
		queue.end = 0
	}
}

func (queue *ArrayQueue) PushFront(value interface{}) {
	queue.grow()

	queue.start = queue.start - 1
	if queue.start < 0 {
		queue.start = queue.cap - 1
	}

	queue.data[queue.start] = value
	queue.size++
}

func (queue *ArrayQueue) Index(idx int64) interface{} {
	if idx < queue.size {
		idx = queue.start + idx
		if idx >= queue.cap {
			idx = idx - queue.cap
		}

		return queue.data[idx]
	}

	panic(fmt.Errorf("out of size(%d): %d", queue.size, idx))
}

func (queue *ArrayQueue) Front() interface{} {
	return queue.data[queue.start]
}

func (queue *ArrayQueue) Back() interface{} {
	idx := queue.end - 1
	if idx < 0 {
		idx = queue.cap - 1
	}
	return queue.data[idx]
}

func (queue *ArrayQueue) Size() int64 {
	return queue.size
}

func (queue *ArrayQueue) PopBack() interface{} {
	if queue.size == 0 {
		return nil
	}

	if queue.end == 0 {
		queue.end = queue.cap - 1
	} else {
		queue.end--
	}

	tail := queue.data[queue.end]
	queue.size--

	queue.slimming()
	return tail
}

func (queue *ArrayQueue) PopFront() interface{} {
	if queue.size == 0 {
		return nil
	}
	head := queue.data[queue.start]
	queue.start++
	queue.size--
	if queue.start >= queue.cap {
		queue.start = 0
	}

	queue.slimming()
	return head
}

func (queue *ArrayQueue) Traverse(do func(idx int64, value interface{}) bool) {

	var idx int64 = 0
	var count int64 = 0
	for i := queue.start; i < queue.cap && count < queue.size; i++ {
		if !do(idx, queue.data[i]) {
			return
		}
		count++
		idx++
	}

	if count < queue.size {
		for i := int64(0); i < queue.end; i++ {
			if !do(idx, queue.data[i]) {
				return
			}
			count++
			idx++
		}
	}
}

// Values 返回所有队列数据. 遍历不推荐这样
func (queue *ArrayQueue) Values() []interface{} {
	var result []interface{}

	var count int64 = 0
	for i := queue.start; i < queue.cap && count < queue.size; i++ {
		result = append(result, queue.data[i])
		count++
	}

	if count < queue.size {
		for i := int64(0); i < queue.end; i++ {
			result = append(result, queue.data[i])
		}
		count++
	}

	return result
}
