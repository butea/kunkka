package std

import (
	"github.com/butea/kunkka/transport"
	"net"
)

type stdSocket struct {
	conn net.Conn
}

type stdTransport struct {
	opts transport.Options
}

type stdClient struct {
	*stdSocket
	t *stdTransport
	opts transport.DialOptions
}

type stdListener struct {
	l net.Listener
	t *stdTransport
	opts transport.ListenOptions
}

func (c *stdClient) Close() error {
	return c.stdSocket.conn.Close()
}

func (s *stdSocket) Send(m *transport.Message) error {
	// todo:
	return nil
}

func (s *stdSocket) Recv(m *transport.Message) error {
	// todo:
	return nil
}

func (s *stdSocket) Close() error {
	return s.conn.Close()
}

func (s *stdSocket) Local() string {
	return s.conn.LocalAddr().String()
}

func (s *stdSocket) Remote() string {
	return s.conn.RemoteAddr().String()
}

func (l *stdListener) Addr() string {
	return l.l.Addr().String()
}

func (l *stdListener) Close() error {
	return l.l.Close()
}

func (l *stdListener) Accept(fn func(socket transport.Socket)) error {
	for {
		conn, err := l.l.Accept()
		if err != nil {
			return err
		}

		go func() {
			fn(&stdSocket{
				conn: conn,
			})
		}()
	}
}

func (t *stdTransport) Init(opts ...transport.Option) error {
	for _, o := range opts {
		o(&t.opts)
	}
	return nil
}

func (t *stdTransport) Options() transport.Options {
	return t.opts
}

func (t *stdTransport) Dial(addr string, opts ...transport.DialOption) (transport.Client, error) {
	var options transport.DialOptions
	for _, o := range opts {
		o(&options)
	}

	// todo: do sth
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &stdClient{
		stdSocket: &stdSocket{
			conn: conn,
		},
		t:         t,
		opts:      options,
	}, nil
}

func (t *stdTransport) Listen(addr string, opts ...transport.ListenOption) (transport.Listener, error) {
	var options transport.ListenOptions
	for _, o := range opts {
		o(&options)
	}

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &stdListener{
		l:    l,
		t:    t,
		opts: options,
	}, nil
}

func (t *stdTransport) String() string {
	return "std"
}

func NewTransport(opts ...transport.Option) transport.Transport {
	options := transport.Options{}

	for _, o := range opts {
		o(&options)
	}

	return &stdTransport{
		opts: options,
	}
}