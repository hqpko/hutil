package hutils

import (
	"cmp"
	"math/rand/v2"
	"sort"
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
	n = Clamp(n, 0, l)
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

func RandSliceWeightFunc[T any](slice []T, weightFunc func(T) int) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}

	prefixSums := make([]int, len(slice))
	totalWeight := 0
	for i, item := range slice {
		totalWeight += weightFunc(item)
		prefixSums[i] = totalWeight
	}
	rnd := rand.IntN(totalWeight)
	idx := sort.Search(len(prefixSums), func(j int) bool {
		return prefixSums[j] > rnd
	})
	return slice[idx]
}

// RandSliceWeightFuncN 根据权重随机选择 n 个元素，权重由 weightFunc 提供，可以重复选择同一个元素
func RandSliceWeightFuncN[T any](slice []T, n int, weightFunc func(T) int) []T {
	if n <= 0 || len(slice) == 0 {
		return nil
	}

	prefixSums := make([]int, len(slice))
	totalWeight := 0
	for i, item := range slice {
		totalWeight += weightFunc(item)
		prefixSums[i] = totalWeight
	}
	selectedItems := make([]T, 0, n)
	for i := 0; i < n; i++ {
		rnd := rand.IntN(totalWeight)
		idx := sort.Search(len(prefixSums), func(j int) bool {
			return prefixSums[j] > rnd
		})
		selectedItems = append(selectedItems, slice[idx])
	}
	return selectedItems
}

func Clamp[T cmp.Ordered](v, min, max T) T {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
