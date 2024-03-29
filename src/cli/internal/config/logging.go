// Copyright (c) Ollion
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"strings"
	"time"

	stdlog "log"

	"github.com/charmbracelet/log"
	"github.com/cldcvr/terrarium/src/pkg/confighelper"
)

// LogFormat JSON or TEXT (case insensitive)
func LogFormat() string {
	return confighelper.MustGetString("log.format")
}

// LogLevel one of: "fatal", "error", "warn", "info", "debug" (case sensitive)
func LogLevel() string {
	return confighelper.MustGetString("log.level")
}

// LoggerConfig sets up the given charmbracelet/log.Logger with defined configuration.
//
// Example:
// To update standard logger configuration, do this:
//
// LoggerConfig(log.Default())
func LoggerConfig(logger *log.Logger) {
	// Set log formatter
	formatter := LogFormat()
	if strings.EqualFold(formatter, "text") {
		logger.SetFormatter(log.TextFormatter)
	} else if strings.EqualFold(formatter, "json") {
		logger.SetFormatter(log.JSONFormatter)
	}

	// Set log level
	levelStr := LogLevel()
	level := log.ParseLevel(levelStr)
	logger.SetLevel(level)

	if level == log.DebugLevel {
		logger.SetReportCaller(true)
	} else {
		logger.SetReportTimestamp(false)
		logger.SetTimeFormat(time.Kitchen)
	}
}

// LoggerConfigDefault sets up the default loggers with defined configuration.
func LoggerConfigDefault() {
	logger := log.Default()
	LoggerConfig(logger) // setup `logger` object from config

	// create standard logger object
	stdLogger := logger.StandardLog(log.StandardLogOptions{ForceLevel: log.DebugLevel})

	// update default standard logger
	stdlog.Default().SetOutput(stdLogger.Writer())
	stdlog.Default().SetFlags(stdLogger.Flags())
}
