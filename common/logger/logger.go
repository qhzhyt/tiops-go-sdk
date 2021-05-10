package logger

import (
	"context"
	"fmt"
	"path"
	"runtime"
	"strings"
	tiopsApiClient "tiops/common/api-client"
	"tiops/common/models"
	"tiops/common/utils"
)

const (
	LoglevelDebug    = "DEBUG"
	LoglevelInfo     = "INFO"
	LoglevelWarning  = "WARNING"
	LoglevelError    = "ERROR"
	LoglevelCritical = "CRITICAL"
)

var (
	projectRoot = ""
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	projectRoot = strings.Split(file, "common")[0]
}

type Logger struct {
	ID        int64
	Type      string
	Level     models.LogLevel
	logChan   chan *models.Log
	currentSn int32
	remote    bool
	source    string
	apiClient *tiopsApiClient.APIClient
}

type LogConfig struct {
}

func GetCallerPosition(skip int) (string, int) {
	_, file, line, _ := runtime.Caller(skip)
	return strings.ReplaceAll(file, projectRoot, ""), line
}

func NewLogger(_type string, level models.LogLevel) *Logger {
	logger := &Logger{ID: utils.SnowflakeID(), Type: _type, Level: level, logChan: make(chan *models.Log, 1000)}
	logger.start()
	return logger
}

func NewRemoteLogger(_type, id, source string, level models.LogLevel, client *tiopsApiClient.APIClient) *Logger {
	logger := &Logger{
		ID:        utils.DecimalToInt64(id),
		source:    source,
		Type:      _type,
		Level:     level,
		logChan:   make(chan *models.Log, 1000),
		remote:    true,
		apiClient: client,
	}
	logger.start()
	return logger
}

func (l *Logger) start() {
	go func() {
		sn := int32(0)
		for log := range l.logChan {
			log.Sn = sn
			sn++
			l.processLog(log)
		}
	}()
}

func (l *Logger) processLog(_log *models.Log) {
	if l.remote {
		_, _ = l.apiClient.PostLog(context.TODO(), _log)
	}
	//fmt.Println(_log)
}

func (l *Logger) newLog(msg interface{}, level models.LogLevel) *models.Log {
	pc, file, line, _ := runtime.Caller(2)
	module, name := path.Split(runtime.FuncForPC(pc).Name())

	module = path.Join(module, strings.Split(name, ".")[0])

	file = strings.ReplaceAll(file, projectRoot, "")
	return &models.Log{
		Level:    level,
		Time:     utils.CurrentTimeStampMS(),
		Type:     l.Type,
		LogID:    l.ID,
		Source:   l.source,
		Content:  fmt.Sprint(msg),
		Module:   module,
		Position: fmt.Sprintf("%s:%d", file, line),
	}
}

func (l *Logger) NewLog(msg interface{}, level models.LogLevel) *models.Log {
	return l.newLog(msg, level)
}

func (l *Logger) NewErrorLog(msg interface{}) *models.Log {
	return l.newLog(msg, models.LogLevel_ERROR)
}

func (l *Logger) NewInfoLog(msg interface{}) *models.Log {
	return l.newLog(msg, models.LogLevel_INFO)
}

func (l *Logger) Emit(log *models.Log) {
	l.logChan <- log
}

func (l *Logger) Debug(msg interface{}) {
	//logger := zap.New(core, zap.AddCaller())
	if l.Level <= models.LogLevel_DEBUG {
		l.logChan <- l.newLog(msg, models.LogLevel_DEBUG)
	}
}

func (l *Logger) Info(msg interface{}) {
	//logger := zap.New(core, zap.AddCaller())
	if l.Level <= models.LogLevel_INFO {
		l.logChan <- l.newLog(msg, models.LogLevel_INFO)
	}
}

func (l *Logger) Warning(msg interface{}) {
	//logger := zap.New(core, zap.AddCaller())
	if l.Level <= models.LogLevel_WARNING {
		l.logChan <- l.newLog(msg, models.LogLevel_WARNING)
	}
}

func (l *Logger) Error(msg interface{}) {
	//logger := zap.New(core, zap.AddCaller())
	if l.Level <= models.LogLevel_ERROR {
		l.logChan <- l.newLog(msg, models.LogLevel_ERROR)
	}
}

func (l *Logger) Critical(msg interface{}) {
	//logger := zap.New(core, zap.AddCaller())
	if l.Level <= models.LogLevel_CRITICAL {
		l.logChan <- l.newLog(msg, models.LogLevel_CRITICAL)
	}
}

func (l *Logger) Println(msg interface{}) {
	_log := l.newLog(msg, models.LogLevel_DEBUG)
	fmt.Printf("%s", _log.Content)
	//logger := zap.New(core, zap.AddCaller())
}

func StringToLogLevel(level string) models.LogLevel {
	switch level {
	case LoglevelCritical:
		return models.LogLevel_CRITICAL
	case LoglevelError:
		return models.LogLevel_ERROR
	case LoglevelWarning:
		return models.LogLevel_WARNING
	case LoglevelInfo:
		return models.LogLevel_INFO
	case LoglevelDebug:
		return models.LogLevel_DEBUG
	default:
		return models.LogLevel_INFO
	}
}
