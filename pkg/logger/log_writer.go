package logger

import "io"

// LogWriter Log writer.
type LogWriter interface {
	// GetLogWriter Get the io writer.
	GetLogWriter() io.Writer
}

func name() {

}
