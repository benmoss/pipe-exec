.PHONY: all clean proto docker

all: build docker

build: proto out/client.exe out/server.exe out/hello.exe

out/client.exe: cmd/client/*
	GOOS=windows go build -o out/client.exe ./cmd/client

out/server.exe: cmd/server/* server/* proto
	GOOS=windows go build -o out/server.exe ./cmd/server

out/hello.exe: cmd/hello/* proto
	GOOS=windows go build -o out/hello.exe ./cmd/hello

proto:
	protoc -I server/ server/server.proto --go_out=plugins=grpc:server

clean:
	rm -rf out

docker: out/client.exe out/server.exe out/hello.exe
	docker build . --tag benmoss/pipe-exec:latest
	docker push benmoss/pipe-exec:latest
