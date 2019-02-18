package log

// KV are key value pairs to be logged
type KV struct {
	K string
	V interface{}
}

// InfoWriter for logs
type InfoWriter interface {
	Info(msg string, kv ...KV)
}

// Logger logs messages and errors
type Logger interface {
	InfoWriter
	Error(err error, msg string, kv ...KV)
	SetLevel(int)
	V(int) InfoWriter
}
