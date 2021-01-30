package contract

// IService Контракт сервиса ядра
type IService interface {
	IProcess
	Run() error
	AddComponent(IComponent)
	AddPlugin(path, name string) error
	RemovePlugin(name string) error
}
