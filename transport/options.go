package transport

import (
	"context"
	"github.com/butea/kunkka/codec"
	"time"

)

type Option func(*Options)

type DialOption func(*DialOptions)

type ListenOption func(*ListenOptions)

type Options struct {
	// Addrs is the list of intermediary addresses to connect to
	Addrs []string
	// Codec is the codec interface to use where headers are not supported
	// by the transport and the entire payload must be encoded
	Codec codec.Codec

	// Timeout sets the timeout for Send/Recv
	Timeout time.Duration
	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}


type DialOptions struct {
	// Tells the transport this is a streaming connection with
	// multiple calls to send/recv and that send may not even be called
	Stream bool
	// Timeout for dialing
	Timeout time.Duration

	// TODO: add tls options when dialling
	// Currently set in global options

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

type ListenOptions struct {
	// TODO: add tls options when listening
	// Currently set in global options

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context
}

// Addrs to use for transport
func Addrs(addrs ...string) Option {
	return func(o *Options) {
		o.Addrs = addrs
	}
}

// Codec sets the codec used for encoding where the transport
// does not support message headers
func Codec(c codec.Codec) Option {
	return func(o *Options) {
		o.Codec = c
	}
}

// Timeout sets the timeout for Send/Recv execution
func Timeout(t time.Duration) Option {
	return func(o *Options) {
		o.Timeout = t
	}
}

// Indicates whether this is a streaming connection
func WithStream() DialOption {
	return func(o *DialOptions) {
		o.Stream = true
	}
}

// Timeout used when dialling the remote side
func WithTimeout(d time.Duration) DialOption {
	return func(o *DialOptions) {
		o.Timeout = d
	}
}
