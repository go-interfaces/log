package log

// NoLog won't output any log.
type NoLog struct{}

// Info wont output any log
func (l *NoLog) Info(msg string, kv ...KV) {}

// Error wont output any log
func (l *NoLog) Error(err error, msg string, kv ...KV) {}

// V will return its same NoLog instance
func (l *NoLog) V(int) InfoWriter { return l }

var _ Logger = (*NoLog)(nil)
