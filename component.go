package contract

// IComponent Компоненты
type IComponent interface {
	IService
	Configure() error
	Route() string
	Write(message IMessage) error
	Read() string
	Send(message IMessage)
}
