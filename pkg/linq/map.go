package linq

func Map[TFrom interface{}, TTo interface{}](from []TFrom, handler func(TFrom) TTo) []TTo {
	result := make([]TTo, len(from))
	for i, v := range from {
		result[i] = handler(v)
	}
	return result
}
