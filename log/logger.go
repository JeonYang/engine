package log

import (
	"github.com/sirupsen/logrus"
)

const (
	PanicLevel logrus.Level = iota

	FatalLevel

	ErrorLevel

	WarnLevel

	InfoLevel

	DebugLevel

	TraceLevel
)

var JSONFormatter = &logrus.JSONFormatter{}

type Logger struct {
	*logrus.Logger
}

type Entry struct {
	*logrus.Entry
}

type emptyWrite struct {
}

func (write *emptyWrite) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func NewLogger(hook logrus.Hook, level logrus.Level) *Logger {
	log = logrus.New()
	log.AddHook(hook)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(&emptyWrite{})
	log.SetLevel(level)
	return &Logger{Logger: log}
}

func (logger *Logger) Author(author string) *Entry {
	return logger.WithField("author", author)
}

// If you want multiple fields, use `WithFields`.
func (logger *Logger) WithField(key string, value interface{}) *Entry {
	return &Entry{Entry: logger.Logger.WithField(key, value)}
}

var log *logrus.Logger
