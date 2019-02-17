package log

// KV are key value pairs to be logged
type KV map[string]interface{}

// InfoWriter for logs
type InfoWriter interface {
	Info(msg string, kv ...KV)
}

// Logger logs messages and errors
type Logger interface {
	InfoWriter
	Error(err error, msg string, kv ...KV)
	V(int) InfoWriter
}
