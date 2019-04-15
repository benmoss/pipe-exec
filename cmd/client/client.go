package main

import (
	"context"
	"io/ioutil"
	"log"
	"net"

	"github.com/Microsoft/go-winio"
	"github.com/benmoss/pipe-exec/server"
	"google.golang.org/grpc"
)

const (
	sddl      = "D:P(A;;GA;;;BA)(A;;GA;;;SY)"
	pipeName  = `\\.\pipe\foo`
	chunkSize = 64 * 1024 // 64 KiB
)

func main() {
	conn, err := grpc.Dial(pipeName, grpc.WithContextDialer(func(_ context.Context, target string) (net.Conn, error) {
		return winio.DialPipe(target, nil)
	}), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("err dialing: %#v", err)
	}
	defer conn.Close()

	client := server.NewExecClient(conn)
	payload, err := ioutil.ReadFile("hello.exe")
	if err != nil {
		log.Fatalf("err reading file: %#v", err)
	}
	stream, err := client.Exec(context.Background())
	if err != nil {
		log.Fatalf("err creating client: %#v", err)
	}

	exe := &server.Executable{Args: []string{"hey", "cool"}}
	for currentByte := 0; currentByte < len(payload); currentByte += chunkSize {
		if currentByte+chunkSize > len(payload) {
			exe.Chunk = payload[currentByte:len(payload)]
		} else {
			exe.Chunk = payload[currentByte : currentByte+chunkSize]
		}
		if err := stream.Send(exe); err != nil {
			log.Fatalf("err sending chunk: %#v", err)
			return
		}
	}
	log.Println("sent all teh chunks")

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("err closing: %#v", err)
	}

	log.Printf("got stdout: %q, stderr: %q", string(resp.Stdout), string(resp.Stderr))
}
