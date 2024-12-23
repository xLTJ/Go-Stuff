package strain

func Keep[T any](collection []T, predicate func(T) bool) []T {
	var filteredCollection []T
	for _, element := range collection {
		if predicate(element) {
			filteredCollection = append(filteredCollection, element)
		}
	}

	return filteredCollection
}

func Discard[T any](collection []T, predicate func(T) bool) []T {
	var filteredCollection []T
	for _, element := range collection {
		if !predicate(element) {
			filteredCollection = append(filteredCollection, element)
		}
	}

	return filteredCollection
}
