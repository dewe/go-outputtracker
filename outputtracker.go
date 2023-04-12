package outputtracker

// Used in tests of the Nullable infrastructure

type OutputTracker[T any] struct {
	output        []T
	outputListener *OutputListener[T]
}

func (ot *OutputTracker[T]) add(data T) {
	ot.output = append(ot.output, data)
}

func (ot *OutputTracker[T]) Data() []T {
	dataCopy := make([]T, len(ot.output))
	copy(dataCopy, ot.output)
	return dataCopy
}

func (ot *OutputTracker[T]) Clear() []T {
	data := ot.Data()
	ot.output = ot.output[:0]
	return data
}

func (ot *OutputTracker[T]) Stop() {
	ot.outputListener.remove(ot)
}
