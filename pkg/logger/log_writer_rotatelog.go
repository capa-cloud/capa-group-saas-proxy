package logger

import (
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"io"
	"time"
)

// RotateFileLogWriter Support the log writing to file, and log file can be rotated by time
type RotateFileLogWriter struct {
	FilePathName            string
	RotationFileNamePattern string
	MaxAge                  time.Duration
	RotationTime            time.Duration
}

func (writer RotateFileLogWriter) GetLogWriter() io.Writer {
	// Set rotate log
	logWriter, _ := rotateLogs.New(
		// Split file name
		writer.FilePathName+"."+writer.RotationFileNamePattern+".log",
		// Set soft link to the latest log file
		rotateLogs.WithLinkName(writer.FilePathName),
		// Maximum save time of log file
		rotateLogs.WithMaxAge(writer.MaxAge),
		// Cutting interval of log file
		rotateLogs.WithRotationTime(writer.RotationTime),
	)
	return logWriter
}
