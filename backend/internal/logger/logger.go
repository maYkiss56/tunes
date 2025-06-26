package logger

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

type Logger struct {
	console *slog.Logger
	file    *slog.Logger
	queue   chan slog.Record
	wg      sync.WaitGroup
}

func getCallerInfo() (file string, line int) {
	// 4 означает глубина стека вызовов:
	// 0 - runtime.Caller
	// 1 - getCallerInfo
	// 2 - метод логгера (Debug/Info/Error)
	// 3 - место, где вызвали логгер
	_, file, line, _ = runtime.Caller(3)

	// Укорачиваем путь до файла (убираем часть пути до проекта)
	if idx := strings.Index(file, "tunes/"); idx != -1 {
		file = file[idx+len("tunes/"):]
	} else {
		file = filepath.Base(file)
	}

	return file, line
}

func New(ctx context.Context, logFilePath string) (*Logger, error) {
	consoleHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	fileHandler := slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	l := &Logger{
		console: slog.New(consoleHandler),
		file:    slog.New(fileHandler),
		queue:   make(chan slog.Record, 1000),
	}

	l.wg.Add(1)
	go l.proccessLogs(ctx)

	return l, nil
}

func (l *Logger) proccessLogs(ctx context.Context) {
	defer l.wg.Done()

	for {
		select {
		case r := <-l.queue:
			_ = l.console.Handler().Handle(context.Background(), r)
			_ = l.file.Handler().Handle(context.Background(), r)
		case <-ctx.Done():
			for len(l.queue) > 0 {
				r := <-l.queue
				_ = l.console.Handler().Handle(context.Background(), r)
				_ = l.file.Handler().Handle(context.Background(), r)
			}
			return
		}
	}
}

func (l *Logger) Debug(msg string, args ...any) {
	file, line := getCallerInfo()
	args = append(args, "file", file, "line", line)
	l.queue <- l.newRecord(slog.LevelDebug, msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.queue <- l.newRecord(slog.LevelInfo, msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	file, line := getCallerInfo()
	args = append(args, "file", file, "line", line)
	l.queue <- l.newRecord(slog.LevelError, msg, args...)
}

func (l *Logger) newRecord(level slog.Level, msg string, args ...any) slog.Record {
	rec := slog.NewRecord(time.Now(), level, msg, 0)

	if len(args)%2 == 0 {
		for i := 0; i < len(args); i += 2 {
			key, ok := args[i].(string)
			if !ok {
				continue
			}
			rec.AddAttrs(slog.Any(key, args[i+1]))
		}
	}

	return rec
}

func (l *Logger) Shutdown() {
	close(l.queue)
	l.wg.Wait()
}
