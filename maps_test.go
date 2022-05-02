//	Copyright 2022 Game Design Analytics
//	All Rights Reserved
//
//	Authors:
//		Eleanor McHugh (eleanor@games-with-brains.com)
//
package goGeneric

import (
	"sort"
	"testing"
)

func TestMaps_MakeMap(t *testing.T) {
	ConfirmMakeMap := func(l int, m Map[int, bool]) {
		if r := len(m); r != l {
			t.Errorf("len(%v) should be %v not %v", m, l, r)
		}
	}

	ConfirmMakeMap(0, nil)
	ConfirmMakeMap(0, make(Map[int, bool]))
	ConfirmMakeMap(0, Map[int, bool]{})
	ConfirmMakeMap(1, Map[int, bool]{ 0: true })
	ConfirmMakeMap(2, Map[int, bool]{ 0: false, 1: true })
	ConfirmMakeMap(3, Map[int, bool]{ 0: true, 1: false, 2: true })
	ConfirmMakeMap(3, Map[int, bool]{ 0: false, 1: true, 2: false })
}

func TestMaps_Len(t *testing.T) {
	ConfirmLen := func(m Map[int, bool], l int) {
		if r := m.Len(); r != l {
			t.Errorf("%v.Len() should be %v not %v", m, l, r)
		}
	}
	ConfirmLen(nil, 0)
	ConfirmLen(Map[int, bool]{}, 0)
	ConfirmLen(Map[int, bool]{ 0: true }, 1)
	ConfirmLen(Map[int, bool]{ 0: true, 1: false }, 2)
	ConfirmLen(Map[int, bool]{ 0: true, 1: false, 2: true }, 3)	
}

func TestMaps_Get(t *testing.T) {
	ConfirmGet := func(m Map[int, bool], k int, v bool) {
		switch r, ok := m.Get(k); {
		case !ok:
			t.Errorf("%v.Get(%v) did not find a value", m, k)
		case r != v:
			t.Errorf("%v.Get(%v) should be %v not %v", m, k, v, r)
		}
	}

	ConfirmGet(Map[int, bool]{ 0: true, 1: false, 2: true }, 0, true)
	ConfirmGet(Map[int, bool]{ 0: true, 1: false, 2: true }, 1, false)	
	ConfirmGet(Map[int, bool]{ 0: true, 1: false, 2: true }, 2, true)	
}

func TestMaps_Set(t *testing.T) {
	ConfirmSet := func(m Map[int, bool], k int, v bool) {
		m.Set(k, v)
		switch r, ok := m.Get(k); {
		case !ok:
			t.Errorf("%v.Set(%v) did not set a value", m, k)
		case r != v:
			t.Errorf("%v.Set(%v) should store %v not %v", m, k, v, r)
		}
	}

	ConfirmSet(Map[int, bool]{}, 0, true)
	ConfirmSet(Map[int, bool]{}, 0, false)
	ConfirmSet(Map[int, bool]{}, 1, true)
	ConfirmSet(Map[int, bool]{}, 1, false)
	ConfirmSet(Map[int, bool]{ 0: true }, 0, false)
	ConfirmSet(Map[int, bool]{ 0: false }, 0, true)
	ConfirmSet(Map[int, bool]{ 1: true }, 1, false)
	ConfirmSet(Map[int, bool]{ 1: false }, 1, true)
}

func TestMaps_Keys(t *testing.T){
	ConfirmKeys := func(m Map[int, bool], k ...int) {
		switch r := m.Keys(); {
		case len(m) != len(r):
			t.Errorf("%v.Keys() only returned %v keys of %v", m, len(r), len(m))
		case len(r) != len(k):
			t.Errorf("wrong number of keys expected: %v rather than %v", len(k), len(m))
		default:
			sort.Ints(k)
			sort.Ints(r)
			for i, v := range k {
				if r[i] != v {
					t.Errorf("%v.Keys()[%v] should be %v not %v", m, i, v, r[i])
				}
			}
		}
	}

	ConfirmKeys(nil)
	ConfirmKeys(make(Map[int, bool]))
	ConfirmKeys(Map[int, bool]{})
	ConfirmKeys(Map[int, bool]{ 0: true }, 0)
	ConfirmKeys(Map[int, bool]{ 0: true, 1: true }, 0, 1)
	ConfirmKeys(Map[int, bool]{ 0: true, 1: true, 2: true }, 0, 1, 2)
	ConfirmKeys(Map[int, bool]{ 0: false, 1: true, 2: true }, 0, 1, 2)
	ConfirmKeys(Map[int, bool]{ 0: true, 1: false, 2: true }, 0, 1, 2)
	ConfirmKeys(Map[int, bool]{ 0: true, 1: true, 2: false }, 0, 1, 2)
}

func TestMaps_KeysMatch(t *testing.T) {
	ConfirmKeysMatch := func(m1, m2 Map[int, bool]) {
		switch {
		case !m1.KeysMatch(m2):
			t.Errorf("%v.MatchKeys(%v) should be true", m1, m2)
		case !m2.KeysMatch(m1):
			t.Errorf("%v.MatchKeys(%v) should be true", m2, m1)
		}
	}

	ConfirmKeysMatch(nil, nil)
	ConfirmKeysMatch(Map[int, bool]{ 0: true }, Map[int, bool]{ 0: true })
	ConfirmKeysMatch(Map[int, bool]{ 0: true }, Map[int, bool]{ 0: false })
	ConfirmKeysMatch(Map[int, bool]{ 0: true, 1: true }, Map[int, bool]{ 0: true, 1: false })
	ConfirmKeysMatch(Map[int, bool]{ 0: true, 1: true, 2: false }, Map[int, bool]{ 0: true, 1: false, 2: false })

	RefuteKeysMatch := func(m1, m2 Map[int, bool]) {
		switch {
		case m1.KeysMatch(m2):
			t.Errorf("%v.MatchKeys(%v) should be false", m1, m2)
		case m2.KeysMatch(m1):
			t.Errorf("%v.MatchKeys(%v) should be false", m2, m1)
		}
	}

	RefuteKeysMatch(nil, Map[int, bool]{ 0: true })
	RefuteKeysMatch(Map[int, bool]{ 0: true }, Map[int, bool]{ 1: true })
	RefuteKeysMatch(Map[int, bool]{ 0: true, 1: true }, Map[int, bool]{ 1: true, 2: true })
}

func TestMaps_Copy(t *testing.T) {
	ConfirmCopy := func(m Map[int, bool]) {
		switch r := m.Copy(); {
		case len(m) != len(r):
			t.Errorf("%v.Copy() length should be %v not %v", m, len(m), len(r))
		case !r.KeysMatch(m):
			t.Errorf("%v.Copy() has incorrect keys %v", m, r.Keys())
		default:
			for k, v := range r {
				switch vo, ok := m[k]; {
				case !ok:
					t.Errorf("%v.Copy()[%v] should exist in copy", m, k)
				case vo != v:
					t.Errorf("%v.Copy()[%v] should be %v not %v", m, k, v, vo)
				}
			}
		}
	}

	ConfirmCopy(nil)
	ConfirmCopy(Map[int, bool]{})
	ConfirmCopy(Map[int, bool]{ 0: true })
	ConfirmCopy(Map[int, bool]{ 0: true, 1: false })
	ConfirmCopy(Map[int, bool]{ 0: true, 1: false, 2: true })
	ConfirmCopy(Map[int, bool]{ 0: true, 1: false, 2: true, 3: true })
}

func TestMaps_Merge(t *testing.T) {
	ConfirmMerge := func(r, m Map[int, bool], o ...Map[int, bool]) {
		switch x := m.Merge(o...); {
		case len(r) != len(x):
			t.Errorf("%v.Merge(%v) is the wrong size", m, o)
		case !x.KeysMatch(r):
			t.Errorf("%v.Merge(%v) has incorrect keys", m, o)
		default:
			for k, v := range r {
				switch vo, ok := x[k]; {
				case !ok:
					t.Errorf("%v.Merge(%v)[%v] should exist in copy", m, o, k)
				case vo != v:
					t.Errorf("%v.Merge(%v)[%v] should be %v not %v", m, o, k, v, vo)
				}
			}
		}
	}

	ConfirmMerge(nil, nil, nil)
	ConfirmMerge(Map[int, bool]{}, nil, Map[int, bool]{})
	ConfirmMerge(Map[int, bool]{}, Map[int, bool]{}, nil)
	ConfirmMerge(Map[int, bool]{ 0: true }, nil, Map[int, bool]{ 0: true })
	ConfirmMerge(Map[int, bool]{ 0: true }, Map[int, bool]{ 0: true }, nil)
	ConfirmMerge(Map[int, bool]{ 0: true, 1: false }, nil, Map[int, bool]{ 0: true, 1: false })
	ConfirmMerge(Map[int, bool]{ 0: true, 1: false }, Map[int, bool]{ 0: true, 1: false }, nil)

	ConfirmMerge(Map[int, bool]{ 0: true, 1: false }, nil, Map[int, bool]{ 0: true }, Map[int, bool]{ 1: false })
	ConfirmMerge(Map[int, bool]{ 0: true, 1: false }, Map[int, bool]{ 0: true }, nil, Map[int, bool]{ 1: false })
	ConfirmMerge(Map[int, bool]{ 0: true, 1: false }, Map[int, bool]{ 0: true }, Map[int, bool]{ 1: false }, nil)

	ConfirmMerge(Map[int, bool]{ 0: true, 1: false, 2: true }, nil, Map[int, bool]{}, Map[int, bool]{ 0: true, 1: false, 2: true })
	ConfirmMerge(Map[int, bool]{ 0: true, 1: false, 2: true }, nil, Map[int, bool]{ 0: true }, Map[int, bool]{ 1: false, 2: true })
	ConfirmMerge(Map[int, bool]{ 0: true, 1: false, 2: true  }, Map[int, bool]{ 0: true }, Map[int, bool]{ 1: false }, Map[int, bool]{ 2: true })
	ConfirmMerge(Map[int, bool]{ 0: true, 1: false, 2: true  }, Map[int, bool]{ 0: true, 1: false }, Map[int, bool]{ 2: true }, nil)
	ConfirmMerge(Map[int, bool]{ 0: true, 1: false, 2: true  }, Map[int, bool]{ 0: true, 1: false, 2: true }, Map[int, bool]{}, nil)

	ConfirmMerge(Map[int, bool]{ 0: true }, Map[int, bool]{ 0: true }, Map[int, bool]{ 0: true }, Map[int, bool]{ 0: true })
	ConfirmMerge(Map[int, bool]{ 0: true }, Map[int, bool]{ 0: false }, Map[int, bool]{ 0: true }, Map[int, bool]{ 0: true })
	ConfirmMerge(Map[int, bool]{ 0: true }, Map[int, bool]{ 0: true }, Map[int, bool]{ 0: false }, Map[int, bool]{ 0: true })
	ConfirmMerge(Map[int, bool]{ 0: false }, Map[int, bool]{ 0: true }, Map[int, bool]{ 0: true }, Map[int, bool]{ 0: false })
}

func TestMaps_Select(t *testing.T) {
	ConfirmSelect := func(r, m Map[int, bool], k ...int) {
		x := m.Select(k...)
		if len(r) != len(x) {
			t.Errorf("%v.Select(%v) should contain %v elements not %v", m, k, len(r), len(x))
		}
		for _, v := range k {
			switch vo, ok := x[v]; {
			case !ok:
				t.Errorf("%v.Select(%v) should contain key %v", m, k, v)
			case vo != r[v]:
				t.Errorf("%v.Select(%v)[%v] should be %v not %v", m, k, v, r[v], vo)
			}
		}
	}

	ConfirmSelect(Map[int, bool]{}, Map[int, bool]{ 0: true, 1: false, 2: true })
	ConfirmSelect(Map[int, bool]{ 0: true }, Map[int, bool]{ 0: true, 1: false, 2: true }, 0)
	ConfirmSelect(Map[int, bool]{ 1: false }, Map[int, bool]{ 0: true, 1: false, 2: true }, 1)
	ConfirmSelect(Map[int, bool]{ 2: true }, Map[int, bool]{ 0: true, 1: false, 2: true }, 2)

	ConfirmSelect(Map[int, bool]{ 0: true }, Map[int, bool]{ 0: true, 1: false, 2: true }, 0, 0)
	ConfirmSelect(Map[int, bool]{ 1: false }, Map[int, bool]{ 0: true, 1: false, 2: true }, 1, 1)
	ConfirmSelect(Map[int, bool]{ 2: true }, Map[int, bool]{ 0: true, 1: false, 2: true }, 2, 2)


	ConfirmSelect(Map[int, bool]{ 0: true, 1: false }, Map[int, bool]{ 0: true, 1: false, 2: true }, 0, 1)
	ConfirmSelect(Map[int, bool]{ 1: false, 2: true }, Map[int, bool]{ 0: true, 1: false, 2: true }, 1, 2)
	ConfirmSelect(Map[int, bool]{ 0: true, 2: true }, Map[int, bool]{ 0: true, 1: false, 2: true }, 0, 2)


	ConfirmSelect(Map[int, bool]{ 1: false }, Map[int, bool]{ 0: true, 1: false, 2: true }, 1)


	ConfirmSelect(Map[int, bool]{ 2: true }, Map[int, bool]{ 0: true, 1: false, 2: true }, 2)

}