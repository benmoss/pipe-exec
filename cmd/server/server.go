package main

import (
	"log"

	"github.com/Microsoft/go-winio"
	"github.com/benmoss/pipe-exec/server"
	"google.golang.org/grpc"
)

var (
	sddl = "D:P(A;;GA;;;BU)"
)

func main() {
	l, err := winio.ListenPipe(server.PipeName, &winio.PipeConfig{
		SecurityDescriptor: sddl,
		MessageMode:        true,
		InputBufferSize:    65536,
		OutputBufferSize:   65536,
	})
	if err != nil {
		log.Fatalf("err listening: %#v", err)
	}
	defer l.Close()

	g := grpc.NewServer()
	server.RegisterExecServer(g, &server.ExecImpl{})
	log.Println("listening on the pipe")
	log.Fatalln(g.Serve(l))
}
