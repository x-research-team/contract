package contract

import "github.com/google/uuid"

// IMessage Сообщение сигнала ядра
type IMessage interface {
	ID() uuid.UUID
	Route() string
	Command() string
	Data() string
}
