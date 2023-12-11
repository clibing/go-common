package util

import "sync"

type Set[T string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64] struct {
	m            map[T]struct{}
	len          int
	sync.RWMutex // lock
}

func NewSet[T string | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64](cap int64) *Set[T] {
	temp := make(map[T]struct{}, cap)
	return &Set[T]{
		m: temp,
	}
}

// 增加一个元素
func (s *Set[T]) Add(value T) {
	s.Lock()
	defer s.Unlock()
	s.m[value] = struct{}{}
	s.len = len(s.m) // 重新计算元素数量
}

// 移除一个元素
func (s *Set[T]) Remove(item T) {
	s.Lock()
	defer s.Unlock()
	// 集合没元素直接返回
	if s.len == 0 {
		return
	}
	delete(s.m, item) // 实际从字典删除这个键
	s.len = len(s.m)  // 重新计算元素数量
}

// 查看是否存在元素
func (s *Set[T]) Has(item T) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

// 查看集合大小
func (s *Set[T]) Len() int {
	return s.len
}

// 集合是够为空
func (s *Set[T]) IsEmpty() bool {
	return s.Len() == 0 
}

// 清除集合所有元素
func (s *Set[T]) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[T]struct{}{} // 字典重新赋值
	s.len = 0              // 大小归零
}

func (s *Set[T]) List() []T {
	s.RLock()
	defer s.RUnlock()
	list := make([]T, 0, s.len)
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
