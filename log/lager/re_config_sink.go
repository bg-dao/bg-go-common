package lager

import "sync/atomic"

//ReConfigSink is a struct
type ReConfigSink struct {
	sink Sink

	minLogLevel int32
}

//NewReConfigSink is a function which returns struct object
func NewReConfigSink(sink Sink, initialMinLogLevel LogLevel) *ReConfigSink {
	return &ReConfigSink{
		sink: sink,

		minLogLevel: int32(initialMinLogLevel),
	}
}

//Log is a method which returns log level and log
func (sink *ReConfigSink) Log(level LogLevel, log []byte) {
	minLogLevel := LogLevel(atomic.LoadInt32(&sink.minLogLevel))

	if level < minLogLevel {
		return
	}

	sink.sink.Log(level, log)
}

//SetMinLevel is a function which sets minimum log level
func (sink *ReConfigSink) SetMinLevel(level LogLevel) {
	atomic.StoreInt32(&sink.minLogLevel, int32(level))
}

//GetMinLevel is a method which gets minimum log level
func (sink *ReConfigSink) GetMinLevel() LogLevel {
	return LogLevel(atomic.LoadInt32(&sink.minLogLevel))
}
