<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# linkedlist

```go
import "github.com/474420502/structure/list/linked_list"
```

## Index


- [type LinkedList](<#type-linkedlist>)
  - [func New[T any](comp compare.Compare[T]) *LinkedList[T]](<#func-new>)
  - [func (l *LinkedList[T]) Back() (result T, found bool)](<#func-linkedlistt-back>)
  - [func (l *LinkedList[T]) CircularIterator() *CircularIterator[T]](<#func-linkedlistt-circulariterator>)
  - [func (l *LinkedList[T]) Clear()](<#func-linkedlistt-clear>)
  - [func (l *LinkedList[T]) Contains(values ...T) (count int)](<#func-linkedlistt-contains>)
  - [func (l *LinkedList[T]) Empty() bool](<#func-linkedlistt-empty>)
  - [func (l *LinkedList[T]) Front() (result T, found bool)](<#func-linkedlistt-front>)
  - [func (l *LinkedList[T]) Index(idx int) (result T, ok bool)](<#func-linkedlistt-index>)
  - [func (l *LinkedList[T]) Iterator() *Iterator[T]](<#func-linkedlistt-iterator>)
  - [func (l *LinkedList[T]) PopBack() (result T, found bool)](<#func-linkedlistt-popback>)
  - [func (l *LinkedList[T]) PopFront() (result T, found bool)](<#func-linkedlistt-popfront>)
  - [func (l *LinkedList[T]) Push(value T)](<#func-linkedlistt-push>)
  - [func (l *LinkedList[T]) PushBack(values ...T)](<#func-linkedlistt-pushback>)
  - [func (l *LinkedList[T]) PushFront(values ...T)](<#func-linkedlistt-pushfront>)
  - [func (l *LinkedList[T]) Size() uint](<#func-linkedlistt-size>)
  - [func (l *LinkedList[T]) String() string](<#func-linkedlistt-string>)
  - [func (l *LinkedList[T]) Traverse(every func(value T) bool)](<#func-linkedlistt-traverse>)
  - [func (l *LinkedList[T]) Values() (result []T)](<#func-linkedlistt-values>)
- [type CircularIterator](<#type-circulariterator>)
  - [func (iter *CircularIterator[T]) InsertAfter(values ...T)](<#func-circulariteratort-insertafter>)
  - [func (iter *CircularIterator[T]) InsertBefore(values ...T)](<#func-circulariteratort-insertbefore>)
  - [func (iter *CircularIterator[T]) Move(step int)](<#func-circulariteratort-move>)
  - [func (iter *CircularIterator[T]) MoveAfter(mark *CircularIterator[T])](<#func-circulariteratort-moveafter>)
  - [func (iter *CircularIterator[T]) MoveBefore(mark *CircularIterator[T])](<#func-circulariteratort-movebefore>)
  - [func (iter *CircularIterator[T]) Next()](<#func-circulariteratort-next>)
  - [func (iter *CircularIterator[T]) Prev()](<#func-circulariteratort-prev>)
  - [func (iter *CircularIterator[T]) RemoveToNext()](<#func-circulariteratort-removetonext>)
  - [func (iter *CircularIterator[T]) RemoveToPrev()](<#func-circulariteratort-removetoprev>)
  - [func (iter *CircularIterator[T]) SetValue(v T)](<#func-circulariteratort-setvalue>)
  - [func (iter *CircularIterator[T]) Swap(other *CircularIterator[T])](<#func-circulariteratort-swap>)
  - [func (iter *CircularIterator[T]) ToHead()](<#func-circulariteratort-tohead>)
  - [func (iter *CircularIterator[T]) ToTail()](<#func-circulariteratort-totail>)
  - [func (iter *CircularIterator[T]) Vaild() bool](<#func-circulariteratort-vaild>)
  - [func (iter *CircularIterator[T]) Value() interface{}](<#func-circulariteratort-value>)
- [type Iterator](<#type-iterator>)
  - [func (iter *Iterator[T]) InsertAfter(values ...T)](<#func-iteratort-insertafter>)
  - [func (iter *Iterator[T]) InsertBefore(values ...T)](<#func-iteratort-insertbefore>)
  - [func (iter *Iterator[T]) Move(step int)](<#func-iteratort-move>)
  - [func (iter *Iterator[T]) MoveAfter(mark *Iterator[T])](<#func-iteratort-moveafter>)
  - [func (iter *Iterator[T]) MoveBefore(mark *Iterator[T])](<#func-iteratort-movebefore>)
  - [func (iter *Iterator[T]) Next()](<#func-iteratort-next>)
  - [func (iter *Iterator[T]) Prev()](<#func-iteratort-prev>)
  - [func (iter *Iterator[T]) RemoveToNext()](<#func-iteratort-removetonext>)
  - [func (iter *Iterator[T]) RemoveToPrev()](<#func-iteratort-removetoprev>)
  - [func (iter *Iterator[T]) SetValue(v T)](<#func-iteratort-setvalue>)
  - [func (iter *Iterator[T]) Swap(other *Iterator[T])](<#func-iteratort-swap>)
  - [func (iter *Iterator[T]) ToHead()](<#func-iteratort-tohead>)
  - [func (iter *Iterator[T]) ToTail()](<#func-iteratort-totail>)
  - [func (iter *Iterator[T]) Vaild() bool](<#func-iteratort-vaild>)
  - [func (iter *Iterator[T]) Value() T](<#func-iteratort-value>)

- [examples](<#examples>)

## type [CircularIterator](<https://github.com/474420502/structure/blob/master/list/linked_list/circular_iterator.go#L3-L6>)

```go
type CircularIterator[T any] struct {
    // contains filtered or unexported fields
}
```

### func \(\*CircularIterator\[T\]\) [InsertAfter](<https://github.com/474420502/structure/blob/master/list/linked_list/circular_iterator.go#L190>)

```go
func (iter *CircularIterator[T]) InsertAfter(values ...T)
```

InsertAfter insert T after the iterator\. must iter\.Vaild\(\) == true

### func \(\*CircularIterator\[T\]\) [InsertBefore](<https://github.com/474420502/structure/blob/master/list/linked_list/circular_iterator.go#L163>)

```go
func (iter *CircularIterator[T]) InsertBefore(values ...T)
```

InsertBefore insert T before the iterator\. must iter\.Vaild\(\) == true

### func \(\*CircularIterator\[T\]\) [Move](<https://github.com/474420502/structure/blob/master/list/linked_list/circular_iterator.go#L79>)

```go
func (iter *CircularIterator[T]) Move(step int)
```

Move move next\(prev\[if step \< 0\]\) by step must iter\.Vaild\(\) == true

### func \(\*CircularIterator\[T\]\) [MoveAfter](<https://github.com/474420502/structure/blob/master/list/linked_list/circular_iterator.go#L137>)

```go
func (iter *CircularIterator[T]) MoveAfter(mark *CircularIterator[T])
```

MoveAfter Move After the mark iterator\.

### func \(\*CircularIterator\[T\]\) [MoveBefore](<https://github.com/474420502/structure/blob/master/list/linked_list/circular_iterator.go#L112>)

```go
func (iter *CircularIterator[T]) MoveBefore(mark *CircularIterator[T])
```

MoveBefore Move before the mark iterator\.

### func \(\*CircularIterator\[T\]\) [Next](<https://github.com/474420502/structure/blob/master/list/linked_list/circular_iterator.go#L50>)

```go
func (iter *CircularIterator[T]) Next()
```

Next the next element

### func \(\*CircularIterator\[T\]\) [Prev](<https://github.com/474420502/structure/blob/master/list/linked_list/circular_iterator.go#L31>)

```go
func (iter *CircularIterator[T]) Prev()
```

Prev the prev element

### func \(\*CircularIterator\[T\]\) [RemoveToNext](<https://github.com/474420502/structure/blob/master/list/linked_list/circular_iterator.go#L217>)

```go
func (iter *CircularIterator[T]) RemoveToNext()
```

RemoveToNext Remove self and to Next\.

### func \(\*CircularIterator\[T\]\) [RemoveToPrev](<https://github.com/474420502/structure/blob/master/list/linked_list/circular_iterator.go#L231>)

```go
func (iter *CircularIterator[T]) RemoveToPrev()
```

RemoveToNext Remove self and to Prev\.

### func \(\*CircularIterator\[T\]\) [SetValue](<https://github.com/474420502/structure/blob/master/list/linked_list/circular_iterator.go#L13>)

```go
func (iter *CircularIterator[T]) SetValue(v T)
```

SetValue set the value of current iter

### func \(\*CircularIterator\[T\]\) [Swap](<https://github.com/474420502/structure/blob/master/list/linked_list/circular_iterator.go#L8>)

```go
func (iter *CircularIterator[T]) Swap(other *CircularIterator[T])
```

### func \(\*CircularIterator\[T\]\) [ToHead](<https://github.com/474420502/structure/blob/master/list/linked_list/circular_iterator.go#L69>)

```go
func (iter *CircularIterator[T]) ToHead()
```

ToHead to list head element

### func \(\*CircularIterator\[T\]\) [ToTail](<https://github.com/474420502/structure/blob/master/list/linked_list/circular_iterator.go#L74>)

```go
func (iter *CircularIterator[T]) ToTail()
```

ToTail to list tail element

### func \(\*CircularIterator\[T\]\) [Vaild](<https://github.com/474420502/structure/blob/master/list/linked_list/circular_iterator.go#L23>)

```go
func (iter *CircularIterator[T]) Vaild() bool
```

Vaild current is Vaild ?

### func \(\*CircularIterator\[T\]\) [Value](<https://github.com/474420502/structure/blob/master/list/linked_list/circular_iterator.go#L18>)

```go
func (iter *CircularIterator[T]) Value() interface{}
```

Value get the value of element\. must iter\.Vaild\(\) == true

## type [Iterator](<https://github.com/474420502/structure/blob/master/list/linked_list/iterator.go#L4-L7>)

Iterator an iterator is an object that enables a programmer to traverse a container

```go
type Iterator[T any] struct {
    // contains filtered or unexported fields
}
```

### func \(\*Iterator\[T\]\) [InsertAfter](<https://github.com/474420502/structure/blob/master/list/linked_list/iterator.go#L37>)

```go
func (iter *Iterator[T]) InsertAfter(values ...T)
```

InsertAfter insert T after the iterator\.  must iter\.Vaild\(\) == true

### func \(\*Iterator\[T\]\) [InsertBefore](<https://github.com/474420502/structure/blob/master/list/linked_list/iterator.go#L10>)

```go
func (iter *Iterator[T]) InsertBefore(values ...T)
```

InsertBefore insert T before the iterator\. must iter\.Vaild\(\) == true

### func \(\*Iterator\[T\]\) [Move](<https://github.com/474420502/structure/blob/master/list/linked_list/iterator.go#L154>)

```go
func (iter *Iterator[T]) Move(step int)
```

Move move next\(prev\[if step \< 0\]\) by step\. must iter\.Vaild\(\) == true

### func \(\*Iterator\[T\]\) [MoveAfter](<https://github.com/474420502/structure/blob/master/list/linked_list/iterator.go#L89>)

```go
func (iter *Iterator[T]) MoveAfter(mark *Iterator[T])
```

MoveAfter Move After the mark iterator\.

### func \(\*Iterator\[T\]\) [MoveBefore](<https://github.com/474420502/structure/blob/master/list/linked_list/iterator.go#L64>)

```go
func (iter *Iterator[T]) MoveBefore(mark *Iterator[T])
```

MoveBefore Move before the mark iterator\.

### func \(\*Iterator\[T\]\) [Next](<https://github.com/474420502/structure/blob/master/list/linked_list/iterator.go#L189>)

```go
func (iter *Iterator[T]) Next()
```

Next must iter\.Vaild\(\) == true

### func \(\*Iterator\[T\]\) [Prev](<https://github.com/474420502/structure/blob/master/list/linked_list/iterator.go#L184>)

```go
func (iter *Iterator[T]) Prev()
```

Prev must iter\.Vaild\(\) == true

### func \(\*Iterator\[T\]\) [RemoveToNext](<https://github.com/474420502/structure/blob/master/list/linked_list/iterator.go#L115>)

```go
func (iter *Iterator[T]) RemoveToNext()
```

RemoveToNext Remove self and to Next\. If iterator is removed\. return true\.  must iter\.Vaild\(\) == true

### func \(\*Iterator\[T\]\) [RemoveToPrev](<https://github.com/474420502/structure/blob/master/list/linked_list/iterator.go#L123>)

```go
func (iter *Iterator[T]) RemoveToPrev()
```

RemoveToNext Remove self and to Prev\. If iterator is removed\. return true\.  must iter\.Vaild\(\) == true

### func \(\*Iterator\[T\]\) [SetValue](<https://github.com/474420502/structure/blob/master/list/linked_list/iterator.go#L136>)

```go
func (iter *Iterator[T]) SetValue(v T)
```

SetValue  must iter\.Vaild\(\) == true

### func \(\*Iterator\[T\]\) [Swap](<https://github.com/474420502/structure/blob/master/list/linked_list/iterator.go#L131>)

```go
func (iter *Iterator[T]) Swap(other *Iterator[T])
```

Swap  must iter\.Vaild\(\) == true

### func \(\*Iterator\[T\]\) [ToHead](<https://github.com/474420502/structure/blob/master/list/linked_list/iterator.go#L194>)

```go
func (iter *Iterator[T]) ToHead()
```

ToHead\. to head and must iter\.Vaild\(\) == true

### func \(\*Iterator\[T\]\) [ToTail](<https://github.com/474420502/structure/blob/master/list/linked_list/iterator.go#L199>)

```go
func (iter *Iterator[T]) ToTail()
```

ToTail\. to tail and must iter\.Vaild\(\) == true

### func \(\*Iterator\[T\]\) [Vaild](<https://github.com/474420502/structure/blob/master/list/linked_list/iterator.go#L146>)

```go
func (iter *Iterator[T]) Vaild() bool
```

Vaild current is Vaild ?

### func \(\*Iterator\[T\]\) [Value](<https://github.com/474420502/structure/blob/master/list/linked_list/iterator.go#L141>)

```go
func (iter *Iterator[T]) Value() T
```

Value must iter\.Vaild\(\) == true

## type [LinkedList](<https://github.com/474420502/structure/blob/master/list/linked_list/linked_list.go#L25-L30>)

LinkedList struct of LinkedList

```go
type LinkedList[T any] struct {
    // contains filtered or unexported fields
}
```

### func [New](<https://github.com/474420502/structure/blob/master/list/linked_list/linked_list.go#L33>)

```go
func New[T any](comp compare.Compare[T]) *LinkedList[T]
```

New create a object of LinkedList

### func \(\*LinkedList\[T\]\) [Back](<https://github.com/474420502/structure/blob/master/list/linked_list/linked_list.go#L175>)

```go
func (l *LinkedList[T]) Back() (result T, found bool)
```

Back return the head of list

### func \(\*LinkedList\[T\]\) [CircularIterator](<https://github.com/474420502/structure/blob/master/list/linked_list/linked_list.go#L54>)

```go
func (l *LinkedList[T]) CircularIterator() *CircularIterator[T]
```

Iterator an iterator is an object that enables a programmer to traverse a container and can circulate

### func \(\*LinkedList\[T\]\) [Clear](<https://github.com/474420502/structure/blob/master/list/linked_list/linked_list.go#L59>)

```go
func (l *LinkedList[T]) Clear()
```

Clear clear the list

### func \(\*LinkedList\[T\]\) [Contains](<https://github.com/474420502/structure/blob/master/list/linked_list/linked_list.go#L231>)

```go
func (l *LinkedList[T]) Contains(values ...T) (count int)
```

Contains is the \[\]T  in list?

### func \(\*LinkedList\[T\]\) [Empty](<https://github.com/474420502/structure/blob/master/list/linked_list/linked_list.go#L68>)

```go
func (l *LinkedList[T]) Empty() bool
```

Empty   if the list is empty\, return true\. else return false

### func \(\*LinkedList\[T\]\) [Front](<https://github.com/474420502/structure/blob/master/list/linked_list/linked_list.go#L166>)

```go
func (l *LinkedList[T]) Front() (result T, found bool)
```

Front return the head of list

### func \(\*LinkedList\[T\]\) [Index](<https://github.com/474420502/structure/blob/master/list/linked_list/linked_list.go#L184>)

```go
func (l *LinkedList[T]) Index(idx int) (result T, ok bool)
```

Index slowly\. is a list\. need to move with idx step

### func \(\*LinkedList\[T\]\) [Iterator](<https://github.com/474420502/structure/blob/master/list/linked_list/linked_list.go#L47>)

```go
func (l *LinkedList[T]) Iterator() *Iterator[T]
```

Iterator  an iterator is an object that enables a programmer to traverse a container\, particularly lists

### func \(\*LinkedList\[T\]\) [PopBack](<https://github.com/474420502/structure/blob/master/list/linked_list/linked_list.go#L148>)

```go
func (l *LinkedList[T]) PopBack() (result T, found bool)
```

PopBack pop the back of the list

### func \(\*LinkedList\[T\]\) [PopFront](<https://github.com/474420502/structure/blob/master/list/linked_list/linked_list.go#L130>)

```go
func (l *LinkedList[T]) PopFront() (result T, found bool)
```

PopFront pop the head of the list

### func \(\*LinkedList\[T\]\) [Push](<https://github.com/474420502/structure/blob/master/list/linked_list/linked_list.go#L78>)

```go
func (l *LinkedList[T]) Push(value T)
```

Push Push a value to the tail of the list

### func \(\*LinkedList\[T\]\) [PushBack](<https://github.com/474420502/structure/blob/master/list/linked_list/linked_list.go#L112>)

```go
func (l *LinkedList[T]) PushBack(values ...T)
```

PushBack Push  values to the tail of the list

### func \(\*LinkedList\[T\]\) [PushFront](<https://github.com/474420502/structure/blob/master/list/linked_list/linked_list.go#L94>)

```go
func (l *LinkedList[T]) PushFront(values ...T)
```

PushFront Push values to the head of the list

### func \(\*LinkedList\[T\]\) [Size](<https://github.com/474420502/structure/blob/master/list/linked_list/linked_list.go#L73>)

```go
func (l *LinkedList[T]) Size() uint
```

Size return the size of list

### func \(\*LinkedList\[T\]\) [String](<https://github.com/474420502/structure/blob/master/list/linked_list/linked_list.go#L253>)

```go
func (l *LinkedList[T]) String() string
```

String fmt\.Sprintf\("%v"\, l\.Values\(\)\)

### func \(\*LinkedList\[T\]\) [Traverse](<https://github.com/474420502/structure/blob/master/list/linked_list/linked_list.go#L258>)

```go
func (l *LinkedList[T]) Traverse(every func(value T) bool)
```

Traverse from the list of head to the tail\. iterator can do it also\.

### func \(\*LinkedList\[T\]\) [Values](<https://github.com/474420502/structure/blob/master/list/linked_list/linked_list.go#L244>)

```go
func (l *LinkedList[T]) Values() (result []T)
```

Values get the values of list


## examples

```go
package main

import (
	"fmt"
	"log"

	"github.com/474420502/structure/compare"
	linkedlist "github.com/474420502/structure/list/linked_list"
)

func main() {
	l := linkedlist.New(compare.Any[int])
	log.Println("Push PushFront PushBack")
	l.Push(3)
	l.Push(5)
	l.Push(7)
	l.PushBack(7)
	l.PushFront(10)
	l.PushFront(2)

	log.Println("Size Empty Values String")
	log.Println(l.String()) // [2 10 3 5 7 7]
	log.Println(l.Size())   // 6
	log.Println(l.Empty())  // false
	log.Println(l.Values()) // [2 10 3 5 7 7]

	log.Println("Back Front Index Contains")
	log.Println(l.Back())         // 7 true
	log.Println(l.Front())        // 2 true
	log.Println(l.Index(1))       // 10
	log.Println(l.Contains(7))    // 2
	log.Println(l.Contains(1))    // 0
	log.Println(l.Contains(2, 7)) // 3

	log.Println("Traverse")

	l.Traverse(func(value int) bool {
		log.Println(value) // 2 10 3 5 7 7
		return true
	})

	log.Println("Remove PopBack PopFront Clear")
	// [2 10 3 5 7 7]
	l.PopBack()              // [2 10 3 5 7]
	l.PopFront()             // [10 3 5 7]
	l.Clear()                // []
	log.Println(ShowList(l)) // []

}

func main2() {
	l := linkedlist.New(compare.Any[int])
	for i := 0; i < 10; i += 2 {
		l.Push(i)
	}
	log.Println("Iterator{ToHead ToTail Value}")
	iter := l.Iterator()
	iter.ToHead()             // to head
	log.Println(iter.Value()) // value:0 index:0
	iter.ToTail()             // to tail
	log.Println(iter.Value()) // value:8 index:4

	log.Println("Iterator{Prev Next Vaild}")
	iter = l.Iterator()
	iter.ToHead()
	for iter.Vaild() {
		log.Println(iter.Value()) // 0 2 4 6 8
		iter.Next()
	}

	iter.ToTail()
	for iter.Vaild() {
		log.Println(iter.Value()) // 8 6 4 2 0
		iter.Prev()
	}

	log.Println("Iterator{SetValue Swap}")
	iter1 := l.Iterator()
	iter1.ToHead()
	iter1.SetValue(8)
	iter2 := l.Iterator()
	iter2.ToTail()
	iter2.SetValue(0)
	log.Println(ShowList(l)) // [0 2 4 6 8] -> [8 2 4 6 0]
	iter1.Swap(iter2)
	log.Println(ShowList(l)) // [8 2 4 6 0] -> [0 2 4 6 8]

	log.Println("Iterator{RemoveToNext RemoveToPrev}")
	iter = l.Iterator()
	iter.ToHead()
	iter.Next()
	iter.Next()

	iter.RemoveToNext()       // cur:6
	log.Println(iter.Value()) // 6  [0 2 6 8]
	iter.RemoveToPrev()       // cur: 2
	log.Println(iter.Value()) // 2  [0 2 8]
	log.Println(ShowList(l))  // [0 2 8]
}

func main3() {
	l := linkedlist.New(compare.Any[int])
	for i := 0; i < 10; i += 2 {
		l.Push(i)
	}
	log.Println("Iterator{ToHead ToTail Value}")
	iter := l.CircularIterator()
	iter.ToHead()             // to head
	log.Println(iter.Value()) // value:0 index:0
	iter.ToTail()             // to tail
	log.Println(iter.Value()) // value:8 index:4

	log.Println("Iterator{Prev Next Vaild}")
	iter = l.CircularIterator()

	var result []int
	var count int

	count = 0
	iter.ToHead()
	for iter.Vaild() {
		result = append(result, iter.Value())
		iter.Next()
		if iter.Value() == 0 {
			count++
			if count >= 2 {
				break
			}
		}
	}
	log.Println(result) // [2 4 6 8 0 2 4 6 8]

	result = nil
	count = 0
	iter.ToTail()
	for iter.Vaild() {
		result = append(result, iter.Value())
		iter.Prev()
		if iter.Value() == 8 {
			count++
			if count >= 2 {
				break
			}
		}
	}
	log.Println(result) // [8 6 4 2 0 8 6 4 2 0]

	log.Println("Iterator{SetValue Swap}")
	iter1 := l.CircularIterator()
	iter1.ToHead()
	iter1.SetValue(8)
	iter2 := l.CircularIterator()
	iter2.ToTail()
	iter2.SetValue(0)
	log.Println(ShowList(l)) // [0 2 4 6 8] -> [8 2 4 6 0]
	iter1.Swap(iter2)
	log.Println(ShowList(l)) // [8 2 4 6 0] -> [0 2 4 6 8]

	log.Println("Iterator{RemoveToNext RemoveToPrev}")
	iter = l.CircularIterator()

	iter.ToHead()
	iter.Next()
	iter.Next()

	// [0 2 4 6 8] 
	iter.RemoveToNext()       // cur:6
	log.Println(iter.Value()) // 6  [0 2 6 8]
	iter.RemoveToPrev()       // cur: 2
	log.Println(iter.Value()) // 2  [0 2 8]
	log.Println(ShowList(l))  // [0 2 8]
}

func ShowList[T any](a *linkedlist.LinkedList[T]) string {
	var result []T
	a.Traverse(func(value T) bool {
		result = append(result, value)
		return true
	})
	return fmt.Sprintf("%v", result)
}

```