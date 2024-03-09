CMD = ./cmd
BIN = ./bin/glua

clean:
	rm -rf ./bin

test:
	go test ${CMD}

build:
	go build -o ${BIN} ${CMD}
