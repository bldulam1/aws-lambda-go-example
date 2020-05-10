BIN = ../../bin
SRC = ./functions

# Services
S_FIB = fibonacci
S_FAC = factorial

all:
	go get ./...
	cd $(SRC)/$(S_FIB) && go build -o $(BIN)/$(S_FIB) ./...
	cd $(SRC)/$(S_FAC) && go build -o $(BIN)/$(S_FAC) ./...
