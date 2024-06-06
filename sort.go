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

func SliceToMap[K comparable, V any](a []V, key func(V) K) map[K]V {
	m := make(map[K]V, len(a))
	for _, v := range a {
		m[key(v)] = v
	}
	return m
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
