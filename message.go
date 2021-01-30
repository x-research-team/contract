package contract

// IMessage Сообщение сигнала ядра
type IMessage interface {
	Route() string
	Command() string
	Data() string
}
