package toolbox

import "context"

// Logger interface for logging with level
type Logger interface {
	Warn(ctx context.Context, keyvals ...interface{})
	Info(ctx context.Context, keyvals ...interface{})
}
