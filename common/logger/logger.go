// @Title  logger.go
// @Description  远程日志客户端实现
// @Create  heyitong  2022/6/23 15:55
// @Update  heyitong  2022/6/23 15:55

package logger

import (
	"context"
	"fmt"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"
	client "tiops/common/api-client"
	tiopsApiClient "tiops/common/api-client"
	tiopsConfigs "tiops/common/config"
	"tiops/common/models"
	"tiops/common/utils"
)

var (
	projectRoot = ""
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	projectRoot = strings.Split(file, "common")[0]
}

// Logger 远程日志客户端
type Logger struct {
	ID        string
	Type      string
	Level     models.LogLevel
	logChan   chan *models.Log
	currentSn int32
	remote    bool
	source    string
	closed    bool
	working   *sync.WaitGroup
	apiClient *tiopsApiClient.APIClient
}

type LogConfig struct {
}

func GetCallerPosition(skip int) (string, int) {
	_, file, line, _ := runtime.Caller(skip)
	return strings.ReplaceAll(file, projectRoot, ""), line
}

func NewLogger(_type string, level models.LogLevel) *Logger {
	logger := &Logger{ID: fmt.Sprintf("%x", utils.SnowflakeID()), Type: _type, Level: level, logChan: make(chan *models.Log, 1000), working: &sync.WaitGroup{}}
	logger.start()
	return logger
}

func NewRemoteLogger(_type, id, source string, level models.LogLevel, client *tiopsApiClient.APIClient) *Logger {
	logger := &Logger{
		ID:        id,
		source:    source,
		Type:      _type,
		Level:     level,
		logChan:   make(chan *models.Log, 1000),
		remote:    true,
		apiClient: client,
		working:   &sync.WaitGroup{},
	}
	logger.start()
	return logger
}

func NewActionLogger(name string) *Logger {
	logger := &Logger{
		ID:        tiopsConfigs.LogId,
		source:    name,
		Type:      "action",
		Level:     StringToLogLevel(tiopsConfigs.LogLevel),
		logChan:   make(chan *models.Log, 1000),
		remote:    true,
		apiClient: _defaultLogger.apiClient,
		working:   &sync.WaitGroup{},
	}
	logger.start()
	return logger
}

func (l *Logger) start() {
	l.working.Add(1)
	go func() {
		sn := int32(0)
		for log := range l.logChan {
			log.Sn = sn
			sn++
			l.processLog(log)
		}
		time.Sleep(time.Second)
		l.working.Done()
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
	if !l.closed {
		l.logChan <- log
	} else {
		l.Println(log.Content)
	}
}

func (l *Logger) Debug(msg ...interface{}) {
	//logger := zap.New(core, zap.AddCaller())
	if l.Level <= models.LogLevel_DEBUG {
		l.Emit(l.newLog(fmt.Sprint(msg...), models.LogLevel_DEBUG))
	}
}

func (l *Logger) Info(msg ...interface{}) {
	//logger := zap.New(core, zap.AddCaller())
	if l.Level <= models.LogLevel_INFO {
		l.Emit(l.newLog(fmt.Sprint(msg...), models.LogLevel_INFO))
	}
}

func (l *Logger) Warning(msg ...interface{}) {
	//logger := zap.New(core, zap.AddCaller())
	if l.Level <= models.LogLevel_WARNING {
		//l.logChan <- l.newLog(fmt.Sprint(msg...), models.LogLevel_WARNING)
		l.Emit(l.newLog(fmt.Sprint(msg...), models.LogLevel_WARNING))
	}
}

func (l *Logger) Error(msg ...interface{}) {
	//logger := zap.New(core, zap.AddCaller())
	if l.Level <= models.LogLevel_ERROR {
		//l.logChan <- l.newLog(fmt.Sprint(msg...), models.LogLevel_ERROR)
		l.Emit(l.newLog(fmt.Sprint(msg...), models.LogLevel_ERROR))
	}
}

func (l *Logger) Critical(msg ...interface{}) {
	//logger := zap.New(core, zap.AddCaller())
	if l.Level <= models.LogLevel_CRITICAL {
		//l.logChan <- l.newLog(fmt.Sprint(msg...), models.LogLevel_CRITICAL)
		l.Emit(l.newLog(fmt.Sprint(msg...), models.LogLevel_CRITICAL))
	}
}

func (l *Logger) Println(msg ...interface{}) {
	_log := l.newLog(fmt.Sprint(msg...), models.LogLevel_DEBUG)
	fmt.Printf("%s", _log.Content)
	//logger := zap.New(core, zap.AddCaller())
}

func (l *Logger) Close() {
	l.closed = true
	close(l.logChan)
}

func (l *Logger) CloseAndWait() {
	l.Close()
	l.working.Wait()
}

func StringToLogLevel(level string) models.LogLevel {
	switch level {
	case tiopsConfigs.LoglevelCritical:
		return models.LogLevel_CRITICAL
	case tiopsConfigs.LoglevelError:
		return models.LogLevel_ERROR
	case tiopsConfigs.LoglevelWarning:
		return models.LogLevel_WARNING
	case tiopsConfigs.LoglevelInfo:
		return models.LogLevel_INFO
	case tiopsConfigs.LoglevelDebug:
		return models.LogLevel_DEBUG
	default:
		return models.LogLevel_INFO
	}
}

func LogLevelToString(level models.LogLevel) string {
	switch level {
	case models.LogLevel_DEBUG:
		return tiopsConfigs.LoglevelDebug
	case models.LogLevel_INFO:
		return tiopsConfigs.LoglevelInfo
	case models.LogLevel_WARNING:
		return tiopsConfigs.LoglevelWarning
	case models.LogLevel_ERROR:
		return tiopsConfigs.LoglevelError
	case models.LogLevel_CRITICAL:
		return tiopsConfigs.LoglevelCritical
	default:
		return tiopsConfigs.LoglevelInfo
	}
}

var _defaultLogger *Logger

func init() {
	apiClient := client.NewAPIClient(tiopsConfigs.ApiServerHost, tiopsConfigs.ApiServerGrpcPort)
	source := "unknown"
	if tiopsConfigs.InMainEngine() {
		source = "main-engine"
	} else {
		source = tiopsConfigs.AppName
	}
	//switch tiopsConfigs.AppType {
	//case tiopsConfigs.AppTypeWorkflowEngine:
	//	source = "workflow-engine"
	//case tiopsConfigs.AppTypeActionServer:
	//	source = fmt.Sprintf("action-server-%s", tiopsConfigs.ProjectId)
	//}
	_defaultLogger = NewRemoteLogger(
		tiopsConfigs.AppType,
		tiopsConfigs.LogId,
		source,
		StringToLogLevel(tiopsConfigs.LogLevel),
		apiClient,
	)
}

func GetDefaultLogger() *Logger {
	return _defaultLogger
}

func Debug(msg ...interface{}) {
	//logger := zap.New(core, zap.AddCaller())
	if _defaultLogger.Level <= models.LogLevel_DEBUG {
		//_defaultLogger.logChan <- _defaultLogger.newLog(fmt.Sprint(msg...), models.LogLevel_DEBUG)
		_defaultLogger.Emit(_defaultLogger.newLog(fmt.Sprint(msg...), models.LogLevel_DEBUG))
	}
}

func Info(msg ...interface{}) {
	//logger := zap.New(core, zap.AddCaller())
	if _defaultLogger.Level <= models.LogLevel_INFO {
		//_defaultLogger.logChan <- _defaultLogger.newLog(fmt.Sprint(msg...), models.LogLevel_INFO)
		_defaultLogger.Emit(_defaultLogger.newLog(fmt.Sprint(msg...), models.LogLevel_INFO))
	}
}

func Warning(msg ...interface{}) {
	//logger := zap.New(core, zap.AddCaller())
	if _defaultLogger.Level <= models.LogLevel_WARNING {
		//_defaultLogger.logChan <- _defaultLogger.newLog(fmt.Sprint(msg...), models.LogLevel_WARNING)
		_defaultLogger.Emit(_defaultLogger.newLog(fmt.Sprint(msg...), models.LogLevel_WARNING))
	}
}

func Error(msg ...interface{}) {
	//logger := zap.New(core, zap.AddCaller())
	if _defaultLogger.Level <= models.LogLevel_ERROR {
		//_defaultLogger.logChan <- _defaultLogger.newLog(fmt.Sprint(msg...), models.LogLevel_ERROR)
		_defaultLogger.Emit(_defaultLogger.newLog(fmt.Sprint(msg...), models.LogLevel_ERROR))
	}
}

func Critical(msg ...interface{}) {
	//logger := zap.New(core, zap.AddCaller())
	if _defaultLogger.Level <= models.LogLevel_CRITICAL {
		//_defaultLogger.logChan <- _defaultLogger.newLog(fmt.Sprint(msg...), models.LogLevel_CRITICAL)
		_defaultLogger.Emit(_defaultLogger.newLog(fmt.Sprint(msg...), models.LogLevel_CRITICAL))
	}
}

func Println(msg ...interface{}) {
	_log := _defaultLogger.newLog(fmt.Sprint(msg...), models.LogLevel_DEBUG)
	fmt.Printf("%s", _log.Content)
	//logger := zap.New(core, zap.AddCaller())
}
