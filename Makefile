.PHONY: all clean proto

all: proto out/client.exe out/server.exe out/hello.exe
	cp out/* ~/Downloads/

out/client.exe: cmd/client/*
	GOOS=windows go build -o out/client.exe ./cmd/client

out/server.exe: cmd/server/* server/*
	GOOS=windows go build -o out/server.exe ./cmd/server

out/hello.exe: cmd/hello/*
	GOOS=windows go build -o out/hello.exe ./cmd/hello

proto:
	protoc -I server/ server/server.proto --go_out=plugins=grpc:server

clean:
	rm -rf out
