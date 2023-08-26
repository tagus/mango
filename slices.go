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
