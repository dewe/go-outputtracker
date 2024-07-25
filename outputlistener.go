package outputtracker

// Used internally in the Nullable infrastructure

type OutputListener[T any] struct {
	listeners []*OutputTracker[T]
}

func CreateListener[T any]() *OutputListener[T] {
	return &OutputListener[T]{listeners: make([]*OutputTracker[T], 0)}
}

func (ol *OutputListener[T]) CreateTracker() *OutputTracker[T] {
	tracker := &OutputTracker[T]{outputListener: ol}
	ol.listeners = append(ol.listeners, tracker)
	return tracker
}

func (ol *OutputListener[T]) Track(data T) {
	for _, tracker := range ol.listeners {
		tracker.add(data)
	}	
}	

func (ol *OutputListener[T]) remove(tracker *OutputTracker[T]) {
	for i, t := range ol.listeners {
		if t == tracker {
			ol.listeners = append(ol.listeners[:i], ol.listeners[i+1:]...)
			break
		}
	}
}
