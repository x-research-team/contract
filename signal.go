package contract

// ISignal Сигнал ядру
type ISignal interface {
	IProcess
	Message() IMessage // Message
	Send() error
	Read() string // Read signal from buffer
}

// ISignals Сигналы ядра
type ISignals []ISignal

// ISignalBus Шина сигналов
type ISignalBus chan ISignal

// ISignalsBus Шины сигналов
type ISignalsBus []ISignalBus

// Generate Сгенерировать сигалы из массивов шин
func (channels ISignalsBus) Generate() <-chan ISignal {
	signal := make(chan ISignal)
	go func(x ISignalsBus) {
		defer close(signal)
		for _, c := range x {
			select {
			case a := <-c:
				signal <- a
			}
		}
	}(channels)
	return signal
}

// Send signals to the channel
func (c ISignalBus) Send(signals ...ISignal) {
	for _, s := range signals {
		select {
		case c <- s:
		}
	}
}

// Send singals to the channels
//	channels.Send(signal, signal, ...)
func (channels ISignalsBus) Send(signals ...ISignal) {
	for i := range channels {
		channels[i].Send(signals...)
	}
}
