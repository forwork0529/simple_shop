package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strconv"
)

var appLogger *zap.SugaredLogger

// New инициализирует логер в main прямым вызовом а не скрытым через init чтобы передать в логгер уровень логирования из env переменных
func New(inputLoggerLevel string) {

	loggerLevel := getUsableLoggerLevel(inputLoggerLevel)

	writerSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, loggerLevel)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.FatalLevel))
	appLogger = logger.Sugar()
	defer appLogger.Sync()
}

func getUsableLoggerLevel(loggerLevel string) zapcore.Level {

	if len(loggerLevel) < 1 {
		log.Println("INFO: no LOGGER_LEVEL given, chosen 0")
		return zapcore.InfoLevel
	}

	levelInt, err := strconv.Atoi(loggerLevel)
	if err != nil {
		log.Println("INFO: strconv.Atoi(); cant parse given LOGGER_LEVEL, chosen 0")
		return zapcore.InfoLevel
	}
	if levelInt < -1 || levelInt > 5 {
		log.Println("INFO: invalid given LOGGER_LEVEL, chosen 0")
		return zapcore.InfoLevel
	}
	log.Printf("INFO: chosed LOGGER_LEVEL = %v\n", levelInt)
	return zapcore.Level(levelInt)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func Debug(args ...interface{}) {
	appLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	appLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	appLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	appLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	appLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	appLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	appLogger.Error(args...)
}

// FastErr is wrapper for ErrorF, returns Errorf("funcName+(): %v", err)
func FastErr(funcName string, err error) {
	Errorf(funcName+"(): %v", err)
}

func Errorf(template string, args ...interface{}) {
	appLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	appLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	appLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	appLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	appLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	appLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	appLogger.Fatalf(template, args...)
}
