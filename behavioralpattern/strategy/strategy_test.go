package strategy

import (
	"testing"
)

func TestLogStrategy(t *testing.T) {
	l := LogContext{
		logStrategy: LogFunc(DBLog),
	}
	defer func() {
		msg := recover()
		l.logStrategy = LogFunc(FileLog)
		l.logStrategy.Log(msg.(string))
	}()
	l.logStrategy.Log("aaaaaa")
}
