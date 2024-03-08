CMD = ./cmd
BIN = ./bin/glua

test:
	go test ${CMD}

build:
	go build -o ${BIN} ${CMD}
