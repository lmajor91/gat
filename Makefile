BIN_NAME=gat
BIN_SRC =core/*.go

all: build

$(BIN_NAME): $(BIN_SRC)
	go build -o $@ $^

build: $(BIN_NAME)

clean:
	rm $(BIN_NAME)

