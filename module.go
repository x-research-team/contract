package contract

// KernelModule Модуль ядра
type KernelModule func(IService)

// KernelModules Модули ядра
type KernelModules []KernelModule

// ComponentModule Модуль компонента
type ComponentModule func(IComponent)

// ComponentModules Модули компонента
type ComponentModules []ComponentModule
