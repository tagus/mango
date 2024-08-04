package mango

func Map[F any, G any](list []F, mapper func(F) G) []G {
	mapped := make([]G, len(list))
	for i := range list {
		mapped[i] = mapper(list[i])
	}
	return mapped
}

func Filter[F any](list []F, predicate func(F) bool) []F {
	filtered := make([]F, 0, len(list))
	for _, item := range list {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func Find[F any](list []F, predicate func(F) bool) (F, bool) {
	var empty F
	for _, item := range list {
		if predicate(item) {
			return item, true
		}
	}
	return empty, false
}

/******************************************************************************/

type Comparator[F any] func(F, F) bool

func SliceEqual[F any](first, second []F, comparator Comparator[F]) bool {
	if len(first) != len(second) {
		return false
	}
	for i := range first {
		if !comparator(first[i], second[i]) {
			return false
		}
	}
	return true
}

func StringSliceEqual(first, second []string) bool {
	return SliceEqual(first, second, func(a, b string) bool { return a == b })
}
