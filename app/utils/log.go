package utils

import (
	"github.com/rs/zerolog"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// Log sets log file location and options
var Log = zerolog.New(&lumberjack.Logger{
	Filename:   "/log/error.log",
	MaxSize:    50, // megabytes
	MaxBackups: 30,
	MaxAge:     90, //days
}).With().Timestamp().Logger()
