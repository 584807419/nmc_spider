package log_manage

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var FSLogger = InitFSLogger()

func InitFSLogger() *zap.SugaredLogger {
	writeSyncer := getFileMoreLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout)), zapcore.DebugLevel)
	//file_logger := zap.New(core)
	file_logger := zap.New(core, zap.AddCaller()) //添加将调用函数信息记录到日志中的功能。
	return file_logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	//return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	//更改时间编码并添加调用者详细信息
	encoderConfig := zap.NewProductionEncoderConfig()
	// encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getFileMoreLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./doc/nmc_spider.log", //日志文件的位置
		MaxSize:    10,                     //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 5,                      //保留旧文件的最大个数
		MaxAge:     30,                     //保留旧文件的最大天数
		Compress:   false,                  //是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}
