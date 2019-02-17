package log

type KV map[string]interface{}

// Writer for log messages
type Writer interface {
	Info(msg string, kv ...KV)
	Error(err error, msg string, kv ...KV)
}

// Log messages and errors
type Logger interface {
	Writer
	V(int) Writer
}
