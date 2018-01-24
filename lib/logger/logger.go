package logger

import (
	gologger "github.com/FZambia/go-logger"
)

// LevelMatches is a map that matches string level to logger level constants.
var LevelMatches = map[string]gologger.Level{
	"DEBUG":    gologger.LevelDebug,
	"INFO":     gologger.LevelInfo,
	"WARN":     gologger.LevelWarn,
	"ERROR":    gologger.LevelError,
	"CRITICAL": gologger.LevelCritical,
	"FATAL":    gologger.LevelFatal,
	"NONE":     gologger.LevelNone,
}

var (
	// LevelDebug level.
	LevelDebug = gologger.LevelDebug
	// LevelInfo level.
	LevelInfo = gologger.LevelInfo
	// LevelWarn level.
	LevelWarn = gologger.LevelWarn
	// LevelError level.
	LevelError = gologger.LevelError
	// LevelCritical level.
	LevelCritical = gologger.LevelCritical
	// LevelFatal level.
	LevelFatal = gologger.LevelFatal
	// LevelNone level.
	LevelNone = gologger.LevelNone
)

var (
	// DEBUG logger.
	DEBUG = gologger.DEBUG
	// INFO logger.
	INFO = gologger.INFO
	// WARN logger.
	WARN = gologger.WARN
	// ERROR logger.
	ERROR = gologger.ERROR
	// CRITICAL logger.
	CRITICAL = gologger.CRITICAL
	// FATAL logger
	FATAL = gologger.FATAL
)

// SetLogThreshold establishes a threshold where anything matching or above will be logged.
func SetLogThreshold(level gologger.Level) {
	gologger.SetLogThreshold(level)
}

// SetStdoutThreshold establishes a threshold where anything matching or above will be output.
func SetStdoutThreshold(level gologger.Level) {
	gologger.SetStdoutThreshold(level)
}

// SetLogFile sets the LogHandle to a io.writer created for the file behind the given file path.
// Will append to this file.
func SetLogFile(path string) error {
	return gologger.SetLogFile(path)
}

// SetLogFlag sets global log flag used in package.
func SetLogFlag(flag int) {
	gologger.SetLogFlag(flag)
}
