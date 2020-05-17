package log

import (
	"engine/conf"
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"time"
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

func (logger *Logger) Author(author string) *Entry {
	return logger.WithField("author", author)
}

// If you want multiple fields, use `WithFields`.
func (logger *Logger) WithField(key string, value interface{}) *Entry {
	return &Entry{Entry: logger.Logger.WithField(key, value)}
}

/*
logrus_amqp：Logrus hook for Activemq。
logrus-logstash-hook:Logstash hook for logrus。
mgorus:Mongodb Hooks for Logrus。
logrus_influxdb:InfluxDB Hook for Logrus。
logrus-redis-hook:Hook for Logrus which enables logging to RELK stack (Redis, Elasticsearch, Logstash and Kibana)。
 */
func InitLog() error {
	baseLogPath := path.Join(conf.EngineConf.Logger.Path, conf.EngineConf.Logger.FileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y-%m-%d",
		//baseLogPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(baseLogPath),                                                    // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(time.Second*time.Duration(conf.EngineConf.Logger.MaxAge)),             // 文件最大保存时间
		rotatelogs.WithRotationTime(time.Second*time.Duration(conf.EngineConf.Logger.RotationTime)), // 日志切割时间间隔
	)
	if err != nil {
		return fmt.Errorf("make rotatelogs writer is fail: %v", err)
	}
	writerMap := make(lfshook.WriterMap)
	writerMap[PanicLevel] = writer
	writerMap[FatalLevel] = writer
	writerMap[ErrorLevel] = writer
	writerMap[WarnLevel] = writer
	writerMap[InfoLevel] = writer
	writerMap[DebugLevel] = writer
	writerMap[TraceLevel] = writer
	hook := lfshook.NewHook(writerMap, JSONFormatter)

	level := DebugLevel
	switch conf.EngineConf.Logger.Level {
	case "panic":
		level = PanicLevel
	case "fatal":
		level = FatalLevel
	case "error":
		level = ErrorLevel
	case "warn":
		level = WarnLevel
	case "info":
		level = InfoLevel
	case "debug":
		level = DebugLevel
	case "trace":
		level = TraceLevel

	}

	log := logrus.New()
	log.AddHook(hook)
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(&emptyWrite{})
	log.SetLevel(level)
	Log = &Logger{Logger: log}

	return nil
}
