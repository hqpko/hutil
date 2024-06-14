package hutils

import "sort"

func TopSliceN[T any](a []T, n int, less func(T, T) bool) []T {
	if n > len(a) {
		n = len(a)
	}
	b := append([]T(nil), a...)
	sort.Slice(b, func(i, j int) bool { return less(b[i], b[j]) })
	return b[:n]
}

func TopMapN[K comparable, V any](m map[K]V, n int, less func(V, V) bool) []V {
	var a []V
	for _, v := range m {
		a = append(a, v)
	}
	sort.Slice(a, func(i, j int) bool { return less(a[i], a[j]) })
	if n > len(a) {
		n = len(a)
	}
	return a[:n]
}

func SliceToMap[K comparable, V, T any](a []V, transFunc func(V) (K, T, bool)) map[K]T {
	m := make(map[K]T, len(a))
	for _, v := range a {
		if k, t, ok := transFunc(v); ok {
			m[k] = t
		}
	}
	return m
}

func SliceToSlice[T, V any](a []T, transFunc func(T) (V, bool)) []V {
	b := make([]V, 0, len(a))
	for _, v := range a {
		if r, ok := transFunc(v); ok {
			b = append(b, r)
		}
	}
	return b
}

func MapToSlice[K comparable, V, T any](m map[K]V, transFunc func(K, V) (T, bool)) []T {
	a := make([]T, 0, len(m))
	for k, v := range m {
		if t, ok := transFunc(k, v); ok {
			a = append(a, t)
		}
	}
	return a
}

func MapToSliceK[K comparable, V any](m map[K]V) []K {
	a := make([]K, 0, len(m))
	for k := range m {
		a = append(a, k)
	}
	return a
}

func MapToSliceV[K comparable, V any](m map[K]V) []V {
	a := make([]V, 0, len(m))
	for _, v := range m {
		a = append(a, v)
	}
	return a
}
