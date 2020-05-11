package log

var Log *Logger

func Error(args ...interface{}) {
	Log.Error(args...)
}

func Warn(args ...interface{}) {
	Log.Warn(args...)
}

func Info(args ...interface{}) {
	Log.Info(args...)
}

func Debug(args ...interface{}) {
	Log.Debug(args...)
}

func Fatal(args ...interface{}) {
	Log.Fatal(args...)
}

func Errorf(format string, v ...interface{}) {
	Log.Errorf(format, v...)
}

func Warnf(format string, v ...interface{}) {
	Log.Warnf(format, v...)
}

func Infof(format string, v ...interface{}) {
	Log.Infof(format, v...)
}

func Debugf(format string, v ...interface{}) {
	Log.Debugf(format, v...)
}

func Fatalf(format string, args ...interface{}) {
	Log.Fatalf(format, args...)
}
