package types

//

type Anigo struct {
	Anime, Image, Url string
	Episode           int

	UseService   func(string, string, ...string) (interface{}, bool)
	GetDirectUrl func(string) (string, bool)
	GetServices  func() []string
}

//

type Process struct {
	Handler    func(*Anigo)
	Persistent bool
}

type Provider struct {
	Handler func(string) (string, bool)
	Solvers []string
}

type Service struct {
	Handler map[string]func(...string) []any
	Solver  string
}

//

type PluginType interface {
	Process | Provider | Service
}

type Plugin[T PluginType] struct {
	Author, Name, Url, Version string
	This                       T
}

//

type Anime struct {
	Anime, Image, Url string
	Episode           int
}

//

type SliceContent interface {
	complex64 | complex128 | float32 | float64 |
		uint8 | uint16 | uint32 | uint | uintptr |
		int8 | int16 | int32 | int |
		bool | string
}
