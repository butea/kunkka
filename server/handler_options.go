package server


type HandlerOption func(*HandlerOptions)

type HandlerOptions struct {
	Internal bool
	Metadata map[string]map[string]string
}
