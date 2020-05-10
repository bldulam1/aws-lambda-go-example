BIN = ../../bin
SRC = ./functions

# Services
S_FIB = fibonacci
S_TOD = todos

all:
	go get ./...
	cd $(SRC)/$(S_FIB) && go build -o $(BIN)/$(S_FIB) ./...
	cd $(SRC)/$(S_TOD) && go build -o $(BIN)/$(S_TOD) ./...
