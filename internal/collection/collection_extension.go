package collection

func Map[T any, U any](input []T, fn func(T) U) []U {
	output := make([]U, len(input))
	for i, v := range input {
		output[i] = fn(v)
	}
	return output
}

func ToSet[T comparable](input []T) map[T]struct{} {
	set := make(map[T]struct{}, len(input))
	for _, v := range input {
		set[v] = struct{}{}
	}
	return set
}
