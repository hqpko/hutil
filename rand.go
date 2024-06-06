package hutils

import (
	"math/rand/v2"
)

const strBase = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var strMaxLen = len(strBase)

func RandString(n int) string {
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = strBase[rand.IntN(strMaxLen)]
	}
	return string(buf)
}

func RandSlice[T any](slice []T, n int) []T {
	l := len(slice)
	if n > l {
		n = l
	}
	r := make([]T, n)
	for i := 0; i < n; i++ {
		j := rand.IntN(l)
		r[i] = slice[j]
		slice[j], slice[l-1] = slice[l-1], slice[j]
		l--
	}
	return r
}

func RandMapK[K comparable, V any](m map[K]V, n int) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return RandSlice(keys, n)
}

func RandMapV[K comparable, V any](m map[K]V, n int) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return RandSlice(values, n)
}
