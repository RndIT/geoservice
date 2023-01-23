package logger

import (
	"sync"
)

// какого типа будет логгер?
// enum
type LogType int

const (
	EPrintLog LogType = iota + 1
	EStdLog
)

// какой интерфейс будет предоставлять логер?
type Logger interface {
	Debug(v ...any)
}


// Создает внутри себя нужный тип логгера
// с применением принципа единоначалия, т.е. синглетона
// Работа с логгером - через интерфейсные методы.
// Сам объект логера - скрыт внутри реализации

// где будет храниться экземпляр?
// описание типа структуры
type logger struct {
	Log Logger
}

// выделение переменной
var loggerInstance *logger
// типовой вариант обработчика
func Debug(v ...any) {

	loggerInstance.Log.Debug(v...)
}

// для синглтона - логгер инициализируется только один раз!
var once sync.Once

func NewLogger(logType LogType) {
	// сработает только один раз в этом модуле!
	once.Do(func() {
		loggerObject := &logger{}
		switch logType {
		case EPrintLog:
			{
				loggerObject.Log = BuildPrintLogger()
			}
		case EStdLog:
			{
				loggerObject.Log = BuildSimpleLogger()
			}
		}
		loggerInstance = loggerObject
	})
}


