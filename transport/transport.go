package transport

// Transport is an interface which is used for communication between
// services. It uses connection based socket send/recv semantics and
// has various implementations; http, grpc, quic.
type Transport interface {
	Init(...Option) error
	Options() Options
	Dial(addr string, opts ...DialOption) (Client, error)
	Listen(addr string, opts ...ListenOption) (Listener, error)
	String() string
}

type Socket interface {
	Recv(*Message) error
	Send(*Message) error
	Close() error
	Local() string
	Remote() string
}


type Client interface {
	Socket
}

type Listener interface {
	Addr() string
	Close() error
	Accept(func(Socket)) error
}