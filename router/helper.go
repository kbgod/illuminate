package router

func firstNotNil[T any](fields ...*T) *T {
	for _, field := range fields {
		if field != nil {
			return field
		}
	}

	return nil
}
