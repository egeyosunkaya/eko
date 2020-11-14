

SRC = 
OBJ = $(SRC:.cc=.o)
EXEC = eko
PKG := example.com/EKO

build:
	GO111MODULE=on go build -o ./bin ./cmd/main.go

run: 
	./bin/main