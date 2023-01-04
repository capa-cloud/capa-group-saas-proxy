package logger

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)

const (
	TimeFormatPattern = "2006-01-02 15:04:05"
)

// InitLog init log settings
func InitLog() {
	logConfig := config.NewLogConfig()
	// Log file name
	fileName := path.Join(logConfig.Get(config.LogFilePathKey), logConfig.Get(config.LogFileNameKey))
	rotationFileNamePattern := logConfig.Get(config.RotationPatternKey)

	// Set log level
	level, err := logrus.ParseLevel(logConfig.Get(config.LogLevelKey))
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   TimeFormatPattern,
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyFunc:  "caller",
		},
		PrettyPrint: false,
	})

	rotateFileLogWriter := &RotateFileLogWriter{
		FilePathName:            fileName,
		RotationFileNamePattern: rotationFileNamePattern,
		MaxAge:                  7 * 24 * time.Hour,
		RotationTime:            1 * 24 * time.Hour,
	}

	// Get rotate log writer
	logWriter := rotateFileLogWriter.GetWriter()

	// hook
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: TimeFormatPattern,
	})
	logrus.AddHook(lfHook)
}
