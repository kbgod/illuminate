package plugin

func Void(returns ...any) error {
	lastValue := returns[len(returns)-1]
	if err, ok := lastValue.(error); ok {
		return err
	}
	return nil
}
