// Code generated by Kitex v0.8.0. DO NOT EDIT.

package testservice1

import (
	server "github.com/cloudwego/kitex/server"
	echo1 "github.com/felix021/dubbo-p1/kitex_gen/echo1"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler echo1.TestService1, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
