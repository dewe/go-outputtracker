# outputtracker

A Go output tracker for creating Nullable infrastructure.

Inspired by https://www.jamesshore.com/v2/projects/nullables/testing-without-mocks

## Usage

```go
package log

import "github.com/dewe/go-outputtracker"

type Log struct {
	listener *outputtracker.OutputListener[LogData]
}

func CreateLog() *Log {
	return &Log{listener: outputtracker.CreateListener[LogData]()}
}

// TrackOutput creates the output tracker
func (l *Log) TrackOutput() *outputtracker.OutputTracker[LogData] {
	return l.listener.CreateTracker()
}

func (l *Log) Info(data LogData) {
	// ...

	// Emit the event
	l.listener.Track(data)
}
```
