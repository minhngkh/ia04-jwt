package logger

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type logger struct {
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
	Logger                zerolog.Logger
}

func NewGormLogger() *logger {
	return &logger{
		Logger:                log.Logger,
		SkipErrRecordNotFound: true,
	}
}

func (l *logger) LogMode(gormLogger.LogLevel) gormLogger.Interface {
	return l
}

func (l *logger) Info(ctx context.Context, s string, args ...interface{}) {
	l.Logger.Info().Msgf(s, args...)
}

func (l *logger) Warn(ctx context.Context, s string, args ...interface{}) {
	l.Logger.Warn().Msgf(s, args...)
}

func (l *logger) Error(ctx context.Context, s string, args ...interface{}) {
	l.Logger.Error().Msgf(s, args...)
}

func (l *logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := map[string]interface{}{
		"sql":      sql,
		"duration": elapsed,
	}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		l.Logger.Error().Err(err).Fields(fields).Msg("[GORM] query error")
		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		l.Logger.Warn().Fields(fields).Msgf("[GORM] slow query")
		return
	}

	l.Logger.Debug().Fields(fields).Msgf("[GORM] query")
}
