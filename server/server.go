package server

import (
	"github.com/butea/kunkka/registry"
	"github.com/butea/kunkka/transport"
	"github.com/micro/go-micro/v2/codec"
)

// Server is a simple micro server abstraction
type Server interface {
	Options() Options
	Init(...Option) error
	Handle(Handler) error
	NewHandler(interface{}, ...HandlerOption) Handler
	Start() error
	Stop() error
	String() string
}

type Option func(*Options)

type Handler interface {
	Name() string
	Handler() interface{}
	Endpoints() []*registry.Endpoint
	Options() HandlerOptions
}

// Request is a synchronous request interface
type Request interface {
	// Service name requested
	Service() string
	// The action requested
	Method() string
	// Endpoint name requested
	Endpoint() string
	// Content type provided
	ContentType() transport.ContentType
	// Header of the request
	Header() map[string]string
	// Body is the initial decoded value
	Body() interface{}
	// Read the undecoded request body
	Read() ([]byte, error)
	// The encoded message stream
	Codec() codec.Reader
	// Indicates whether its a stream
	Stream() bool
}

// Response is the response writer for unencoded messages
type Response interface {
	// Encoded writer
	Codec() codec.Writer
	// Write the header
	WriteHeader(map[string]string)
	// write a response directly to the client
	Write([]byte) error
}
