//	Copyright 2022 Game Design Analytics
//	All Rights Reserved
//
//	Authors:
//		Eleanor McHugh (eleanor@games-with-brains.com)
//
package goGeneric

type Map[K comparable, V any] map[K]V

func MakeMap[T comparable, V any]() Map[T, V] {
	return make(Map[T,V])
}

func (m Map[K, V]) Len() (int) {
	return len(m)
}

func (m Map[K, V]) Get(k K) (V, bool) {
	v, ok := m[k]
	return v, ok
}

func (m Map[K, V]) Set(k K, v V) {
	m[k] = v
}


func (m Map[K, V]) Keys() (r []K) {
	for k, _ := range m {
		r = append(r, k)
	}
	return
}

func (m Map[K, V]) KeysMatch(o Map[K, V]) (r bool) {
	if r = len(m) == len(o); r {
		for k, _ := range m {
			if _, r = o[k]; !r {
				break
			}
		}
	}
	return
}

func (m Map[K, V]) Copy() (r Map[K, V]) {
	r = MakeMap[K, V]()
	for k, v := range m {
		r[k] = v
	}
	return
}

func (m Map[K, V]) Merge(o ...Map[K, V]) (r Map[K, V]) {
	r = m.Copy()
	for _, on := range o {
		for k, v := range on {
			r[k] = v
		}
	}
	return
}

func (m Map[K, V]) Select(k ...K) (r Map[K, V]) {
	r = MakeMap[K, V]()
	for _, k := range k {
		if v, ok := m[k]; ok {
			r[k] = v
		}
	}
	return
}
