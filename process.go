package contract

import "time"

// IProcess Процесс ядра
type IProcess interface {
	Pid() string
	Name() string
	Up(graceful bool) error
	Down(graceful bool) error
	Sleep(time.Duration) error
	Restart(graceful bool) error
	Pause() error
	Cron(rule string) error
	Stop() error
	Kill() error
	Sync(with string) error
	Backup(to string) error
}
