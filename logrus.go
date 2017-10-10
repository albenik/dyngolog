package dyngolog

import (
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

type Wrapper interface {
	Normal() logrus.FieldLogger
	Verbose(uint8) logrus.FieldLogger
}

type wrapper struct {
	logger logrus.FieldLogger
	level  uint8
}

type formatter struct{}

var (
	silent = &logrus.Logger{
		Out:       ioutil.Discard,
		Formatter: new(formatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.PanicLevel,
	}
	zero [0]byte
)

func (f *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	return zero[:], nil
}

func (wr *wrapper) Normal() logrus.FieldLogger {
	return wr.logger
}

func (wr *wrapper) Verbose(level uint8) logrus.FieldLogger {
	if level <= wr.level {
		return wr.logger
	}
	return silent
}

func Wrap(logger logrus.FieldLogger, level uint8) Wrapper {
	return &wrapper{logger: logger, level: level}
}
